package files

import (
	"testing"

	projects_application "github.com/steve-rodrigue/architect-lang/architect/internal/applications/projects"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/model"
	parser_antlr "github.com/steve-rodrigue/architect-lang/architect/internal/infrastructure/antlr"
)

func TestProjectApplicationParsesLoadsAndReturnsProject(t *testing.T) {
	root := t.TempDir()

	writeFile(t, root, "project.arch", `
project Stadan {
  version "0.1.0" {
    objects {
      "objects/User.arch"
      "objects/Post.arch"
    }

    services {
      "services/maindb/service.arch"
      "services/blog_api/service.arch"
    }

    deployments {
      "deployments/dev.arch"
    }
  }
}
`)

	writeFile(t, root, "0.1.0/objects/User.arch", `
object User {
  id UUID primary
  email String unique
}
`)

	writeFile(t, root, "0.1.0/objects/Post.arch", `
object Post {
  id UUID primary
  title String
  postedBy User
}
`)

	writeFile(t, root, "0.1.0/services/maindb/service.arch", `
service MainDB postgres {
  version "17"
  stores User
  stores Post
}
`)

	writeFile(t, root, "0.1.0/services/blog_api/service.arch", `
service BlogAPI go {
  exposes 8080
  application "application/application.arch"
}
`)

	writeFile(t, root, "0.1.0/services/blog_api/application/application.arch", `
application BlogAPI {
  emits rest on 8080

  endpoints {
    "endpoints/CreatePost.arch"
  }

  consumers {
  }
}
`)

	writeFile(t, root, "0.1.0/services/blog_api/application/endpoints/CreatePost.arch", `
endpoint CreatePost POST "/posts" {
  input {
    title String body
    postedBy UUID body
  }

  create post:Post {
    title = input.title
    postedBy = input.postedBy
  }

  store post to MainDB

  return post
}
`)

	writeFile(t, root, "0.1.0/deployments/dev.arch", `
deployment dev {
  vendor docker_compose

  service BlogAPI {
    replicas 1
  }

  service MainDB {
    volume "main-db-data"
  }
}
`)

	parserApplication := parser_antlr.NewParserApplication()
	resolver := model.NewResolver()
	projectApp := projects_application.NewApplication(resolver)

	fileApp := NewProjectApplication(parserApplication, projectApp)

	if err := fileApp.Parse(root); err != nil {
		t.Fatalf("expected parse no error, got %v", err)
	}

	if err := fileApp.Load(); err != nil {
		t.Fatalf("expected load no error, got %v", err)
	}

	project, err := fileApp.Project()
	if err != nil {
		t.Fatalf("expected project no error, got %v", err)
	}

	if project.Name() != "Stadan" {
		t.Fatalf("expected project Stadan, got %s", project.Name())
	}

	if len(project.Versions()) != 1 {
		t.Fatalf("expected 1 version, got %d", len(project.Versions()))
	}

	version := project.Versions()[0]

	if version.Number() != "0.1.0" {
		t.Fatalf("expected version 0.1.0, got %s", version.Number())
	}

	if version.ObjectByName("User") == nil {
		t.Fatalf("expected User object")
	}

	if version.ObjectByName("Post") == nil {
		t.Fatalf("expected Post object")
	}

	mainDB := version.ServiceByName("MainDB")
	if mainDB == nil {
		t.Fatalf("expected MainDB service")
	}

	if len(mainDB.StoredObjects()) != 2 {
		t.Fatalf("expected MainDB to store 2 objects, got %d", len(mainDB.StoredObjects()))
	}

	blogAPI := version.ServiceByName("BlogAPI")
	if blogAPI == nil {
		t.Fatalf("expected BlogAPI service")
	}

	if !blogAPI.HasApplication() {
		t.Fatalf("expected BlogAPI application")
	}

	if len(blogAPI.Application().Endpoints()) != 1 {
		t.Fatalf("expected 1 endpoint, got %d", len(blogAPI.Application().Endpoints()))
	}

	deployment := version.DeploymentByName("dev")
	if deployment == nil {
		t.Fatalf("expected dev deployment")
	}

	if len(deployment.Services()) != 2 {
		t.Fatalf("expected 2 deployed services, got %d", len(deployment.Services()))
	}
}

