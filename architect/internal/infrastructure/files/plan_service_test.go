package files

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/model"
	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/plan"
)

func TestPlanServiceSaveAllEntities(t *testing.T) {
	dir := t.TempDir()
	service := NewPlanService(dir, "project.plan.json")

	project := buildProject(t, "Stadan", "1")
	version := buildVersion(t, project, "0.1.0")
	rootSection := buildSection(t, version, plan.SectionKindVersion, "0.1.0", "")
	serviceSection := buildSection(t, version, plan.SectionKindService, "blog_api", rootSection.ID())
	dependency := buildDependency(t, version)
	fileArtifact := buildFileArtifact(t, serviceSection)
	directoryArtifact := buildDirectoryArtifact(t, serviceSection)
	rootTask := buildTask(t, serviceSection, "Generate root", nil, "", nil, nil)
	childTask := buildTask(
		t,
		serviceSection,
		"Generate child",
		buildTaskSpec(t),
		rootTask.ID(),
		[]plan.TaskID{rootTask.ID()},
		nil,
	)

	must(t, service.SaveProject(project))
	must(t, service.SaveVersion(version))
	must(t, service.SaveSection(rootSection))
	must(t, service.SaveSection(serviceSection))
	must(t, service.SaveDependency(dependency))
	must(t, service.SaveArtifact(fileArtifact))
	must(t, service.SaveArtifact(directoryArtifact))
	must(t, service.SaveTask(rootTask))
	must(t, service.SaveTask(childTask))

	data := readPlanFile(t, filepath.Join(dir, "project.plan.json"))

	if len(data.Projects) != 1 {
		t.Fatalf("expected 1 project, got %d", len(data.Projects))
	}
	if len(data.Versions) != 1 {
		t.Fatalf("expected 1 version, got %d", len(data.Versions))
	}
	if len(data.Sections) != 2 {
		t.Fatalf("expected 2 sections, got %d", len(data.Sections))
	}
	if len(data.Dependencies) != 1 {
		t.Fatalf("expected 1 dependency, got %d", len(data.Dependencies))
	}
	if len(data.Artifacts) != 2 {
		t.Fatalf("expected 2 artifacts, got %d", len(data.Artifacts))
	}
	if len(data.Tasks) != 2 {
		t.Fatalf("expected 2 tasks, got %d", len(data.Tasks))
	}

	if data.Projects[0].Name != "Stadan" {
		t.Fatalf("expected project name Stadan, got %q", data.Projects[0].Name)
	}
	if data.Artifacts[0].File == nil {
		t.Fatal("expected file artifact spec")
	}
	if data.Artifacts[1].File != nil {
		t.Fatal("expected directory artifact without file spec")
	}
	if data.Tasks[1].Spec == nil {
		t.Fatal("expected task spec")
	}
}

func TestPlanServiceUpsertsEntities(t *testing.T) {
	dir := t.TempDir()
	service := NewPlanService(dir, "project.plan.json")

	projectV1 := buildProject(t, "Stadan", "1")
	projectV2 := buildProject(t, "Stadan", "2")

	must(t, service.SaveProject(projectV1))
	must(t, service.SaveProject(projectV2))

	data := readPlanFile(t, filepath.Join(dir, "project.plan.json"))

	if len(data.Projects) != 1 {
		t.Fatalf("expected 1 project, got %d", len(data.Projects))
	}
	if data.Projects[0].SchemaVersion != "2" {
		t.Fatalf("expected schema version 2, got %q", data.Projects[0].SchemaVersion)
	}
}

