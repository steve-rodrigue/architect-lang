package plan

import "fmt"

type taskBuilder struct {
	section       Section
	kind          TaskKind
	action        TaskAction
	name          string
	description   string
	spec          TaskSpec
	parentID      TaskID
	dependencyIDs []TaskID
	subTaskIDs    []TaskID
}

func (b *taskBuilder) Section(section Section) TaskBuilder {
	b.section = section
	return b
}

func (b *taskBuilder) Kind(kind TaskKind) TaskBuilder {
	b.kind = kind
	return b
}

func (b *taskBuilder) Action(action TaskAction) TaskBuilder {
	b.action = action
	return b
}

func (b *taskBuilder) Name(name string) TaskBuilder {
	b.name = name
	return b
}

func (b *taskBuilder) Description(description string) TaskBuilder {
	b.description = description
	return b
}

func (b *taskBuilder) Spec(spec TaskSpec) TaskBuilder {
	b.spec = spec
	return b
}

func (b *taskBuilder) ParentID(task TaskID) TaskBuilder {
	b.parentID = task
	return b
}

func (b *taskBuilder) AddDependency(task TaskID) TaskBuilder {
	b.dependencyIDs = append(b.dependencyIDs, task)
	return b
}

func (b *taskBuilder) AddSubTask(task TaskID) TaskBuilder {
	b.subTaskIDs = append(b.subTaskIDs, task)
	return b
}

func (b *taskBuilder) Build() (Task, error) {
	if b.section == nil {
		return nil, fmt.Errorf("task section is required")
	}

	if b.kind == "" {
		return nil, fmt.Errorf("task kind is required")
	}

	if b.action == "" {
		return nil, fmt.Errorf("task action is required")
	}

	if b.name == "" {
		return nil, fmt.Errorf("task name is required")
	}

	if b.description == "" {
		return nil, fmt.Errorf("task description is required")
	}

	return &task{
		id:            taskID(b.section, b.kind, b.action, b.name),
		section:       b.section,
		kind:          b.kind,
		action:        b.action,
		name:          b.name,
		description:   b.description,
		spec:          b.spec,
		parentID:      b.parentID,
		dependencyIDs: append([]TaskID(nil), b.dependencyIDs...),
		subTaskIDs:    append([]TaskID(nil), b.subTaskIDs...),
	}, nil
}
