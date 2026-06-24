package plan

import (
	"context"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/model"
)

type ProjectID string
type VersionID string
type SectionID string
type DependencyID string
type ArtifactID string
type TaskID string

type DependencyKind string

const (
	DependencyRuntime DependencyKind = "runtime"
	DependencyTool    DependencyKind = "tool"
	DependencyImage   DependencyKind = "image"
	DependencyModule  DependencyKind = "module"
)

type ArtifactKind string

const (
	ArtifactKindFile      ArtifactKind = "file"
	ArtifactKindDirectory ArtifactKind = "directory"
)

type ArtifactFileKind string

const (
	ArtifactFileCode          ArtifactFileKind = "code"
	ArtifactFileTest          ArtifactFileKind = "test"
	ArtifactFileConfig        ArtifactFileKind = "config"
	ArtifactFileDocumentation ArtifactFileKind = "documentation"
)

type TaskKind string

const (
	TaskKindCode       TaskKind = "code"
	TaskKindTest       TaskKind = "test"
	TaskKindDeployment TaskKind = "deployment"
	TaskKindCommand    TaskKind = "command"
)

type TaskAction string

const (
	TaskActionGenerate TaskAction = "generate"
	TaskActionUpdate   TaskAction = "update"
	TaskActionRefactor TaskAction = "refactor"
	TaskActionDelete   TaskAction = "delete"
	TaskActionValidate TaskAction = "validate"
	TaskActionExecute  TaskAction = "execute"
)

type SectionKind string

const (
	SectionKindVersion     SectionKind = "version"
	SectionKindObject      SectionKind = "object"
	SectionKindField       SectionKind = "field"
	SectionKindService     SectionKind = "service"
	SectionKindApplication SectionKind = "application"
	SectionKindEndpoint    SectionKind = "endpoint"
	SectionKindConsumer    SectionKind = "consumer"
	SectionKindWorkflow    SectionKind = "workflow"
	SectionKindAction      SectionKind = "action"
	SectionKindDeployment  SectionKind = "deployment"
)

// NewProjectBuilder creates a new project builder
func NewProjectBuilder() ProjectBuilder {
	return &projectBuilder{}
}

// NewVersionBuilder creates a new version builder
func NewVersionBuilder() VersionBuilder {
	return &versionBuilder{}
}

// NewSectionBuilder creates a new section builder
func NewSectionBuilder() SectionBuilder {
	return &sectionBuilder{}
}

// NewDependencyBuilder creates a new dependency builder
func NewDependencyBuilder() DependencyBuilder {
	return &dependencyBuilder{}
}

// NewArtifactBuilder creates a new artifact builder
func NewArtifactBuilder() ArtifactBuilder {
	return &artifactBuilder{}
}

// NewArtifactFileBuilder creates a new artifact file builder
func NewArtifactFileBuilder() ArtifactFileBuilder {
	return &artifactFileBuilder{
		constraints:        make([]string, 0),
		acceptanceCriteria: make([]string, 0),
	}
}

// NewTaskBuilder creates a new task builder.
func NewTaskBuilder() TaskBuilder {
	return &taskBuilder{
		dependencyIDs: make([]TaskID, 0),
		subTaskIDs:    make([]TaskID, 0),
	}
}

// NewTaskSpecBuilder creates a new task specification builder.
func NewTaskSpecBuilder() TaskSpecBuilder {
	return &taskSpecBuilder{
		constraints:        make([]string, 0),
		acceptanceCriteria: make([]string, 0),
	}
}

// Planner creates a concrete generation plan from a resolved model project.
type Planner interface {
	Plan(project model.Project) (Project, error)
}

// Service represents a plan service
type Service interface {
	SaveProject(ctx context.Context, project Project) error
	SaveVersion(ctx context.Context, version Version) error
	SaveSection(ctx context.Context, section Section) error
	SaveDependency(ctx context.Context, dependency Dependency) error
	SaveArtifact(ctx context.Context, artifact Artifact) error
	SaveTask(ctx context.Context, task Task) error
}

