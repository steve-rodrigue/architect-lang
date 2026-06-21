package antlr

import (
	"testing"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/applications"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/endpoints"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/services"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
)

func TestParserApplicationObjectParsesSimpleObject(t *testing.T) {
	app := NewParserApplication()

	obj, err := app.Object(`
object User {
  id UUID primary
  email String unique
  displayName String?
}
`)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if obj.Name() != "User" {
		t.Fatalf("expected object name User, got %s", obj.Name())
	}

	if obj.HistoryOf() != nil {
		t.Fatalf("expected no history_of")
	}

	if len(obj.Fields()) != 3 {
		t.Fatalf("expected 3 fields, got %d", len(obj.Fields()))
	}

	id := obj.Fields()[0]
	if id.Name() != "id" {
		t.Fatalf("expected first field id, got %s", id.Name())
	}

	if id.Type().Name() != "UUID" {
		t.Fatalf("expected id type UUID, got %s", id.Type().Name())
	}

	if !hasModifier(id, objects.FieldModifierPrimary) {
		t.Fatalf("expected id to have primary modifier")
	}

	email := obj.Fields()[1]
	if !hasModifier(email, objects.FieldModifierUnique) {
		t.Fatalf("expected email to have unique modifier")
	}

	displayName := obj.Fields()[2]
	if !displayName.Type().IsOptional() {
		t.Fatalf("expected displayName to be optional")
	}
}

func TestParserApplicationObjectParsesHistoryObject(t *testing.T) {
	app := NewParserApplication()

	obj, err := app.Object(`
object PostHistory history_of Post {
  id UUID primary
  post Post
  changedAt Time
}
`)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if obj.Name() != "PostHistory" {
		t.Fatalf("expected PostHistory, got %s", obj.Name())
	}

	if obj.HistoryOf() == nil {
		t.Fatalf("expected history_of Post")
	}

	if obj.HistoryOf().Name() != "Post" {
		t.Fatalf("expected history_of Post, got %s", obj.HistoryOf().Name())
	}
}

func TestParserApplicationObjectParsesDefaultValueOnOptionalString(t *testing.T) {
	app := NewParserApplication()

	obj, err := app.Object(`
object User {
  displayName String? "Default Name"
}
`)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	field := obj.Fields()[0]

	if !field.Type().IsOptional() {
		t.Fatalf("expected field to be optional")
	}

	if !field.Type().HasDefaultValue() {
		t.Fatalf("expected default value")
	}

	if field.Type().DefaultValue().Kind() != common.ValueString {
		t.Fatalf("expected string default")
	}

	if field.Type().DefaultValue().Raw() != "Default Name" {
		t.Fatalf("expected Default Name, got %s", field.Type().DefaultValue().Raw())
	}
}

