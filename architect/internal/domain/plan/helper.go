package plan

import (
	"path/filepath"
	"strings"
)

// GenProjectID generates a project identifer
func GenProjectID(projectName string) ProjectID {
	return ProjectID("project:" + cleanIDPart(projectName))
}

// GetVersionID generates a version identifier
func GenVersionID(project Project, number string) VersionID {
	return VersionID(
		strings.Join([]string{
			"version",
			string(project.ID()),
			cleanIDPart(number),
		}, ":"),
	)
}

// GenSectionID generates a section identifier
func GenSectionID(version Version, parentID SectionID, kind SectionKind, name string) SectionID {
	parts := []string{
		"section",
		string(version.ID()),
		string(kind),
	}

	if parentID != "" {
		parts = append(parts, string(parentID))
	}

	parts = append(parts, cleanIDPart(name))

	return SectionID(strings.Join(parts, ":"))
}

// GenSectionPath generates a section path
func GenSectionPath(parentID SectionID, kind SectionKind, name string) string {
	current := string(kind) + ":" + cleanIDPart(name)

	if parentID == "" {
		return current
	}

	return string(parentID) + "/" + current
}

// GenDependencyID generates a dependency identifier
func GenDependencyID(version Version, kind DependencyKind, name string, depVersion string) DependencyID {
	return DependencyID(
		strings.Join([]string{
			"dependency",
			string(version.ID()),
			string(kind),
			cleanIDPart(name),
			cleanIDPart(depVersion),
		}, ":"),
	)
}

// GenArtifactID generates an artifact identifier
func GenArtifactID(section Section, path string) ArtifactID {
	return ArtifactID(
		strings.Join([]string{
			"artifact",
			string(section.ID()),
			cleanPath(path),
		}, ":"),
	)
}

// GenTaskID generates a task identifier
func GenTaskID(section Section, kind TaskKind, action TaskAction, name string) TaskID {
	return TaskID(
		strings.Join([]string{
			"task",
			string(section.ID()),
			string(kind),
			string(action),
			cleanIDPart(name),
		}, ":"),
	)
}

func cleanIDPart(value string) string {
	value = strings.TrimSpace(value)
	value = strings.ReplaceAll(value, "\\", "/")
	value = strings.ReplaceAll(value, " ", "_")
	value = strings.ReplaceAll(value, ":", "_")

	return value
}

func cleanPath(value string) string {
	value = strings.TrimSpace(value)
	value = strings.ReplaceAll(value, "\\", "/")
	value = filepath.Clean(value)
	value = strings.ReplaceAll(value, "\\", "/")
	value = strings.TrimPrefix(value, "./")
	value = strings.ReplaceAll(value, ":", "_")

	return value
}
