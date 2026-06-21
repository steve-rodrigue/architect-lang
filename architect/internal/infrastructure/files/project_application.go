package files

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/steve-rodrigue/architect-lang/architect/internal/applications/parsers"
	projects_application "github.com/steve-rodrigue/architect-lang/architect/internal/applications/projects"
	ast_applications "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/applications"
	ast_consumers "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/consumers"
	ast_deployments "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/deployments"
	ast_endpoints "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/endpoints"
	ast_objects "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/objects"
	ast_projects "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/projects"
	ast_services "github.com/steve-rodrigue/architect-lang/architect/internal/domain/ast/services"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/model"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/model/sources"
)

type projectApplication struct {
	parserApplication  parsers.Application
	projectApplication projects_application.Application
	source             sources.Project
}

func (a *projectApplication) Parse(relativePath string) error {
	if a.parserApplication == nil {
		return fmt.Errorf("parser application is required")
	}

	if relativePath == "" {
		return fmt.Errorf("relative path is required")
	}

	projectFile := filepath.Join(relativePath, "project.arch")

	projectScript, err := readFile(projectFile)
	if err != nil {
		return err
	}

	projectAST, err := a.parserApplication.Project(projectScript)
	if err != nil {
		return fmt.Errorf("parse project %q: %w", projectFile, err)
	}

	projectBuilder := sources.NewProjectBuilder().
		Project(projectAST)

	for _, versionAST := range projectAST.Versions() {
		versionPath := filepath.Join(relativePath, versionAST.Number())

		versionSource, err := a.parseVersion(versionPath, versionAST)
		if err != nil {
			return err
		}

		projectBuilder.AddVersion(versionSource)
	}

	source, err := projectBuilder.Build()
	if err != nil {
		return err
	}

	a.source = source

	return nil
}

func (a *projectApplication) Load() error {
	if a.projectApplication == nil {
		return fmt.Errorf("project application is required")
	}

	if a.source == nil {
		return fmt.Errorf("project source is not parsed")
	}

	return a.projectApplication.Load(a.source)
}

func (a *projectApplication) Project() (model.Project, error) {
	if a.projectApplication == nil {
		return nil, fmt.Errorf("project application is required")
	}

	return a.projectApplication.Project()
}

func (a *projectApplication) parseVersion(
	versionPath string,
	versionAST ast_projects.Version,
) (sources.Version, error) {
	versionBuilder := sources.NewVersionBuilder().
		Version(versionAST)

	for _, objectFile := range versionAST.ObjectFiles() {
		objectAST, err := a.parseObject(filepath.Join(versionPath, objectFile))
		if err != nil {
			return nil, err
		}

		versionBuilder.AddObject(objectAST)
	}

	for _, serviceFile := range versionAST.ServiceFiles() {
		serviceSource, err := a.parseService(versionPath, serviceFile)
		if err != nil {
			return nil, err
		}

		versionBuilder.AddService(serviceSource)
	}

	for _, deploymentFile := range versionAST.DeploymentFiles() {
		deploymentAST, err := a.parseDeployment(filepath.Join(versionPath, deploymentFile))
		if err != nil {
			return nil, err
		}

		versionBuilder.AddDeployment(deploymentAST)
	}

	return versionBuilder.Build()
}

func (a *projectApplication) parseService(
	versionPath string,
	serviceFile string,
) (sources.Service, error) {
	servicePath := filepath.Join(versionPath, serviceFile)

	serviceAST, err := a.parseServiceFile(servicePath)
	if err != nil {
		return nil, err
	}

	var applicationAST ast_applications.Application
	endpointList := make([]ast_endpoints.Endpoint, 0)
	consumerList := make([]ast_consumers.Consumer, 0)

	if serviceAST.ApplicationFile() != "" {
		serviceDirectory := filepath.Dir(servicePath)
		applicationPath := filepath.Join(serviceDirectory, serviceAST.ApplicationFile())

		applicationAST, err = a.parseApplication(applicationPath)
		if err != nil {
			return nil, err
		}

		applicationDirectory := filepath.Dir(applicationPath)

		for _, endpointFile := range applicationAST.EndpointFiles() {
			endpointAST, err := a.parseEndpoint(filepath.Join(applicationDirectory, endpointFile))
			if err != nil {
				return nil, err
			}

			endpointList = append(endpointList, endpointAST)
		}

		for _, consumerFile := range applicationAST.ConsumerFiles() {
			consumerAST, err := a.parseConsumer(filepath.Join(applicationDirectory, consumerFile))
			if err != nil {
				return nil, err
			}

			consumerList = append(consumerList, consumerAST)
		}
	}

	return sources.FromService(
		serviceAST,
		applicationAST,
		endpointList,
		consumerList,
	)
}

func (a *projectApplication) parseObject(path string) (ast_objects.Object, error) {
	script, err := readFile(path)
	if err != nil {
		return nil, err
	}

	object, err := a.parserApplication.Object(script)
	if err != nil {
		return nil, fmt.Errorf("parse object %q: %w", path, err)
	}

	return object, nil
}

func (a *projectApplication) parseServiceFile(path string) (ast_services.Service, error) {
	script, err := readFile(path)
	if err != nil {
		return nil, err
	}

	service, err := a.parserApplication.Service(script)
	if err != nil {
		return nil, fmt.Errorf("parse service %q: %w", path, err)
	}

	return service, nil
}

func (a *projectApplication) parseApplication(path string) (ast_applications.Application, error) {
	script, err := readFile(path)
	if err != nil {
		return nil, err
	}

	application, err := a.parserApplication.Application(script)
	if err != nil {
		return nil, fmt.Errorf("parse application %q: %w", path, err)
	}

	return application, nil
}

func (a *projectApplication) parseEndpoint(path string) (ast_endpoints.Endpoint, error) {
	script, err := readFile(path)
	if err != nil {
		return nil, err
	}

	endpoint, err := a.parserApplication.Endpoint(script)
	if err != nil {
		return nil, fmt.Errorf("parse endpoint %q: %w", path, err)
	}

	return endpoint, nil
}

func (a *projectApplication) parseConsumer(path string) (ast_consumers.Consumer, error) {
	script, err := readFile(path)
	if err != nil {
		return nil, err
	}

	consumer, err := a.parserApplication.Consumer(script)
	if err != nil {
		return nil, fmt.Errorf("parse consumer %q: %w", path, err)
	}

	return consumer, nil
}

func (a *projectApplication) parseDeployment(path string) (ast_deployments.Deployment, error) {
	script, err := readFile(path)
	if err != nil {
		return nil, err
	}

	deployment, err := a.parserApplication.Deployment(script)
	if err != nil {
		return nil, fmt.Errorf("parse deployment %q: %w", path, err)
	}

	return deployment, nil
}

func readFile(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read file %q: %w", path, err)
	}

	return string(content), nil
}
