package antlr

import (
	"testing"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/common"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/endpoints"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/workflows"
)

func assertAssignmentSelector(t *testing.T, assignment workflows.Assignment, name string, parts []string) {
	t.Helper()

	if assignment.Name() != name {
		t.Fatalf("expected assignment %s, got %s", name, assignment.Name())
	}

	selector, ok := assignment.Value().(workflows.SelectorExpression)
	if !ok {
		t.Fatalf("expected selector expression for %s", name)
	}

	actual := selector.Parts()
	if len(actual) != len(parts) {
		t.Fatalf("expected %d selector parts, got %d", len(parts), len(actual))
	}

	for i := range parts {
		if actual[i] != parts[i] {
			t.Fatalf("expected selector part %d to be %s, got %s", i, parts[i], actual[i])
		}
	}
}

func assertLiteral(t *testing.T, expression workflows.Expression, kind common.ValueKind, raw string) {
	t.Helper()

	literal, ok := expression.(workflows.LiteralExpression)
	if !ok {
		t.Fatalf("expected literal expression")
	}

	if literal.Value().Kind() != kind {
		t.Fatalf("expected literal kind %s, got %s", kind, literal.Value().Kind())
	}

	if literal.Value().Raw() != raw {
		t.Fatalf("expected literal raw %s, got %s", raw, literal.Value().Raw())
	}
}

func assertSingleSource(t *testing.T, field endpoints.InputField, kind endpoints.InputSourceKind, optional bool) {
	t.Helper()

	sources := field.Sources().Sources()

	if field.Sources().Kind() != endpoints.InputSourceRuleSingle {
		t.Fatalf("expected single source rule, got %s", field.Sources().Kind())
	}

	if len(sources) != 1 {
		t.Fatalf("expected 1 source, got %d", len(sources))
	}

	if sources[0].Kind() != kind {
		t.Fatalf("expected source %s, got %s", kind, sources[0].Kind())
	}

	if sources[0].IsOptional() != optional {
		t.Fatalf("expected optional %v, got %v", optional, sources[0].IsOptional())
	}
}

func assertAssignmentFunctionCall(t *testing.T, assignment workflows.Assignment, name string, functionName string, argCount int) {
	t.Helper()

	if assignment.Name() != name {
		t.Fatalf("expected assignment %s, got %s", name, assignment.Name())
	}

	call, ok := assignment.Value().(workflows.FunctionCallExpression)
	if !ok {
		t.Fatalf("expected function call expression for %s", name)
	}

	if call.Name() != functionName {
		t.Fatalf("expected function %s, got %s", functionName, call.Name())
	}

	if len(call.Arguments()) != argCount {
		t.Fatalf("expected %d function args, got %d", argCount, len(call.Arguments()))
	}
}

func assertIntSlice(
	t *testing.T,
	actual []int,
	expected []int,
) {
	t.Helper()

	if len(actual) != len(expected) {
		t.Fatalf("expected %d items, got %d: %v", len(expected), len(actual), actual)
	}

	for i := range expected {
		if actual[i] != expected[i] {
			t.Fatalf("expected item %d to be %d, got %d", i, expected[i], actual[i])
		}
	}
}