func TestParserApplicationObjectRejectsDefaultValueOnRequiredField(t *testing.T) {
	app := NewParserApplication()

	_, err := app.Object(`
object User {
  displayName String "Default Name"
}
`)
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestParserApplicationObjectParsesNumberConstraints(t *testing.T) {
	app := NewParserApplication()

	obj, err := app.Object(`
object User {
  onePlus Int+
  zeroPlus Int*
  between Int[0,23]
  threeOrHigher Int[3,]
  rangeNumber Float[34,38]
  optionalDefault Int+? 45
}
`)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	assertNumberConstraint(t, obj.Fields()[0], common.NumberConstraintOneOrMore, false, "", false, "")
	assertNumberConstraint(t, obj.Fields()[1], common.NumberConstraintZeroOrMore, false, "", false, "")
	assertNumberConstraint(t, obj.Fields()[2], common.NumberConstraintRange, true, "0", true, "23")
	assertNumberConstraint(t, obj.Fields()[3], common.NumberConstraintRange, true, "3", false, "")
	assertNumberConstraint(t, obj.Fields()[4], common.NumberConstraintRange, true, "34", true, "38")

	optionalDefault := obj.Fields()[5]
	if !optionalDefault.Type().IsOptional() {
		t.Fatalf("expected optionalDefault to be optional")
	}

	if !optionalDefault.Type().HasDefaultValue() {
		t.Fatalf("expected optionalDefault to have default")
	}

	if optionalDefault.Type().DefaultValue().Kind() != common.ValueInt {
		t.Fatalf("expected optionalDefault default kind int")
	}

	if optionalDefault.Type().DefaultValue().Raw() != "45" {
		t.Fatalf("expected default 45, got %s", optionalDefault.Type().DefaultValue().Raw())
	}
}

func TestParserApplicationObjectParsesRenamedFromAndDeprecated(t *testing.T) {
	app := NewParserApplication()

	obj, err := app.Object(`
object User {
  fullName String renamed_from displayName
  oldField String deprecated
}
`)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	fullName := obj.Fields()[0]
	renamed := renamedFromModifier(fullName)
	if renamed == nil {
		t.Fatalf("expected renamed_from modifier")
	}

	if renamed.Name() != "displayName" {
		t.Fatalf("expected renamed_from displayName, got %s", renamed.Name())
	}

	oldField := obj.Fields()[1]
	if !hasModifier(oldField, objects.FieldModifierDeprecated) {
		t.Fatalf("expected deprecated modifier")
	}
}

func TestParserApplicationEndpointParsesUpdatePost(t *testing.T) {
	app := NewParserApplication()

	endpoint, err := app.Endpoint(`
endpoint UpdatePost PATCH "/posts/{id}" {
  input {
    id UUID (path?|body?)
    title String body?
    content Text body?
    changedBy UUID body
  }

  fetch post:Post from MainDB where id == input.id

  create history:PostHistory {
    post = post
    title = post.title
    content = post.content
    postedBy = post.postedBy
    changedBy = input.changedBy
    changedAt = now()
    operation = "update"
  }

  store history to MainDB

  update post {
    title = input.title
    content = input.content
    updatedAt = now()
  }

  store post to MainDB
  store post to GraphDB

  emit event:PostUpdated {
    post = post
    history = history
  }

  return post
}
`)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if endpoint.Name() != "UpdatePost" {
		t.Fatalf("expected UpdatePost, got %s", endpoint.Name())
	}

	if endpoint.Method() != endpoints.HTTPMethodPatch {
		t.Fatalf("expected PATCH, got %s", endpoint.Method())
	}

	if endpoint.Path() != "/posts/{id}" {
		t.Fatalf("expected /posts/{id}, got %s", endpoint.Path())
	}

	fields := endpoint.Input().Fields()
	if len(fields) != 4 {
		t.Fatalf("expected 4 input fields, got %d", len(fields))
	}

	id := fields[0]
	if id.Name() != "id" {
		t.Fatalf("expected id field, got %s", id.Name())
	}

	if id.Type().Name() != "UUID" {
		t.Fatalf("expected id type UUID, got %s", id.Type().Name())
	}

	if id.Sources().Kind() != endpoints.InputSourceRuleExactlyOneOf {
		t.Fatalf("expected id source rule exactly_one_of, got %s", id.Sources().Kind())
	}

	if len(id.Sources().Sources()) != 2 {
		t.Fatalf("expected 2 id sources, got %d", len(id.Sources().Sources()))
	}

	if id.Sources().Sources()[0].Kind() != endpoints.InputSourcePath || !id.Sources().Sources()[0].IsOptional() {
		t.Fatalf("expected first id source path?")
	}

	if id.Sources().Sources()[1].Kind() != endpoints.InputSourceBody || !id.Sources().Sources()[1].IsOptional() {
		t.Fatalf("expected second id source body?")
	}

	if fields[1].Name() != "title" || fields[1].Sources().Sources()[0].Kind() != endpoints.InputSourceBody || !fields[1].Sources().Sources()[0].IsOptional() {
		t.Fatalf("expected title String body?")
	}

	if fields[3].Name() != "changedBy" || fields[3].Sources().Sources()[0].Kind() != endpoints.InputSourceBody || fields[3].Sources().Sources()[0].IsOptional() {
		t.Fatalf("expected changedBy UUID body")
	}

	actions := endpoint.Actions()
	if len(actions) != 8 {
		t.Fatalf("expected 8 actions, got %d", len(actions))
	}

	fetch := actions[0].(workflows.FetchAction)
	if fetch.Kind() != workflows.ActionKindFetch {
		t.Fatalf("expected action 0 fetch")
	}
	if fetch.Variable().Name() != "post" || fetch.Variable().Type().Name() != "Post" {
		t.Fatalf("expected fetch variable post:Post")
	}
	if fetch.Source() != "MainDB" {
		t.Fatalf("expected fetch source MainDB")
	}
	if fetch.Condition().Operator() != workflows.ComparatorEqual {
		t.Fatalf("expected fetch condition ==")
	}

	create := actions[1].(workflows.CreateAction)
	if create.Kind() != workflows.ActionKindCreate {
		t.Fatalf("expected action 1 create")
	}
	if create.Variable().Name() != "history" || create.Variable().Type().Name() != "PostHistory" {
		t.Fatalf("expected create history:PostHistory")
	}
	if len(create.Assignments()) != 7 {
		t.Fatalf("expected 7 history assignments, got %d", len(create.Assignments()))
	}

	storeHistory := actions[2].(workflows.StoreAction)
	if storeHistory.VariableName() != "history" || storeHistory.Destination() != "MainDB" {
		t.Fatalf("expected store history to MainDB")
	}

	update := actions[3].(workflows.UpdateAction)
	if update.VariableName() != "post" {
		t.Fatalf("expected update post")
	}
	if len(update.Assignments()) != 3 {
		t.Fatalf("expected 3 update assignments, got %d", len(update.Assignments()))
	}

	emit := actions[6].(workflows.EmitAction)
	if emit.Kind() != workflows.ActionKindEmit {
		t.Fatalf("expected action 6 emit")
	}
	if emit.Variable().Name() != "event" || emit.Variable().Type().Name() != "PostUpdated" {
		t.Fatalf("expected emit event:PostUpdated")
	}
	if len(emit.Assignments()) != 2 {
		t.Fatalf("expected 2 emit assignments")
	}

	if actions[7].Kind() != workflows.ActionKindReturn {
		t.Fatalf("expected action 7 return")
	}
}

func TestParserApplicationEndpointParsesCreatePostWithFunctionCallsAndLiterals(t *testing.T) {
	app := NewParserApplication()

	endpoint, err := app.Endpoint(`
endpoint CreatePost POST "/posts" {
  input {
    title String body
    content Text body
    postedBy UUID body
  }

  create post:Post {
    title = input.title
    content = input.content
    postedBy = input.postedBy
    createdAt = now()
    slug = slugify(input.title)
    score = calculateScore(input.title, 12, 12.5, true, false)
  }

  store post to MainDB
  store post to GraphDB

  emit event:PostCreated {
    post = post
  }

  return post
}
`)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if endpoint.Method() != endpoints.HTTPMethodPost {
		t.Fatalf("expected POST, got %s", endpoint.Method())
	}

	actions := endpoint.Actions()
	if len(actions) != 5 {
		t.Fatalf("expected 5 actions, got %d", len(actions))
	}

	create := actions[0].(workflows.CreateAction)
	assignments := create.Assignments()

	assertAssignmentSelector(t, assignments[0], "title", []string{"input", "title"})
	assertAssignmentSelector(t, assignments[1], "content", []string{"input", "content"})
	assertAssignmentFunctionCall(t, assignments[3], "createdAt", "now", 0)
	assertAssignmentFunctionCall(t, assignments[4], "slug", "slugify", 1)
	assertAssignmentFunctionCall(t, assignments[5], "score", "calculateScore", 5)

	scoreCall := assignments[5].Value().(workflows.FunctionCallExpression)
	assertLiteral(t, scoreCall.Arguments()[1], common.ValueInt, "12")
	assertLiteral(t, scoreCall.Arguments()[2], common.ValueFloat, "12.5")
	assertLiteral(t, scoreCall.Arguments()[3], common.ValueBool, "true")
	assertLiteral(t, scoreCall.Arguments()[4], common.ValueBool, "false")

	storeOne := actions[1].(workflows.StoreAction)
	if storeOne.VariableName() != "post" || storeOne.Destination() != "MainDB" {
		t.Fatalf("expected store post to MainDB")
	}

	storeTwo := actions[2].(workflows.StoreAction)
	if storeTwo.VariableName() != "post" || storeTwo.Destination() != "GraphDB" {
		t.Fatalf("expected store post to GraphDB")
	}

	emit := actions[3].(workflows.EmitAction)
	if emit.Variable().Name() != "event" || emit.Variable().Type().Name() != "PostCreated" {
		t.Fatalf("expected emit event:PostCreated")
	}

	ret := actions[4].(workflows.ReturnAction)
	if ret.Expression().Kind() != workflows.ExpressionKindIdentifier {
		t.Fatalf("expected identifier return expression")
	}
}

func TestParserApplicationEndpointParsesAllInputSources(t *testing.T) {
	app := NewParserApplication()

	endpoint, err := app.Endpoint(`
endpoint SourceTest GET "/source" {
  input {
    pathId UUID path
    queryId UUID query?
    bodyId UUID body
    token String header
    session String cookie?
  }

  return pathId
}
`)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	fields := endpoint.Input().Fields()
	if len(fields) != 5 {
		t.Fatalf("expected 5 fields, got %d", len(fields))
	}

	assertSingleSource(t, fields[0], endpoints.InputSourcePath, false)
	assertSingleSource(t, fields[1], endpoints.InputSourceQuery, true)
	assertSingleSource(t, fields[2], endpoints.InputSourceBody, false)
	assertSingleSource(t, fields[3], endpoints.InputSourceHeader, false)
	assertSingleSource(t, fields[4], endpoints.InputSourceCookie, true)
}

func TestParserApplicationEndpointParsesAllHTTPMethods(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		expected endpoints.HTTPMethod
	}{
		{"GetEndpoint", "GET", endpoints.HTTPMethodGet},
		{"PostEndpoint", "POST", endpoints.HTTPMethodPost},
		{"PutEndpoint", "PUT", endpoints.HTTPMethodPut},
		{"PatchEndpoint", "PATCH", endpoints.HTTPMethodPatch},
		{"DeleteEndpoint", "DELETE", endpoints.HTTPMethodDelete},
	}

	for _, test := range tests {
		t.Run(test.method, func(t *testing.T) {
			app := NewParserApplication()

			endpoint, err := app.Endpoint(`
endpoint ` + test.name + ` ` + test.method + ` "/test" {
  input {
    id UUID path
  }

  return id
}
`)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			if endpoint.Method() != test.expected {
				t.Fatalf("expected %s, got %s", test.expected, endpoint.Method())
			}
		})
	}
}