func TestProjectApplicationParseRejectsNilParser(t *testing.T) {
	app := NewProjectApplication(nil, nil)

	if err := app.Parse(t.TempDir()); err == nil {
		t.Fatalf("expected error")
	}
}

func TestProjectApplicationParseRejectsEmptyPath(t *testing.T) {
	app := NewProjectApplication(newParserApplication(t), nil)

	if err := app.Parse(""); err == nil {
		t.Fatalf("expected error")
	}
}

func TestProjectApplicationParseRejectsMissingProjectFile(t *testing.T) {
	app := NewProjectApplication(newParserApplication(t), nil)

	if err := app.Parse(t.TempDir()); err == nil {
		t.Fatalf("expected error")
	}
}

func TestProjectApplicationParseRejectsInvalidProjectFile(t *testing.T) {
	root := t.TempDir()
	writeFile(t, root, "project.arch", `project`)

	app := NewProjectApplication(newParserApplication(t), nil)

	if err := app.Parse(root); err == nil {
		t.Fatalf("expected error")
	}
}

func TestProjectApplicationParseRejectsMissingObjectFile(t *testing.T) {
	root := t.TempDir()

	writeFile(t, root, "project.arch", `
project Stadan {
  version "0.1.0" {
    objects {
      "objects/Missing.arch"
    }
  }
}
`)

	app := NewProjectApplication(newParserApplication(t), nil)

	if err := app.Parse(root); err == nil {
		t.Fatalf("expected error")
	}
}

func TestProjectApplicationParseRejectsInvalidObjectFile(t *testing.T) {
	root := t.TempDir()

	writeFile(t, root, "project.arch", `
project Stadan {
  version "0.1.0" {
    objects {
      "objects/User.arch"
    }
  }
}
`)

	writeFile(t, root, "0.1.0/objects/User.arch", `object User { id }`)

	app := NewProjectApplication(newParserApplication(t), nil)

	if err := app.Parse(root); err == nil {
		t.Fatalf("expected error")
	}
}

func TestProjectApplicationParseRejectsMissingServiceFile(t *testing.T) {
	root := t.TempDir()

	writeFile(t, root, "project.arch", `
project Stadan {
  version "0.1.0" {
    services {
      "services/missing/service.arch"
    }
  }
}
`)

	app := NewProjectApplication(newParserApplication(t), nil)

	if err := app.Parse(root); err == nil {
		t.Fatalf("expected error")
	}
}

func TestProjectApplicationParseRejectsInvalidServiceFile(t *testing.T) {
	root := t.TempDir()

	writeFile(t, root, "project.arch", `
project Stadan {
  version "0.1.0" {
    services {
      "services/api/service.arch"
    }
  }
}
`)

	writeFile(t, root, "0.1.0/services/api/service.arch", `service BlogAPI go { exposes }`)

	app := NewProjectApplication(newParserApplication(t), nil)

	if err := app.Parse(root); err == nil {
		t.Fatalf("expected error")
	}
}

func TestProjectApplicationParseRejectsMissingApplicationFile(t *testing.T) {
	root := t.TempDir()

	writeFile(t, root, "project.arch", `
project Stadan {
  version "0.1.0" {
    services {
      "services/api/service.arch"
    }
  }
}
`)

	writeFile(t, root, "0.1.0/services/api/service.arch", `
service BlogAPI go {
  application "application/application.arch"
}
`)

	app := NewProjectApplication(newParserApplication(t), nil)

	if err := app.Parse(root); err == nil {
		t.Fatalf("expected error")
	}
}

func TestProjectApplicationParseRejectsInvalidApplicationFile(t *testing.T) {
	root := t.TempDir()

	writeFile(t, root, "project.arch", `
project Stadan {
  version "0.1.0" {
    services {
      "services/api/service.arch"
    }
  }
}
`)

	writeFile(t, root, "0.1.0/services/api/service.arch", `
service BlogAPI go {
  application "application/application.arch"
}
`)

	writeFile(t, root, "0.1.0/services/api/application/application.arch", `application BlogAPI { emits rest on }`)

	app := NewProjectApplication(newParserApplication(t), nil)

	if err := app.Parse(root); err == nil {
		t.Fatalf("expected error")
	}
}

