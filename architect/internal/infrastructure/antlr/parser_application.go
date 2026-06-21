package antlr

import (
	"fmt"

	antlr4 "github.com/antlr4-go/antlr/v4"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/endpoints"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"
	endpoint_generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/endpoint"
	object_generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/object"
)

type parserApplication struct{}

func (p *parserApplication) Object(script string) (objects.Object, error) {
	input := antlr4.NewInputStream(script)

	lexer := object_generated.NewObjectLexer(input)
	lexer.RemoveErrorListeners()

	lexerErrors := &syntaxErrorListener{
		DefaultErrorListener: antlr4.NewDefaultErrorListener(),
	}
	lexer.AddErrorListener(lexerErrors)

	tokens := antlr4.NewCommonTokenStream(lexer, antlr4.TokenDefaultChannel)

	parser := object_generated.NewObjectParser(tokens)
	parser.BuildParseTrees = true
	parser.RemoveErrorListeners()

	parserErrors := &syntaxErrorListener{
		DefaultErrorListener: antlr4.NewDefaultErrorListener(),
	}
	parser.AddErrorListener(parserErrors)

	tree := parser.Program()

	if lexerErrors.err != nil {
		return nil, lexerErrors.err
	}

	if parserErrors.err != nil {
		return nil, parserErrors.err
	}

	visitor := &objectVisitor{
		BaseObjectVisitor: &object_generated.BaseObjectVisitor{},
	}

	result := tree.Accept(visitor)

	if visitor.err != nil {
		return nil, visitor.err
	}

	obj, ok := result.(objects.Object)
	if !ok || obj == nil {
		return nil, fmt.Errorf("failed to parse object")
	}

	return obj, nil
}

func (p *parserApplication) Endpoint(script string) (endpoints.Endpoint, error) {
	input := antlr4.NewInputStream(script)

	lexer := endpoint_generated.NewEndpointLexer(input)
	lexer.RemoveErrorListeners()

	lexerErrors := &syntaxErrorListener{
		DefaultErrorListener: antlr4.NewDefaultErrorListener(),
	}
	lexer.AddErrorListener(lexerErrors)

	tokens := antlr4.NewCommonTokenStream(lexer, antlr4.TokenDefaultChannel)

	parser := endpoint_generated.NewEndpointParser(tokens)
	parser.BuildParseTrees = true
	parser.RemoveErrorListeners()

	parserErrors := &syntaxErrorListener{
		DefaultErrorListener: antlr4.NewDefaultErrorListener(),
	}
	parser.AddErrorListener(parserErrors)

	tree := parser.Program()

	if lexerErrors.err != nil {
		return nil, lexerErrors.err
	}

	if parserErrors.err != nil {
		return nil, parserErrors.err
	}

	visitor := &endpointVisitor{
		BaseEndpointVisitor: &endpoint_generated.BaseEndpointVisitor{},
	}

	result := tree.Accept(visitor)
	if visitor.err != nil {
		return nil, visitor.err
	}

	endpoint, ok := result.(endpoints.Endpoint)
	if !ok || endpoint == nil {
		return nil, fmt.Errorf("failed to parse endpoint")
	}

	return endpoint, nil
}
