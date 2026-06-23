package files

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/plan"

type planServiceFileData struct {
	Projects     []projectData    `json:"projects"`
	Versions     []versionData    `json:"versions"`
	Sections     []sectionData    `json:"sections"`
	Dependencies []dependencyData `json:"dependencies"`
	Artifacts    []artifactData   `json:"artifacts"`
	Tasks        []taskData       `json:"tasks"`
}

type projectData struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	SchemaVersion string `json:"schemaVersion"`
}

type versionData struct {
	ID        string `json:"id"`
	ProjectID string `json:"projectId"`
	Number    string `json:"number"`
}

type sectionData struct {
	ID       string `json:"id"`
	Version  string `json:"versionId"`
	Kind     string `json:"kind"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	ParentID string `json:"parentId,omitempty"`
}

type dependencyData struct {
	ID         string `json:"id"`
	VersionID  string `json:"versionId"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	DepVersion string `json:"depVersion"`
	Source     string `json:"source,omitempty"`
}

type artifactData struct {
	ID      string            `json:"id"`
	Kind    string            `json:"kind"`
	Section string            `json:"sectionId"`
	Path    string            `json:"path"`
	File    *artifactFileData `json:"file,omitempty"`
}

type artifactFileData struct {
	Kind               string   `json:"kind"`
	Purpose            string   `json:"purpose"`
	Specification      string   `json:"specification"`
	Constraints        []string `json:"constraints"`
	AcceptanceCriteria []string `json:"acceptanceCriteria"`
}

type taskData struct {
	ID            string        `json:"id"`
	SectionID     string        `json:"sectionId"`
	Kind          string        `json:"kind"`
	Action        string        `json:"action"`
	ParentID      string        `json:"parentId,omitempty"`
	DependencyIDs []string      `json:"dependencyIds"`
	SubTaskIDs    []string      `json:"subTaskIds"`
	Name          string        `json:"name"`
	Description   string        `json:"description"`
	Spec          *taskSpecData `json:"spec,omitempty"`
}

type taskSpecData struct {
	Goal               string   `json:"goal"`
	Instructions       string   `json:"instructions"`
	Constraints        []string `json:"constraints"`
	AcceptanceCriteria []string `json:"acceptanceCriteria"`
}

func (d *planServiceFileData) upsertProject(value projectData) {
	for index := range d.Projects {
		if d.Projects[index].ID == value.ID {
			d.Projects[index] = value
			return
		}
	}

	d.Projects = append(d.Projects, value)
}

func (d *planServiceFileData) upsertVersion(value versionData) {
	for index := range d.Versions {
		if d.Versions[index].ID == value.ID {
			d.Versions[index] = value
			return
		}
	}

	d.Versions = append(d.Versions, value)
}

func (d *planServiceFileData) upsertSection(value sectionData) {
	for index := range d.Sections {
		if d.Sections[index].ID == value.ID {
			d.Sections[index] = value
			return
		}
	}

	d.Sections = append(d.Sections, value)
}

func (d *planServiceFileData) upsertDependency(value dependencyData) {
	for index := range d.Dependencies {
		if d.Dependencies[index].ID == value.ID {
			d.Dependencies[index] = value
			return
		}
	}

	d.Dependencies = append(d.Dependencies, value)
}

func (d *planServiceFileData) upsertArtifact(value artifactData) {
	for index := range d.Artifacts {
		if d.Artifacts[index].ID == value.ID {
			d.Artifacts[index] = value
			return
		}
	}

	d.Artifacts = append(d.Artifacts, value)
}

func (d *planServiceFileData) upsertTask(value taskData) {
	for index := range d.Tasks {
		if d.Tasks[index].ID == value.ID {
			d.Tasks[index] = value
			return
		}
	}

	d.Tasks = append(d.Tasks, value)
}

func (d *planServiceFileData) hasProject(id plan.ProjectID) bool {
	for _, project := range d.Projects {
		if project.ID == string(id) {
			return true
		}
	}

	return false
}

func (d *planServiceFileData) hasVersion(id plan.VersionID) bool {
	for _, version := range d.Versions {
		if version.ID == string(id) {
			return true
		}
	}

	return false
}

func (d *planServiceFileData) hasSection(id plan.SectionID) bool {
	for _, section := range d.Sections {
		if section.ID == string(id) {
			return true
		}
	}

	return false
}

func (d *planServiceFileData) hasTask(id plan.TaskID) bool {
	for _, task := range d.Tasks {
		if task.ID == string(id) {
			return true
		}
	}

	return false
}
