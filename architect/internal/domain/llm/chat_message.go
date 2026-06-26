package llm

type chatMessage struct {
	role ChatRole

	content string
}

func (m *chatMessage) Role() ChatRole {
	return m.role
}

func (m *chatMessage) Content() string {
	return m.content
}