func TestParserApplicationEndpointParsesComparators(t *testing.T) {
	tests := []struct {
		name     string
		operator string
		expected workflows.Comparator
	}{
		{"EqEndpoint", "==", workflows.ComparatorEqual},
		{"NeqEndpoint", "!=", workflows.ComparatorNotEqual},
		{"LtEndpoint", "<", workflows.ComparatorLessThan},
		{"LteEndpoint", "<=", workflows.ComparatorLessThanOrEqual},
		{"GtEndpoint", ">", workflows.ComparatorGreaterThan},
		{"GteEndpoint", ">=", workflows.ComparatorGreaterThanOrEqual},
	}

	for _, test := range tests {
		t.Run(test.operator, func(t *testing.T) {
			app := NewParserApplication()

			endpoint, err := app.Endpoint(`
endpoint ` + test.name + ` GET "/test" {
  input {
    id UUID path
  }

  fetch post:Post from MainDB where post.score ` + test.operator + ` 10

  return post
}
`)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			fetch := endpoint.Actions()[0].(workflows.FetchAction)
			if fetch.Condition().Operator() != test.expected {
				t.Fatalf("expected comparator %s, got %s", test.expected, fetch.Condition().Operator())
			}
		})
	}
}

func TestParserApplicationEndpointParsesListAndNumberConstraints(t *testing.T) {
	app := NewParserApplication()

	endpoint, err := app.Endpoint(`
endpoint SearchPosts GET "/posts" {
  input {
    limit Int[1,100] query?
    offset Int* query?
    scores List<Float>? body?
  }

  fetch posts:List<Post> from MainDB where limit <= 100

  return posts
}
`)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	fields := endpoint.Input().Fields()

	limitConstraint := fields[0].Type().NumberConstraint()
	if limitConstraint.Kind() != common.NumberConstraintRange {
		t.Fatalf("expected range constraint")
	}
	if !limitConstraint.HasMin() || limitConstraint.Min().Raw() != "1" {
		t.Fatalf("expected min 1")
	}
	if !limitConstraint.HasMax() || limitConstraint.Max().Raw() != "100" {
		t.Fatalf("expected max 100")
	}

	offsetConstraint := fields[1].Type().NumberConstraint()
	if offsetConstraint.Kind() != common.NumberConstraintZeroOrMore {
		t.Fatalf("expected zero_or_more constraint")
	}

	if !fields[2].Type().IsList() {
		t.Fatalf("expected scores to be list")
	}
	if fields[2].Type().Name() != "Float" {
		t.Fatalf("expected list type Float, got %s", fields[2].Type().Name())
	}
	if !fields[2].Type().IsOptional() {
		t.Fatalf("expected scores optional")
	}

	fetch := endpoint.Actions()[0].(workflows.FetchAction)
	if !fetch.Variable().Type().IsList() {
		t.Fatalf("expected fetched posts to be list")
	}
	if fetch.Variable().Type().Name() != "Post" {
		t.Fatalf("expected Post list")
	}
}