func TestProjectApplicationParseRejectsMissingEndpointFile(t *testing.T) {
	root := t.TempDir()

	writeProjectWithApplication(t, root, `
endpoints {
  "endpoints/Missing.arch"
}

consumers {
}
`)

	app := NewProjectApplication(newParserApplication(t), nil)

	if err := app.Parse(root); err == nil {
		t.Fatalf("expected error")
	}
}

func TestProjectApplicationParseRejectsInvalidEndpointFile(t *testing.T) {
	root := t.TempDir()

	writeProjectWithApplication(t, root, `
endpoints {
  "endpoints/CreatePost.arch"
}

consumers {
}
`)

	writeFile(t, root, "0.1.0/services/api/application/endpoints/CreatePost.arch", `endpoint Bad GET "/bad" { return }`)

	app := NewProjectApplication(newParserApplication(t), nil)

	if err := app.Parse(root); err == nil {
		t.Fatalf("expected error")
	}
}

func TestProjectApplicationParseRejectsMissingConsumerFile(t *testing.T) {
	root := t.TempDir()

	writeProjectWithApplication(t, root, `
endpoints {
}

consumers {
  "consumers/Missing.arch"
}
`)

	app := NewProjectApplication(newParserApplication(t), nil)

	if err := app.Parse(root); err == nil {
		t.Fatalf("expected error")
	}
}

func TestProjectApplicationParseRejectsInvalidConsumerFile(t *testing.T) {
	root := t.TempDir()

	writeProjectWithApplication(t, root, `
endpoints {
}

consumers {
  "consumers/PostCreated.arch"
}
`)

	writeFile(t, root, "0.1.0/services/api/application/consumers/PostCreated.arch", `consumes PostCreated { return }`)

	app := NewProjectApplication(newParserApplication(t), nil)

	if err := app.Parse(root); err == nil {
		t.Fatalf("expected error")
	}
}

func TestProjectApplicationParseRejectsMissingDeploymentFile(t *testing.T) {
	root := t.TempDir()

	writeFile(t, root, "project.arch", `
project Stadan {
  version "0.1.0" {
    deployments {
      "deployments/missing.arch"
    }
  }
}
`)

	app := NewProjectApplication(newParserApplication(t), nil)

	if err := app.Parse(root); err == nil {
		t.Fatalf("expected error")
	}
}

func TestProjectApplicationParseRejectsInvalidDeploymentFile(t *testing.T) {
	root := t.TempDir()

	writeFile(t, root, "project.arch", `
project Stadan {
  version "0.1.0" {
    deployments {
      "deployments/dev.arch"
    }
  }
}
`)

	writeFile(t, root, "0.1.0/deployments/dev.arch", `deployment dev { vendor }`)

	app := NewProjectApplication(newParserApplication(t), nil)

	if err := app.Parse(root); err == nil {
		t.Fatalf("expected error")
	}
}

func TestProjectApplicationLoadRejectsNilProjectApplication(t *testing.T) {
	root := t.TempDir()
	writeMinimalProject(t, root)

	app := NewProjectApplication(newParserApplication(t), nil)

	if err := app.Parse(root); err != nil {
		t.Fatalf("expected parse no error, got %v", err)
	}

	if err := app.Load(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestProjectApplicationLoadRejectsUnparsedSource(t *testing.T) {
	app := NewProjectApplication(newParserApplication(t), newProjectApplication(t))

	if err := app.Load(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestProjectApplicationProjectRejectsNilProjectApplication(t *testing.T) {
	app := NewProjectApplication(newParserApplication(t), nil)

	if _, err := app.Project(); err == nil {
		t.Fatalf("expected error")
	}
}

func TestProjectApplicationProjectReturnsDelegatedErrorWhenNotLoaded(t *testing.T) {
	app := NewProjectApplication(newParserApplication(t), newProjectApplication(t))

	if _, err := app.Project(); err == nil {
		t.Fatalf("expected error")
	}
}
