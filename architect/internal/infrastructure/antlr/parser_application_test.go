package antlr

import (
	"testing"

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