func TestParserApplicationEndpointRejectsInvalidDefaultSyntax(t *testing.T) {
	app := NewParserApplication()

	_, err := app.Endpoint(`
endpoint Bad GET "/bad" {
  input {
    id UUID path
  }

  return
}
`)
	if err == nil {
		t.Fatalf("expected error")
	}
}

func hasModifier(field objects.Field, kind objects.FieldModifierKind) bool {
	for _, modifier := range field.Modifiers() {
		switch kind {
		case objects.FieldModifierPrimary:
			if modifier.IsPrimary() {
				return true
			}
		case objects.FieldModifierUnique:
			if modifier.IsUnique() {
				return true
			}
		case objects.FieldModifierRenamedFrom:
			if modifier.IsRenameFrom() {
				return true
			}
		case objects.FieldModifierDeprecated:
			if modifier.IsDeprecated() {
				return true
			}
		default:
			if modifier.Kind() == kind {
				return true
			}
		}
	}

	return false
}

func renamedFromModifier(field objects.Field) objects.RenamedFromFieldModifier {
	for _, modifier := range field.Modifiers() {
		if renamed, ok := modifier.(objects.RenamedFromFieldModifier); ok {
			return renamed
		}
	}

	return nil
}

func assertNumberConstraint(
	t *testing.T,
	field objects.Field,
	expectedKind common.NumberConstraintKind,
	expectMin bool,
	expectedMin string,
	expectMax bool,
	expectedMax string,
) {
	t.Helper()

	constraint := field.Type().NumberConstraint()
	if constraint == nil {
		t.Fatalf("expected number constraint")
	}

	if constraint.Kind() != expectedKind {
		t.Fatalf("expected constraint kind %s, got %s", expectedKind, constraint.Kind())
	}

	if constraint.HasMin() != expectMin {
		t.Fatalf("expected HasMin %v, got %v", expectMin, constraint.HasMin())
	}

	if expectMin {
		if constraint.Min() == nil {
			t.Fatalf("expected min %s, got nil", expectedMin)
		}

		if constraint.Min().Raw() != expectedMin {
			t.Fatalf("expected min %s, got %s", expectedMin, constraint.Min().Raw())
		}
	}

	if constraint.HasMax() != expectMax {
		t.Fatalf("expected HasMax %v, got %v", expectMax, constraint.HasMax())
	}

	if expectMax {
		if constraint.Max() == nil {
			t.Fatalf("expected max %s, got nil", expectedMax)
		}

		if constraint.Max().Raw() != expectedMax {
			t.Fatalf("expected max %s, got %s", expectedMax, constraint.Max().Raw())
		}
	}
}

