package llm

import "fmt"

type chatMessageBuilder struct {
	role ChatRole

	content string
}

func (b *chatMessageBuilder) Role(role ChatRole) ChatMessageBuilder {
	b.role = role
	return b
}

func (b *chatMessageBuilder) Content(content string) ChatMessageBuilder {
	b.content = content
	return b
}

func (b *chatMessageBuilder) Build() (ChatMessage, error) {
	if b.role == "" {
		return nil, fmt.Errorf("chat message role is required")
	}

	if b.content == "" {
		return nil, fmt.Errorf("chat message content is required")
	}

	return &chatMessage{
		role:    b.role,
		content: b.content,
	}, nil
}
