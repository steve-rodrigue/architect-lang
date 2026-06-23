package plan

import (
	"path/filepath"
	"strings"
)

func projectID(projectName string) ProjectID {
	return ProjectID("project:" + cleanIDPart(projectName))
}

func versionID(project Project, number string) VersionID {
	return VersionID(
		strings.Join([]string{
			"version",
			string(project.ID()),
			cleanIDPart(number),
		}, ":"),
	)
}

func sectionID(version Version, parentID SectionID, kind SectionKind, name string) SectionID {
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

func sectionPath(parentID SectionID, kind SectionKind, name string) string {
	current := string(kind) + ":" + cleanIDPart(name)

	if parentID == "" {
		return current
	}

	return string(parentID) + "/" + current
}

func dependencyID(version Version, kind DependencyKind, name string, depVersion string) DependencyID {
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

func artifactID(section Section, path string) ArtifactID {
	return ArtifactID(
		strings.Join([]string{
			"artifact",
			string(section.ID()),
			cleanPath(path),
		}, ":"),
	)
}

func taskID(section Section, kind TaskKind, action TaskAction, name string) TaskID {
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