func TestParserApplicationConsumerParsesPostCreated(t *testing.T) {
	app := NewParserApplication()

	consumer, err := app.Consumer(`
consumes PostCreated {
  fetch post:Post from GraphDB where id == event.post.id

  call EmbeddingService.Generate {
    text = post.content
    dimensions = 1024
    normalize = true
  }

  create embedding:PostEmbedding {
    post = post
    vector = result.vector
    score = 0.98
    createdAt = now()
  }

  store embedding to VectorDB

  update post {
    embedding = embedding
    scoredAt = now()
  }

  emit event:PostScored {
    post = post
    embedding = embedding
  }

  return embedding
}
`)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if consumer.EventName() != "PostCreated" {
		t.Fatalf("expected PostCreated, got %s", consumer.EventName())
	}

	actions := consumer.Actions()
	if len(actions) != 7 {
		t.Fatalf("expected 7 actions, got %d", len(actions))
	}

	fetch := actions[0].(workflows.FetchAction)
	if fetch.Variable().Name() != "post" || fetch.Variable().Type().Name() != "Post" {
		t.Fatalf("expected fetch post:Post")
	}
	if fetch.Source() != "GraphDB" {
		t.Fatalf("expected GraphDB")
	}
	if fetch.Condition().Operator() != workflows.ComparatorEqual {
		t.Fatalf("expected ==")
	}

	call := actions[1].(workflows.CallAction)
	if call.Target() != "EmbeddingService.Generate" {
		t.Fatalf("expected EmbeddingService.Generate, got %s", call.Target())
	}
	if len(call.Assignments()) != 3 {
		t.Fatalf("expected 3 call assignments")
	}
	assertAssignmentSelector(t, call.Assignments()[0], "text", []string{"post", "content"})
	assertLiteral(t, call.Assignments()[1].Value(), common.ValueInt, "1024")
	assertLiteral(t, call.Assignments()[2].Value(), common.ValueBool, "true")

	create := actions[2].(workflows.CreateAction)
	if create.Variable().Name() != "embedding" || create.Variable().Type().Name() != "PostEmbedding" {
		t.Fatalf("expected create embedding:PostEmbedding")
	}
	if len(create.Assignments()) != 4 {
		t.Fatalf("expected 4 create assignments")
	}
	assertAssignmentSelector(t, create.Assignments()[1], "vector", []string{"result", "vector"})
	assertLiteral(t, create.Assignments()[2].Value(), common.ValueFloat, "0.98")
	assertAssignmentFunctionCall(t, create.Assignments()[3], "createdAt", "now", 0)

	store := actions[3].(workflows.StoreAction)
	if store.VariableName() != "embedding" || store.Destination() != "VectorDB" {
		t.Fatalf("expected store embedding to VectorDB")
	}

	update := actions[4].(workflows.UpdateAction)
	if update.VariableName() != "post" {
		t.Fatalf("expected update post")
	}
	if len(update.Assignments()) != 2 {
		t.Fatalf("expected 2 update assignments")
	}

	emit := actions[5].(workflows.EmitAction)
	if emit.Variable().Name() != "event" || emit.Variable().Type().Name() != "PostScored" {
		t.Fatalf("expected emit event:PostScored")
	}

	ret := actions[6].(workflows.ReturnAction)
	if ret.Expression().Kind() != workflows.ExpressionKindIdentifier {
		t.Fatalf("expected identifier return")
	}
}

