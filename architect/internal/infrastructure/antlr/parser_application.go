package antlr

import (
	"fmt"

	antlr4 "github.com/antlr4-go/antlr/v4"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/consumers"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/endpoints"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"
	consumer_generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/consumer"
	endpoint_generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/endpoint"
	object_generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/object"
)

type parserApplication struct{}

// Object parses an object script
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

	visitor := newObjectVisitor()

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

// Endpoint parses an endpoint script
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

	visitor := newEndpointVisitor()

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

// Consumer parses a consumer script
func (p *parserApplication) Consumer(script string) (consumers.Consumer, error) {
	input := antlr4.NewInputStream(script)

	lexer := consumer_generated.NewConsumerLexer(input)
	lexer.RemoveErrorListeners()

	lexerErrors := &syntaxErrorListener{
		DefaultErrorListener: antlr4.NewDefaultErrorListener(),
	}
	lexer.AddErrorListener(lexerErrors)

	tokens := antlr4.NewCommonTokenStream(lexer, antlr4.TokenDefaultChannel)

	parser := consumer_generated.NewConsumerParser(tokens)
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

	visitor := newConsumerVisitor()

	result := tree.Accept(visitor)

	if visitor.err != nil {
		return nil, visitor.err
	}

	consumer, ok := result.(consumers.Consumer)
	if !ok || consumer == nil {
		return nil, fmt.Errorf("failed to parse consumer")
	}

	return consumer, nil
}
