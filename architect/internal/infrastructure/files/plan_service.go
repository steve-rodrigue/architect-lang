package files

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/plan"
)

type planService struct {
	path     string
	fileName string
}

func (s *planService) SaveProject(project plan.Project) error {
	if project == nil {
		return fmt.Errorf("project is required")
	}

	data, err := s.read()
	if err != nil {
		return err
	}

	data.upsertProject(toProjectData(project))

	return s.write(data)
}

func (s *planService) SaveVersion(version plan.Version) error {
	if version == nil {
		return fmt.Errorf("version is required")
	}

	data, err := s.read()
	if err != nil {
		return err
	}

	if !data.hasProject(version.Project().ID()) {
		return fmt.Errorf("version references unknown project %q", version.Project().ID())
	}

	data.upsertVersion(toVersionData(version))

	return s.write(data)
}

func (s *planService) SaveSection(section plan.Section) error {
	if section == nil {
		return fmt.Errorf("section is required")
	}

	data, err := s.read()
	if err != nil {
		return err
	}

	if !data.hasVersion(section.Version().ID()) {
		return fmt.Errorf("section references unknown version %q", section.Version().ID())
	}

	if section.HasParentID() && !data.hasSection(section.ParentID()) {
		return fmt.Errorf("section references unknown parent section %q", section.ParentID())
	}

	data.upsertSection(toSectionData(section))

	return s.write(data)
}

func (s *planService) SaveDependency(dependency plan.Dependency) error {
	if dependency == nil {
		return fmt.Errorf("dependency is required")
	}

	data, err := s.read()
	if err != nil {
		return err
	}

	if !data.hasVersion(dependency.Version().ID()) {
		return fmt.Errorf("dependency references unknown version %q", dependency.Version().ID())
	}

	data.upsertDependency(toDependencyData(dependency))

	return s.write(data)
}

func (s *planService) SaveArtifact(artifact plan.Artifact) error {
	if artifact == nil {
		return fmt.Errorf("artifact is required")
	}

	data, err := s.read()
	if err != nil {
		return err
	}

	if !data.hasSection(artifact.Section().ID()) {
		return fmt.Errorf("artifact references unknown section %q", artifact.Section().ID())
	}

	data.upsertArtifact(toArtifactData(artifact))

	return s.write(data)
}

func (s *planService) SaveTask(task plan.Task) error {
	if task == nil {
		return fmt.Errorf("task is required")
	}

	data, err := s.read()
	if err != nil {
		return err
	}

	if !data.hasSection(task.Section().ID()) {
		return fmt.Errorf("task references unknown section %q", task.Section().ID())
	}

	if task.HasParentID() && !data.hasTask(task.ParentID()) {
		return fmt.Errorf("task references unknown parent task %q", task.ParentID())
	}

	for _, dependencyID := range task.DependencyIDs() {
		if !data.hasTask(dependencyID) {
			return fmt.Errorf("task references unknown dependency task %q", dependencyID)
		}
	}

	for _, subTaskID := range task.SubTaskIDs() {
		if !data.hasTask(subTaskID) {
			return fmt.Errorf("task references unknown sub task %q", subTaskID)
		}
	}

	data.upsertTask(toTaskData(task))

	return s.write(data)
}

func (s *planService) read() (*planServiceFileData, error) {
	path, err := s.filePath()
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return &planServiceFileData{
			Projects:     make([]projectData, 0),
			Versions:     make([]versionData, 0),
			Sections:     make([]sectionData, 0),
			Dependencies: make([]dependencyData, 0),
			Artifacts:    make([]artifactData, 0),
			Tasks:        make([]taskData, 0),
		}, nil
	}

	if err != nil {
		return nil, fmt.Errorf("read plan file %q: %w", path, err)
	}

	var data planServiceFileData
	if err := json.Unmarshal(content, &data); err != nil {
		return nil, fmt.Errorf("decode plan file %q: %w", path, err)
	}

	return &data, nil
}

func (s *planService) write(data *planServiceFileData) error {
	path, err := s.filePath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return fmt.Errorf("create plan directory %q: %w", filepath.Dir(path), err)
	}

	content, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("encode plan file: %w", err)
	}

	content = append(content, '\n')

	if err := os.WriteFile(path, content, 0644); err != nil {
		return fmt.Errorf("write plan file %q: %w", path, err)
	}

	return nil
}

func (s *planService) filePath() (string, error) {
	if s.path == "" {
		return "", fmt.Errorf("plan file path is required")
	}

	info, err := os.Stat(s.path)
	if err == nil && info.IsDir() {
		return filepath.Join(s.path, s.fileName), nil
	}

	if os.IsNotExist(err) && filepath.Ext(s.path) == "" {
		return filepath.Join(s.path, s.fileName), nil
	}

	return s.path, nil
}
