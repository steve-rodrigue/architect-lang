package projects

// NewProjectBuilder creates a new project builder
func NewProjectBuilder() ProjectBuilder {
	return &projectBuilder{
		versions: make([]Version, 0),
	}
}

// NewVersionBuilder creates a new version builder
func NewVersionBuilder() VersionBuilder {
	return &versionBuilder{
		serviceFiles:    make([]string, 0),
		objectFiles:     make([]string, 0),
		deploymentFiles: make([]string, 0),
		nextVersions:    make([]Version, 0),
	}
}

// ProjectBuilder represents a project builder
type ProjectBuilder interface {
	Name(name string) ProjectBuilder
	AddVersion(version Version) ProjectBuilder
	Build() (Project, error)
}

// Project represents a project
type Project interface {
	Name() string
	Versions() []Version
}

// VersionBuilder represents a project version builder
type VersionBuilder interface {
	Number(number string) VersionBuilder
	AddServiceFile(file string) VersionBuilder
	AddObjectFile(file string) VersionBuilder
	AddDeploymentFile(file string) VersionBuilder
	AddNextVersion(version Version) VersionBuilder
	Build() (Version, error)
}

// Version represents a project version
type Version interface {
	Number() string
	ServiceFiles() []string
	ObjectFiles() []string
	DeploymentFiles() []string
	NextVersions() []Version
}
