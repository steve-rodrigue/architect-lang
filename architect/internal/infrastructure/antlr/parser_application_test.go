package antlr

import (
	"testing"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/endpoints"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"
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

	if field.Type().DefaultValue().Kind() != objects.ValueString {
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

	assertNumberConstraint(t, obj.Fields()[0], objects.NumberConstraintOneOrMore, false, "", false, "")
	assertNumberConstraint(t, obj.Fields()[1], objects.NumberConstraintZeroOrMore, false, "", false, "")
	assertNumberConstraint(t, obj.Fields()[2], objects.NumberConstraintRange, true, "0", true, "23")
	assertNumberConstraint(t, obj.Fields()[3], objects.NumberConstraintRange, true, "3", false, "")
	assertNumberConstraint(t, obj.Fields()[4], objects.NumberConstraintRange, true, "34", true, "38")

	optionalDefault := obj.Fields()[5]
	if !optionalDefault.Type().IsOptional() {
		t.Fatalf("expected optionalDefault to be optional")
	}

	if !optionalDefault.Type().HasDefaultValue() {
		t.Fatalf("expected optionalDefault to have default")
	}

	if optionalDefault.Type().DefaultValue().Kind() != objects.ValueInt {
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
	expectedKind objects.NumberConstraintKind,
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

	if endpoint.Input() == nil {
		t.Fatalf("expected input")
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

	if id.Sources().Sources()[0].Kind() != endpoints.InputSourcePath {
		t.Fatalf("expected first id source path")
	}

	if !id.Sources().Sources()[0].IsOptional() {
		t.Fatalf("expected path source to be optional")
	}

	if id.Sources().Sources()[1].Kind() != endpoints.InputSourceBody {
		t.Fatalf("expected second id source body")
	}

	if !id.Sources().Sources()[1].IsOptional() {
		t.Fatalf("expected body source to be optional")
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

	if actions[0].Kind() != endpoints.ActionKindFetch {
		t.Fatalf("expected action 0 fetch")
	}

	fetch := actions[0].(endpoints.FetchAction)
	if fetch.Variable().Name() != "post" {
		t.Fatalf("expected fetch variable post")
	}

	if fetch.Variable().Type().Name() != "Post" {
		t.Fatalf("expected fetch variable type Post")
	}

	if fetch.Source() != "MainDB" {
		t.Fatalf("expected fetch source MainDB")
	}

	if fetch.Condition().Operator() != endpoints.ComparatorEqual {
		t.Fatalf("expected fetch condition ==")
	}

	if actions[1].Kind() != endpoints.ActionKindCreate {
		t.Fatalf("expected action 1 create")
	}

	create := actions[1].(endpoints.CreateAction)
	if create.Variable().Name() != "history" {
		t.Fatalf("expected create variable history")
	}

	if create.Variable().Type().Name() != "PostHistory" {
		t.Fatalf("expected create type PostHistory")
	}

	if len(create.Assignments()) != 7 {
		t.Fatalf("expected 7 history assignments, got %d", len(create.Assignments()))
	}

	if actions[2].Kind() != endpoints.ActionKindStore {
		t.Fatalf("expected action 2 store")
	}

	storeHistory := actions[2].(endpoints.StoreAction)
	if storeHistory.VariableName() != "history" || storeHistory.Destination() != "MainDB" {
		t.Fatalf("expected store history to MainDB")
	}

	if actions[3].Kind() != endpoints.ActionKindUpdate {
		t.Fatalf("expected action 3 update")
	}

	update := actions[3].(endpoints.UpdateAction)
	if update.VariableName() != "post" {
		t.Fatalf("expected update post")
	}

	if len(update.Assignments()) != 3 {
		t.Fatalf("expected 3 update assignments, got %d", len(update.Assignments()))
	}

	if actions[6].Kind() != endpoints.ActionKindEmit {
		t.Fatalf("expected action 6 emit")
	}

	emit := actions[6].(endpoints.EmitAction)
	if emit.Variable().Name() != "event" {
		t.Fatalf("expected emit variable event")
	}

	if emit.Variable().Type().Name() != "PostUpdated" {
		t.Fatalf("expected emit type PostUpdated")
	}

	if len(emit.Assignments()) != 2 {
		t.Fatalf("expected 2 emit assignments")
	}

	if actions[7].Kind() != endpoints.ActionKindReturn {
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

	create := actions[0].(endpoints.CreateAction)
	assignments := create.Assignments()

	assertAssignmentSelector(t, assignments[0], "title", []string{"input", "title"})
	assertAssignmentSelector(t, assignments[1], "content", []string{"input", "content"})
	assertAssignmentFunctionCall(t, assignments[3], "createdAt", "now", 0)
	assertAssignmentFunctionCall(t, assignments[4], "slug", "slugify", 1)
	assertAssignmentFunctionCall(t, assignments[5], "score", "calculateScore", 5)

	scoreCall := assignments[5].Value().(endpoints.FunctionCallExpression)
	assertLiteral(t, scoreCall.Arguments()[1], objects.ValueInt, "12")
	assertLiteral(t, scoreCall.Arguments()[2], objects.ValueFloat, "12.5")
	assertLiteral(t, scoreCall.Arguments()[3], objects.ValueBool, "true")
	assertLiteral(t, scoreCall.Arguments()[4], objects.ValueBool, "false")

	storeOne := actions[1].(endpoints.StoreAction)
	if storeOne.VariableName() != "post" || storeOne.Destination() != "MainDB" {
		t.Fatalf("expected store post to MainDB")
	}

	storeTwo := actions[2].(endpoints.StoreAction)
	if storeTwo.VariableName() != "post" || storeTwo.Destination() != "GraphDB" {
		t.Fatalf("expected store post to GraphDB")
	}

	emit := actions[3].(endpoints.EmitAction)
	if emit.Variable().Name() != "event" || emit.Variable().Type().Name() != "PostCreated" {
		t.Fatalf("expected emit event:PostCreated")
	}

	ret := actions[4].(endpoints.ReturnAction)
	if ret.Expression().Kind() != endpoints.ExpressionKindIdentifier {
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
		expected endpoints.Comparator
	}{
		{"EqEndpoint", "==", endpoints.ComparatorEqual},
		{"NeqEndpoint", "!=", endpoints.ComparatorNotEqual},
		{"LtEndpoint", "<", endpoints.ComparatorLessThan},
		{"LteEndpoint", "<=", endpoints.ComparatorLessThanOrEqual},
		{"GtEndpoint", ">", endpoints.ComparatorGreaterThan},
		{"GteEndpoint", ">=", endpoints.ComparatorGreaterThanOrEqual},
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

			fetch := endpoint.Actions()[0].(endpoints.FetchAction)
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
	if limitConstraint.Kind() != objects.NumberConstraintRange {
		t.Fatalf("expected range constraint")
	}
	if !limitConstraint.HasMin() || limitConstraint.Min().Raw() != "1" {
		t.Fatalf("expected min 1")
	}
	if !limitConstraint.HasMax() || limitConstraint.Max().Raw() != "100" {
		t.Fatalf("expected max 100")
	}

	offsetConstraint := fields[1].Type().NumberConstraint()
	if offsetConstraint.Kind() != objects.NumberConstraintZeroOrMore {
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

	fetch := endpoint.Actions()[0].(endpoints.FetchAction)
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