func TestParserApplicationConsumerParsesComparators(t *testing.T) {
	tests := []struct {
		name     string
		operator string
		expected workflows.Comparator
	}{
		{"Equal", "==", workflows.ComparatorEqual},
		{"NotEqual", "!=", workflows.ComparatorNotEqual},
		{"LessThan", "<", workflows.ComparatorLessThan},
		{"LessThanOrEqual", "<=", workflows.ComparatorLessThanOrEqual},
		{"GreaterThan", ">", workflows.ComparatorGreaterThan},
		{"GreaterThanOrEqual", ">=", workflows.ComparatorGreaterThanOrEqual},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			app := NewParserApplication()

			consumer, err := app.Consumer(`
consumes ScoreRequested {
  fetch post:Post from GraphDB where post.score ` + test.operator + ` 10
}
`)
			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}

			fetch := consumer.Actions()[0].(workflows.FetchAction)
			if fetch.Condition().Operator() != test.expected {
				t.Fatalf("expected %s, got %s", test.expected, fetch.Condition().Operator())
			}
		})
	}
}

func TestParserApplicationConsumerParsesTypeReferences(t *testing.T) {
	app := NewParserApplication()

	consumer, err := app.Consumer(`
consumes BatchRequested {
  fetch posts:List<Post> from MainDB where limit <= 100

  create batch:Batch {
    posts = posts
    min = 1
    max = 100
  }

  return posts
}
`)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	fetch := consumer.Actions()[0].(workflows.FetchAction)
	if !fetch.Variable().Type().IsList() {
		t.Fatalf("expected posts to be list")
	}
	if fetch.Variable().Type().Name() != "Post" {
		t.Fatalf("expected Post, got %s", fetch.Variable().Type().Name())
	}
}

