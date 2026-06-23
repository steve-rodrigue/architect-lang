package files

import "github.com/steve-rodrigue/architect-lang/architect/internal/domain/plan"

func toProjectData(project plan.Project) projectData {
	return projectData{
		ID:            string(project.ID()),
		Name:          project.Model().Name(),
		SchemaVersion: project.SchemaVersion(),
	}
}

func toVersionData(version plan.Version) versionData {
	return versionData{
		ID:        string(version.ID()),
		ProjectID: string(version.Project().ID()),
		Number:    version.Number(),
	}
}

func toSectionData(section plan.Section) sectionData {
	return sectionData{
		ID:       string(section.ID()),
		Version:  string(section.Version().ID()),
		Kind:     string(section.Kind()),
		Name:     section.Name(),
		Path:     section.Path(),
		ParentID: string(section.ParentID()),
	}
}

func toDependencyData(dependency plan.Dependency) dependencyData {
	return dependencyData{
		ID:         string(dependency.ID()),
		VersionID:  string(dependency.Version().ID()),
		Kind:       string(dependency.Kind()),
		Name:       dependency.Name(),
		DepVersion: dependency.DepVersion(),
		Source:     dependency.Source(),
	}
}

func toArtifactData(artifact plan.Artifact) artifactData {
	data := artifactData{
		ID:      string(artifact.ID()),
		Kind:    string(artifact.Kind()),
		Section: string(artifact.Section().ID()),
		Path:    artifact.Path(),
	}

	if artifact.IsFile() && artifact.File() != nil {
		data.File = toArtifactFileData(artifact.File())
	}

	return data
}

func toArtifactFileData(file plan.ArtifactFile) *artifactFileData {
	return &artifactFileData{
		Kind:               string(file.Kind()),
		Purpose:            file.Purpose(),
		Specification:      file.Specification(),
		Constraints:        file.Constraints(),
		AcceptanceCriteria: file.AcceptanceCriteria(),
	}
}

func toTaskData(task plan.Task) taskData {
	data := taskData{
		ID:            string(task.ID()),
		SectionID:     string(task.Section().ID()),
		Kind:          string(task.Kind()),
		Action:        string(task.Action()),
		ParentID:      string(task.ParentID()),
		DependencyIDs: taskIDsToStrings(task.DependencyIDs()),
		SubTaskIDs:    taskIDsToStrings(task.SubTaskIDs()),
		Name:          task.Name(),
		Description:   task.Description(),
	}

	if task.HasSpec() && task.Spec() != nil {
		data.Spec = toTaskSpecData(task.Spec())
	}

	return data
}

func toTaskSpecData(spec plan.TaskSpec) *taskSpecData {
	return &taskSpecData{
		Goal:               spec.Goal(),
		Instructions:       spec.Instructions(),
		Constraints:        spec.Constraints(),
		AcceptanceCriteria: spec.AcceptanceCriteria(),
	}
}

func taskIDsToStrings(ids []plan.TaskID) []string {
	values := make([]string, 0, len(ids))

	for _, id := range ids {
		values = append(values, string(id))
	}

	return values
}