// ProjectRepository loads project plans on demand.
type ProjectRepository interface {
	IDs(ctx context.Context) ([]ProjectID, error)
	Names(ctx context.Context) ([]string, error)

	Get(ctx context.Context, id ProjectID) (Project, error)
	GetByName(ctx context.Context, name string) (Project, error)

	CurrentByName(ctx context.Context, name string) (Project, error)
	Previous(ctx context.Context, id ProjectID) (ProjectID, error)
	History(ctx context.Context, name string) ([]ProjectID, error)

	Has(ctx context.Context, id ProjectID) (bool, error)
}

// ProjectBuilder represents a project builder.
type ProjectBuilder interface {
	Model(projectModel model.Project) ProjectBuilder
	SchemaVersion(schemaVersion string) ProjectBuilder
	Build() (Project, error)
}

// Project represents one planned project.
type Project interface {
	ID() ProjectID
	Model() model.Project
	SchemaVersion() string
}

// VersionRepository loads project versions on demand.
type VersionRepository interface {
	IDsByProject(ctx context.Context, projectID ProjectID) ([]VersionID, error)
	Get(ctx context.Context, id VersionID) (Version, error)
	GetByNumber(ctx context.Context, projectID ProjectID, number string) (Version, error)
	Has(ctx context.Context, id VersionID) (bool, error)
}

// VersionBuilder represents a version builder.
type VersionBuilder interface {
	Project(project Project) VersionBuilder
	Number(number string) VersionBuilder
	Build() (Version, error)
}

// Version represents one planned project version.
type Version interface {
	ID() VersionID
	Project() Project
	Number() string
}

// SectionRepository loads sections on demand.
type SectionRepository interface {
	IDsByVersion(ctx context.Context, versionID VersionID) ([]SectionID, error)
	Get(ctx context.Context, id SectionID) (Section, error)
	Children(ctx context.Context, parentID SectionID) ([]SectionID, error)
	Has(ctx context.Context, id SectionID) (bool, error)
}

// SectionBuilder represents a section builder.
type SectionBuilder interface {
	Version(version Version) SectionBuilder
	Kind(kind SectionKind) SectionBuilder
	Name(name string) SectionBuilder
	Parent(parent SectionID) SectionBuilder
	Build() (Section, error)
}

// Section represents a deterministic location in the project graph.
type Section interface {
	ID() SectionID
	Version() Version

	Kind() SectionKind
	Name() string
	Path() string

	ParentID() SectionID
	HasParentID() bool
}

// DependencyRepository loads dependencies on demand.
type DependencyRepository interface {
	IDsByProject(ctx context.Context, projectID ProjectID) ([]DependencyID, error)
	IDsByVersion(ctx context.Context, versionID VersionID) ([]DependencyID, error)
	Get(ctx context.Context, id DependencyID) (Dependency, error)
	ByKind(ctx context.Context, kind DependencyKind) ([]DependencyID, error)
	Has(ctx context.Context, id DependencyID) (bool, error)
}

// DependencyBuilder represents a dependency builder.
type DependencyBuilder interface {
	Version(version Version) DependencyBuilder
	Kind(kind DependencyKind) DependencyBuilder
	Name(name string) DependencyBuilder
	DepVersion(depVersion string) DependencyBuilder
	Source(source string) DependencyBuilder
	Build() (Dependency, error)
}

// Dependency represents a locked runtime, tool, image, or module dependency.
type Dependency interface {
	ID() DependencyID
	Version() Version

	Kind() DependencyKind
	Name() string       // what (ex: protobuf)
	DepVersion() string // which version (ex: v33.0)
	Source() string     // where to obtain it (ex: buf.build/googleapis/googleapis)
}

// ArtifactRepository loads artifacts on demand.
type ArtifactRepository interface {
	IDsByProject(ctx context.Context, projectID ProjectID) ([]ArtifactID, error)
	IDsByVersion(ctx context.Context, versionID VersionID) ([]ArtifactID, error)
	Get(ctx context.Context, id ArtifactID) (Artifact, error)
	ByKind(ctx context.Context, kind ArtifactKind) ([]ArtifactID, error)
	Has(ctx context.Context, id ArtifactID) (bool, error)
}

