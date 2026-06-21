package antlr

import (
	"fmt"

	antlr4 "github.com/antlr4-go/antlr/v4"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/applications"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/consumers"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/deployments"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/endpoints"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/services"
	application_generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/application"
	consumer_generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/consumer"
	deployment_generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/deployment"
	endpoint_generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/endpoint"
	object_generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/object"
	service_generated "github.com/steve-rodrigue/architect-lang/architect/internal/generated/service"
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

// Application parses an application script
func (p *parserApplication) Application(script string) (applications.Application, error) {
	input := antlr4.NewInputStream(script)

	lexer := application_generated.NewApplicationLexer(input)
	lexer.RemoveErrorListeners()

	lexerErrors := &syntaxErrorListener{
		DefaultErrorListener: antlr4.NewDefaultErrorListener(),
	}
	lexer.AddErrorListener(lexerErrors)

	tokens := antlr4.NewCommonTokenStream(lexer, antlr4.TokenDefaultChannel)

	parser := application_generated.NewApplicationParser(tokens)
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

	visitor := newApplicationVisitor()

	result := tree.Accept(visitor)

	if visitor.err != nil {
		return nil, visitor.err
	}

	application, ok := result.(applications.Application)
	if !ok || application == nil {
		return nil, fmt.Errorf("failed to parse application")
	}

	return application, nil
}

// Service parses a service script
func (p *parserApplication) Service(script string) (services.Service, error) {
	input := antlr4.NewInputStream(script)

	lexer := service_generated.NewServiceLexer(input)
	lexer.RemoveErrorListeners()

	lexerErrors := &syntaxErrorListener{
		DefaultErrorListener: antlr4.NewDefaultErrorListener(),
	}
	lexer.AddErrorListener(lexerErrors)

	tokens := antlr4.NewCommonTokenStream(lexer, antlr4.TokenDefaultChannel)

	parser := service_generated.NewServiceParser(tokens)
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

	visitor := newServiceVisitor()

	result := tree.Accept(visitor)

	if visitor.err != nil {
		return nil, visitor.err
	}

	service, ok := result.(services.Service)
	if !ok || service == nil {
		return nil, fmt.Errorf("failed to parse service")
	}

	return service, nil
}

// Deployment parses a deployment script
func (p *parserApplication) Deployment(script string) (deployments.Deployment, error) {
	input := antlr4.NewInputStream(script)

	lexer := deployment_generated.NewDeploymentLexer(input)
	lexer.RemoveErrorListeners()

	lexerErrors := &syntaxErrorListener{
		DefaultErrorListener: antlr4.NewDefaultErrorListener(),
	}
	lexer.AddErrorListener(lexerErrors)

	tokens := antlr4.NewCommonTokenStream(lexer, antlr4.TokenDefaultChannel)

	parser := deployment_generated.NewDeploymentParser(tokens)
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

	visitor := newDeploymentVisitor()

	result := tree.Accept(visitor)

	if visitor.err != nil {
		return nil, visitor.err
	}

	deployment, ok := result.(deployments.Deployment)
	if !ok || deployment == nil {
		return nil, fmt.Errorf("failed to parse deployment")
	}

	return deployment, nil
}