func TestParserApplicationConsumerRejectsInvalidSyntax(t *testing.T) {
	app := NewParserApplication()

	_, err := app.Consumer(`
consumes BadConsumer {
  return
}
`)
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestParserApplicationApplicationParsesFullApplication(t *testing.T) {
	app := NewParserApplication()

	application, err := app.Application(`
application PostService {
  emits rest on 8080
  listens events on 9090

  objects {
    "Post.arch"
    "PostHistory.arch"
    "PostEmbedding.arch"
  }

  endpoints {
    "CreatePost.arch"
    "UpdatePost.arch"
  }

  consumers {
    "PostCreated.arch"
    "PostUpdated.arch"
  }
}
`)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if application.Name() != "PostService" {
		t.Fatalf("expected PostService, got %s", application.Name())
	}

	ports := application.Ports()
	if len(ports) != 2 {
		t.Fatalf("expected 2 ports, got %d", len(ports))
	}

	if ports[0].Direction() != applications.PortDirectionEmit {
		t.Fatalf("expected first port emit, got %s", ports[0].Direction())
	}

	if ports[0].Kind() != applications.PortKind("rest") {
		t.Fatalf("expected first port rest, got %s", ports[0].Kind())
	}

	if ports[0].Number() != 8080 {
		t.Fatalf("expected first port 8080, got %d", ports[0].Number())
	}

	if ports[1].Direction() != applications.PortDirectionListen {
		t.Fatalf("expected second port listen, got %s", ports[1].Direction())
	}

	if ports[1].Kind() != applications.PortKind("events") {
		t.Fatalf("expected second port events, got %s", ports[1].Kind())
	}

	if ports[1].Number() != 9090 {
		t.Fatalf("expected second port 9090, got %d", ports[1].Number())
	}

	assertStringSlice(t, application.ObjectFiles(), []string{
		"Post.arch",
		"PostHistory.arch",
		"PostEmbedding.arch",
	})

	assertStringSlice(t, application.EndpointFiles(), []string{
		"CreatePost.arch",
		"UpdatePost.arch",
	})

	assertStringSlice(t, application.ConsumerFiles(), []string{
		"PostCreated.arch",
		"PostUpdated.arch",
	})
}

func TestParserApplicationApplicationParsesEmptyBlocks(t *testing.T) {
	app := NewParserApplication()

	application, err := app.Application(`
application EmptyService {
  objects {
  }

  endpoints {
  }

  consumers {
  }
}
`)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if application.Name() != "EmptyService" {
		t.Fatalf("expected EmptyService, got %s", application.Name())
	}

	if len(application.Ports()) != 0 {
		t.Fatalf("expected 0 ports, got %d", len(application.Ports()))
	}

	if len(application.ObjectFiles()) != 0 {
		t.Fatalf("expected 0 object files, got %d", len(application.ObjectFiles()))
	}

	if len(application.EndpointFiles()) != 0 {
		t.Fatalf("expected 0 endpoint files, got %d", len(application.EndpointFiles()))
	}

	if len(application.ConsumerFiles()) != 0 {
		t.Fatalf("expected 0 consumer files, got %d", len(application.ConsumerFiles()))
	}
}

func TestParserApplicationApplicationRejectsInvalidSyntax(t *testing.T) {
	app := NewParserApplication()

	_, err := app.Application(`
application BadService {
  emits rest on
}
`)
	if err == nil {
		t.Fatalf("expected error")
	}
}

func assertStringSlice(
	t *testing.T,
	actual []string,
	expected []string,
) {
	t.Helper()

	if len(actual) != len(expected) {
		t.Fatalf("expected %d items, got %d: %v", len(expected), len(actual), actual)
	}

	for i := range expected {
		if actual[i] != expected[i] {
			t.Fatalf("expected item %d to be %q, got %q", i, expected[i], actual[i])
		}
	}
}

func TestParserApplicationServiceParsesApplicationBackedService(t *testing.T) {
	app := NewParserApplication()

	service, err := app.Service(`
service BlogAPI go {
  exposes 8080
  application "blog_api.arch"
}
`)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if service.Name() != "BlogAPI" {
		t.Fatalf("expected BlogAPI, got %s", service.Name())
	}

	if service.Kind() != services.ServiceKindGo {
		t.Fatalf("expected go, got %s", service.Kind())
	}

	assertIntSlice(t, service.ExposedPorts(), []int{8080})

	if service.ApplicationFile() != "blog_api.arch" {
		t.Fatalf("expected blog_api.arch, got %s", service.ApplicationFile())
	}

	if len(service.DependsOn()) != 0 {
		t.Fatalf("expected no explicit dependencies when application exists")
	}
}

func TestParserApplicationServiceParsesInfrastructureService(t *testing.T) {
	app := NewParserApplication()

	service, err := app.Service(`
service MainDB postgres {
  version "17"
  image "postgres:17"
  exposes 5432
  stores User
  stores Post
  stores PostHistory
}
`)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if service.Name() != "MainDB" {
		t.Fatalf("expected MainDB, got %s", service.Name())
	}

	if service.Kind() != services.ServiceKindPostgres {
		t.Fatalf("expected postgres, got %s", service.Kind())
	}

	if service.Version() != "17" {
		t.Fatalf("expected version 17, got %s", service.Version())
	}

	if service.Image() != "postgres:17" {
		t.Fatalf("expected postgres:17, got %s", service.Image())
	}

	assertIntSlice(t, service.ExposedPorts(), []int{5432})
	assertStringSlice(t, service.Stores(), []string{"User", "Post", "PostHistory"})
}

func TestParserApplicationServiceParsesEventBusService(t *testing.T) {
	app := NewParserApplication()

	service, err := app.Service(`
service Events hatchet {
  version "latest"
  event PostCreated
  event PostUpdated
  event PostScored
}
`)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if service.Name() != "Events" {
		t.Fatalf("expected Events, got %s", service.Name())
	}

	if service.Kind() != services.ServiceKindHatchet {
		t.Fatalf("expected hatchet, got %s", service.Kind())
	}

	if service.Version() != "latest" {
		t.Fatalf("expected latest, got %s", service.Version())
	}

	assertStringSlice(t, service.Events(), []string{
		"PostCreated",
		"PostUpdated",
		"PostScored",
	})
}

func TestParserApplicationServiceParsesManualDependenciesWhenNoApplication(t *testing.T) {
	app := NewParserApplication()

	service, err := app.Service(`
service Gateway nginx {
  image "nginx:latest"
  exposes 80
  exposes 443
  depends_on BlogAPI
  depends_on EmbeddingService
}
`)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if service.Name() != "Gateway" {
		t.Fatalf("expected Gateway, got %s", service.Name())
	}

	if service.Kind() != services.ServiceKindNginx {
		t.Fatalf("expected nginx, got %s", service.Kind())
	}

	if service.Image() != "nginx:latest" {
		t.Fatalf("expected nginx:latest, got %s", service.Image())
	}

	assertIntSlice(t, service.ExposedPorts(), []int{80, 443})
	assertStringSlice(t, service.DependsOn(), []string{"BlogAPI", "EmbeddingService"})
}

func TestParserApplicationServiceRejectsDependsOnWithApplication(t *testing.T) {
	app := NewParserApplication()

	_, err := app.Service(`
service BlogAPI go {
  application "blog_api.arch"
  depends_on MainDB
}
`)
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestParserApplicationServiceRejectsInvalidSyntax(t *testing.T) {
	app := NewParserApplication()

	_, err := app.Service(`
service BadService go {
  exposes
}
`)
	if err == nil {
		t.Fatalf("expected error")
	}
}

func TestParserApplicationServiceRejectsInvalidPort(t *testing.T) {
	app := NewParserApplication()

	_, err := app.Service(`
service BadService go {
  exposes 0
}
`)
	if err == nil {
		t.Fatalf("expected error")
	}
}