// ArtifactBuilder represents an artifact builder.
type ArtifactBuilder interface {
	Section(section Section) ArtifactBuilder
	Path(path string) ArtifactBuilder
	File(file ArtifactFile) ArtifactBuilder
	Build() (Artifact, error)
}

// Artifact represents a file or directory that should exist after generation.
type Artifact interface {
	ID() ArtifactID
	Kind() ArtifactKind
	Section() Section
	Path() string

	File() ArtifactFile
	IsFile() bool
}

// ArtifactFileBuilder represents an artifact file builder.
type ArtifactFileBuilder interface {
	Kind(kind ArtifactFileKind) ArtifactFileBuilder
	Purpose(purpose string) ArtifactFileBuilder
	Specification(specification string) ArtifactFileBuilder
	AddConstraint(constraint string) ArtifactFileBuilder
	AddAcceptanceCriteria(criteria string) ArtifactFileBuilder
	Build() (ArtifactFile, error)
}

// ArtifactFile represents what an artifact file should contain.
type ArtifactFile interface {
	Kind() ArtifactFileKind
	Purpose() string
	Specification() string
	Constraints() []string
	AcceptanceCriteria() []string
}

// TaskRepository loads tasks on demand.
type TaskRepository interface {
	IDsByProject(ctx context.Context, projectID ProjectID) ([]TaskID, error)
	IDsByVersion(ctx context.Context, versionID VersionID) ([]TaskID, error)
	Get(ctx context.Context, id TaskID) (Task, error)

	RootsByProject(ctx context.Context, projectID ProjectID) ([]TaskID, error)
	RootsByVersion(ctx context.Context, versionID VersionID) ([]TaskID, error)

	Children(ctx context.Context, parentID TaskID) ([]TaskID, error)
	Dependents(ctx context.Context, taskID TaskID) ([]TaskID, error)
	Upstream(ctx context.Context, taskID TaskID) ([]TaskID, error)   // tasks this task depends on
	Downstream(ctx context.Context, taskID TaskID) ([]TaskID, error) // tasks depending on this task

	ReadyByProject(ctx context.Context, projectID ProjectID) ([]TaskID, error)
	ReadyByVersion(ctx context.Context, versionID VersionID) ([]TaskID, error)

	Has(ctx context.Context, id TaskID) (bool, error)
}

// TaskBuilder represents a task builder.
type TaskBuilder interface {
	Section(section Section) TaskBuilder
	Kind(kind TaskKind) TaskBuilder
	Action(action TaskAction) TaskBuilder
	Name(name string) TaskBuilder
	Description(description string) TaskBuilder
	Spec(spec TaskSpec) TaskBuilder
	ParentID(task TaskID) TaskBuilder
	AddDependency(task TaskID) TaskBuilder
	AddSubTask(task TaskID) TaskBuilder
	Build() (Task, error)
}

// Task represents one executable step in the generation DAG.
type Task interface {
	ID() TaskID
	Section() Section
	Kind() TaskKind
	Action() TaskAction
	ParentID() TaskID
	HasParentID() bool
	DependencyIDs() []TaskID
	SubTaskIDs() []TaskID
	Name() string
	Description() string
	Spec() TaskSpec
	HasSpec() bool
}

// TaskSpecBuilder represents a task specification builder.
type TaskSpecBuilder interface {
	Goal(goal string) TaskSpecBuilder
	Instructions(instructions string) TaskSpecBuilder
	AddConstraint(constraint string) TaskSpecBuilder
	AddAcceptanceCriteria(criteria string) TaskSpecBuilder
	Build() (TaskSpec, error)
}

// TaskSpec represents how a task should be executed.
type TaskSpec interface {
	Goal() string
	Instructions() string
	Constraints() []string
	AcceptanceCriteria() []string
}
