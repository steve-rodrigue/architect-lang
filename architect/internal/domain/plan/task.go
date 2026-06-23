package plan

type task struct {
	id            TaskID
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

func (t *task) ID() TaskID {
	return t.id
}

func (t *task) Section() Section {
	return t.section
}

func (t *task) Kind() TaskKind {
	return t.kind
}

func (t *task) Action() TaskAction {
	return t.action
}

func (t *task) ParentID() TaskID {
	return t.parentID
}

func (t *task) HasParentID() bool {
	return t.parentID != ""
}

func (t *task) DependencyIDs() []TaskID {
	return append([]TaskID(nil), t.dependencyIDs...)
}

func (t *task) SubTaskIDs() []TaskID {
	return append([]TaskID(nil), t.subTaskIDs...)
}

func (t *task) Name() string {
	return t.name
}

func (t *task) Description() string {
	return t.description
}

func (t *task) Spec() TaskSpec {
	return t.spec
}

func (t *task) HasSpec() bool {
	return t.spec != nil
}