func TestPlanServiceReferenceValidation(t *testing.T) {
	dir := t.TempDir()
	service := NewPlanService(dir, "project.plan.json")

	project := buildProject(t, "Stadan", "1")
	version := buildVersion(t, project, "0.1.0")
	rootSection := buildSection(t, version, plan.SectionKindVersion, "0.1.0", "")
	childSection := buildSection(t, version, plan.SectionKindService, "blog_api", rootSection.ID())
	dependency := buildDependency(t, version)
	artifact := buildDirectoryArtifact(t, childSection)
	task := buildTask(t, childSection, "Generate", nil, "", nil, nil)

	assertErrContains(t, service.SaveVersion(version), "unknown project")

	must(t, service.SaveProject(project))
	assertErrContains(t, service.SaveSection(rootSection), "unknown version")
	assertErrContains(t, service.SaveDependency(dependency), "unknown version")

	must(t, service.SaveVersion(version))
	assertErrContains(t, service.SaveSection(childSection), "unknown parent section")
	assertErrContains(t, service.SaveArtifact(artifact), "unknown section")
	assertErrContains(t, service.SaveTask(task), "unknown section")

	must(t, service.SaveSection(rootSection))
	must(t, service.SaveSection(childSection))

	parentTask := buildTask(t, childSection, "Parent", nil, "", nil, nil)
	childTask := buildTask(t, childSection, "Child", nil, parentTask.ID(), nil, nil)
	dependencyTask := buildTask(t, childSection, "Dependency", nil, "", []plan.TaskID{parentTask.ID()}, nil)
	subTaskReference := buildTask(t, childSection, "SubTaskReference", nil, "", nil, []plan.TaskID{parentTask.ID()})

	assertErrContains(t, service.SaveTask(childTask), "unknown parent task")
	assertErrContains(t, service.SaveTask(dependencyTask), "unknown dependency task")
	assertErrContains(t, service.SaveTask(subTaskReference), "unknown sub task")

	must(t, service.SaveTask(parentTask))
	must(t, service.SaveTask(childTask))
	must(t, service.SaveTask(dependencyTask))
	must(t, service.SaveTask(subTaskReference))
}

func TestPlanServiceNilErrors(t *testing.T) {
	service := NewPlanService(t.TempDir(), "project.plan.json")

	assertErrContains(t, service.SaveProject(nil), "project is required")
	assertErrContains(t, service.SaveVersion(nil), "version is required")
	assertErrContains(t, service.SaveSection(nil), "section is required")
	assertErrContains(t, service.SaveDependency(nil), "dependency is required")
	assertErrContains(t, service.SaveArtifact(nil), "artifact is required")
	assertErrContains(t, service.SaveTask(nil), "task is required")
}

func TestPlanServiceInvalidJSONReturnsError(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "project.plan.json")

	must(t, os.WriteFile(path, []byte("{invalid-json"), 0644))

	service := NewPlanService(dir, "project.plan.json")
	project := buildProject(t, "Stadan", "1")

	assertErrContains(t, service.SaveProject(project), "decode plan file")
}

func TestPlanServiceMissingPathReturnsError(t *testing.T) {
	service := NewPlanService("", "project.plan.json")
	project := buildProject(t, "Stadan", "1")

	assertErrContains(t, service.SaveProject(project), "plan file path is required")
}

func TestPlanServiceFilePathBranches(t *testing.T) {
	dir := t.TempDir()

	directoryService := &planService{path: dir, fileName: "custom.json"}
	path, err := directoryService.filePath()
	must(t, err)
	if path != filepath.Join(dir, "custom.json") {
		t.Fatalf("unexpected directory file path: %q", path)
	}

	missingDirectory := filepath.Join(t.TempDir(), "plans")
	missingDirectoryService := &planService{path: missingDirectory, fileName: "custom.json"}
	path, err = missingDirectoryService.filePath()
	must(t, err)
	if path != filepath.Join(missingDirectory, "custom.json") {
		t.Fatalf("unexpected missing directory file path: %q", path)
	}

	explicitFile := filepath.Join(t.TempDir(), "custom.plan.json")
	explicitFileService := &planService{path: explicitFile, fileName: "ignored.json"}
	path, err = explicitFileService.filePath()
	must(t, err)
	if path != explicitFile {
		t.Fatalf("unexpected explicit file path: %q", path)
	}
}

func buildProject(t *testing.T, name string, schemaVersion string) plan.Project {
	t.Helper()

	modelProject, err := model.NewProjectBuilder().
		Name(name).
		Build()
	if err != nil {
		t.Fatal(err)
	}

	project, err := plan.NewProjectBuilder().
		Model(modelProject).
		SchemaVersion(schemaVersion).
		Build()
	if err != nil {
		t.Fatal(err)
	}

	return project
}

