package files

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/steve-rodrigue/architect-lang/architect/internal/applications/parsers"
	projects_application "github.com/steve-rodrigue/architect-lang/architect/internal/applications/projects"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/model"
	"github.com/steve-rodrigue/architect-lang/architect/internal/infrastructure/antlr"
)

func writeFile(t *testing.T, root string, relativePath string, content string) {
	t.Helper()

	path := filepath.Join(root, relativePath)

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatalf("mkdir %q: %v", filepath.Dir(path), err)
	}

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("write %q: %v", path, err)
	}
}

func writeMinimalProject(t *testing.T, root string) {
	t.Helper()

	writeFile(t, root, "project.arch", `
project Stadan {
  version "0.1.0" {
  }
}
`)
}

func writeProjectWithApplication(t *testing.T, root string, applicationBlocks string) {
	t.Helper()

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

	writeFile(t, root, "0.1.0/services/api/application/application.arch", `
application BlogAPI {
`+applicationBlocks+`
}
`)
}

func newParserApplication(t *testing.T) parsers.Application {
	t.Helper()
	return antlr.NewParserApplication()
}

func newProjectApplication(t *testing.T) projects_application.Application {
	t.Helper()
	return projects_application.NewApplication(model.NewResolver())
}
