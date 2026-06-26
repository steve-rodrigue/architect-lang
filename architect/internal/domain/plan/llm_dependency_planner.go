package plan

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/llm"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/model"
)

type llmDependencyPlanner struct {
	repository llm.Repository
	model      string
}

func (p *llmDependencyPlanner) PlanServiceDependencies(
	version Version,
	service model.Service,
) ([]Dependency, error) {
	if p.repository == nil {
		return nil, fmt.Errorf("llm repository is required")
	}

	if p.model == "" {
		return nil, fmt.Errorf("llm model is required")
	}

	if version == nil {
		return nil, fmt.Errorf("version is required")
	}

	if service == nil {
		return nil, fmt.Errorf("service is required")
	}

	systemMessage, err := llm.NewChatMessageBuilder().
		Role(llm.ChatRoleSystem).
		Content(dependencyPlannerSystemPrompt()).
		Build()
	if err != nil {
		return nil, err
	}

	userMessage, err := llm.NewChatMessageBuilder().
		Role(llm.ChatRoleUser).
		Content(dependencyPlannerUserPrompt(service)).
		Build()
	if err != nil {
		return nil, err
	}

	request, err := llm.NewChatRequestBuilder().
		Model(p.model).
		AddMessage(systemMessage).
		AddMessage(userMessage).
		Temperature(0.1).
		TopP(0.9).
		MaxTokens(4096).
		Build()
	if err != nil {
		return nil, err
	}

	response, err := p.repository.Chat(context.Background(), request)
	if err != nil {
		return nil, err
	}

	if response == nil || response.Message() == nil {
		return nil, fmt.Errorf("llm dependency planner returned empty response")
	}

	specs, err := parseDependencyPlannerResponse(response.Message().Content())
	if err != nil {
		return nil, err
	}

	dependencies := make([]Dependency, 0, len(specs))

	for _, spec := range specs {
		dependency, err := NewDependencyBuilder().
			Version(version).
			Kind(DependencyKind(spec.Kind)).
			Scope(DependencyScope(spec.Scope)).
			Ecosystem(DependencyEcosystem(spec.Ecosystem)).
			Name(spec.Name).
			DepVersion(spec.Version).
			Source(spec.Source).
			Build()
		if err != nil {
			return nil, err
		}

		dependencies = append(dependencies, dependency)
	}

	return dependencies, nil
}

type dependencyPlannerResponse struct {
	Dependencies []dependencyPlannerDependency `json:"dependencies"`
}

type dependencyPlannerDependency struct {
	Kind      string `json:"kind"`
	Scope     string `json:"scope"`
	Ecosystem string `json:"ecosystem"`
	Name      string `json:"name"`
	Version   string `json:"version"`
	Source    string `json:"source"`
}

func parseDependencyPlannerResponse(raw string) ([]dependencyPlannerDependency, error) {
	raw = strings.TrimSpace(raw)
	raw = strings.TrimPrefix(raw, "```json")
	raw = strings.TrimPrefix(raw, "```")
	raw = strings.TrimSuffix(raw, "```")
	raw = strings.TrimSpace(raw)

	var response dependencyPlannerResponse
	if err := json.Unmarshal([]byte(raw), &response); err != nil {
		return nil, fmt.Errorf("decode dependency planner response: %w", err)
	}

	return response.Dependencies, nil
}

func dependencyPlannerSystemPrompt() string {
	return fmt.Sprintf(`You are Architect Lang's dependency planner.

Your task is to infer every dependency required by a service.

Return JSON only.

Schema:
{
  "dependencies": [
    {
      "kind": "%s",
      "scope": "%s",
      "ecosystem": "%s",
      "name": "dependency name",
      "version": "dependency version",
      "source": "where it comes from"
    }
  ]
}

Allowed values:
- kind: %s
- scope: %s
- ecosystem: %s

Rules:
- Return valid JSON only.
- Never return markdown.
- Never explain your answer.
- Never invent dependencies that are not required.

Classification rules:
- Every service runs inside Docker. Therefore every runtime service MUST have one dependency with:
  - kind=image
  - ecosystem=docker
  - scope=runtime
- Imported code packages are kind=module.
- Build utilities (protoc, buf, linters, generators, compilers, etc.) are kind=tool.
- Operating-system packages are kind=environment.
- The ecosystem must correspond to the dependency manager responsible for the dependency.
- Prefer explicit versions provided by the service.
- Prefer official sources whenever possible.
- Include every required dependency exactly once.
- If a version cannot be determined, use "latest".`,
		DependencyKindsString(),
		DependencyScopesString(),
		DependencyEcosystemsString(),
		DependencyKindsString(),
		DependencyScopesString(),
		DependencyEcosystemsString(),
	)
}

func dependencyPlannerUserPrompt(service model.Service) string {
	builder := strings.Builder{}

	builder.WriteString("Plan dependencies for this resolved service.\n\n")
	builder.WriteString("Use only these dependency options:\n")
	builder.WriteString(fmt.Sprintf("- kinds: %s\n", DependencyKindsString()))
	builder.WriteString(fmt.Sprintf("- scopes: %s\n", DependencyScopesString()))
	builder.WriteString(fmt.Sprintf("- ecosystems: %s\n\n", DependencyEcosystemsString()))

	builder.WriteString("Service:\n")
	builder.WriteString(fmt.Sprintf("- name: %s\n", service.Name()))
	builder.WriteString(fmt.Sprintf("- kind: %s\n", service.Kind()))
	builder.WriteString(fmt.Sprintf("- version: %s\n", service.Version()))
	builder.WriteString(fmt.Sprintf("- image: %s\n", service.Image()))

	if service.HasApplication() {
		application := service.Application()

		builder.WriteString("\nApplication:\n")
		builder.WriteString(fmt.Sprintf("- name: %s\n", application.Name()))

		builder.WriteString("- ports:\n")
		for _, port := range application.Ports() {
			builder.WriteString(fmt.Sprintf("  - %s %s %d\n", port.Direction(), port.Kind(), port.Number()))
		}

		builder.WriteString("- endpoints:\n")
		for _, endpoint := range application.Endpoints() {
			builder.WriteString(fmt.Sprintf("  - %s %s %s\n", endpoint.Name(), endpoint.Method(), endpoint.Path()))
		}

		builder.WriteString("- consumers:\n")
		for _, consumer := range application.Consumers() {
			if consumer.Event() != nil {
				builder.WriteString(fmt.Sprintf("  - event: %s\n", consumer.Event().Name()))
			}
		}
	}

	builder.WriteString("\nService dependencies:\n")
	for _, dependency := range service.Dependencies() {
		builder.WriteString(fmt.Sprintf("- %s kind=%s version=%s image=%s\n",
			dependency.Name(),
			dependency.Kind(),
			dependency.Version(),
			dependency.Image(),
		))
	}

	builder.WriteString("\nStored objects:\n")
	for _, object := range service.StoredObjects() {
		builder.WriteString(fmt.Sprintf("- %s\n", object.Name()))
	}

	builder.WriteString("\nEvents:\n")
	for _, event := range service.Events() {
		builder.WriteString(fmt.Sprintf("- %s\n", event.Name()))
	}

	return builder.String()
}