func buildVersion(t *testing.T, project plan.Project, number string) plan.Version {
	t.Helper()

	version, err := plan.NewVersionBuilder().
		Project(project).
		Number(number).
		Build()
	if err != nil {
		t.Fatal(err)
	}

	return version
}

func buildSection(
	t *testing.T,
	version plan.Version,
	kind plan.SectionKind,
	name string,
	parentID plan.SectionID,
) plan.Section {
	t.Helper()

	builder := plan.NewSectionBuilder().
		Version(version).
		Kind(kind).
		Name(name)

	if parentID != "" {
		builder.Parent(parentID)
	}

	section, err := builder.Build()
	if err != nil {
		t.Fatal(err)
	}

	return section
}

func buildDependency(t *testing.T, version plan.Version) plan.Dependency {
	t.Helper()

	dependency, err := plan.NewDependencyBuilder().
		Version(version).
		Kind(plan.DependencyRuntime).
		Name("go").
		DepVersion("1.25").
		Source("go.dev").
		Build()
	if err != nil {
		t.Fatal(err)
	}

	return dependency
}

func buildFileArtifact(t *testing.T, section plan.Section) plan.Artifact {
	t.Helper()

	file, err := plan.NewArtifactFileBuilder().
		Kind(plan.ArtifactFileCode).
		Purpose("Implement BlogAPI.").
		Specification("Contains the HTTP server entrypoint.").
		AddConstraint("Must compile.").
		AddAcceptanceCriteria("go test passes.").
		Build()
	if err != nil {
		t.Fatal(err)
	}

	artifact, err := plan.NewArtifactBuilder().
		Section(section).
		Path("cmd/blog_api/main.go").
		File(file).
		Build()
	if err != nil {
		t.Fatal(err)
	}

	return artifact
}

func buildDirectoryArtifact(t *testing.T, section plan.Section) plan.Artifact {
	t.Helper()

	artifact, err := plan.NewArtifactBuilder().
		Section(section).
		Path("cmd/blog_api").
		Build()
	if err != nil {
		t.Fatal(err)
	}

	return artifact
}

func buildTaskSpec(t *testing.T) plan.TaskSpec {
	t.Helper()

	spec, err := plan.NewTaskSpecBuilder().
		Goal("Generate BlogAPI.").
		Instructions("Create the HTTP server.").
		AddConstraint("Do not edit unrelated files.").
		AddAcceptanceCriteria("Server builds.").
		Build()
	if err != nil {
		t.Fatal(err)
	}

	return spec
}

func buildTask(
	t *testing.T,
	section plan.Section,
	name string,
	spec plan.TaskSpec,
	parentID plan.TaskID,
	dependencyIDs []plan.TaskID,
	subTaskIDs []plan.TaskID,
) plan.Task {
	t.Helper()

	builder := plan.NewTaskBuilder().
		Section(section).
		Kind(plan.TaskKindCode).
		Action(plan.TaskActionGenerate).
		Name(name).
		Description(name + " description.")

	if spec != nil {
		builder.Spec(spec)
	}

	if parentID != "" {
		builder.ParentID(parentID)
	}

	for _, dependencyID := range dependencyIDs {
		builder.AddDependency(dependencyID)
	}

	for _, subTaskID := range subTaskIDs {
		builder.AddSubTask(subTaskID)
	}

	task, err := builder.Build()
	if err != nil {
		t.Fatal(err)
	}

	return task
}

func readPlanFile(t *testing.T, path string) planServiceFileData {
	t.Helper()

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	var data planServiceFileData
	if err := json.Unmarshal(content, &data); err != nil {
		t.Fatal(err)
	}

	return data
}

func must(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatal(err)
	}
}

func assertErrContains(t *testing.T, err error, expected string) {
	t.Helper()

	if err == nil {
		t.Fatalf("expected error containing %q", expected)
	}

	if !strings.Contains(err.Error(), expected) {
		t.Fatalf("expected error containing %q, got %q", expected, err.Error())
	}
}
