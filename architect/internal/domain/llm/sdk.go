package llm

import "context"

type ChatRole string

const (
	ChatRoleSystem    ChatRole = "system"
	ChatRoleUser      ChatRole = "user"
	ChatRoleAssistant ChatRole = "assistant"
)

// NewChatRequestBuilder creates a new chat request builder.
func NewChatRequestBuilder() ChatRequestBuilder {
	return &chatRequestBuilder{
		messages: make([]ChatMessage, 0),
		stop:     make([]string, 0),
	}
}

// NewChatResponseBuilder creates a new chat response builder.
func NewChatResponseBuilder() ChatResponseBuilder {
	return &chatResponseBuilder{}
}

// NewChatMessageBuilder creates a new chat message builder.
func NewChatMessageBuilder() ChatMessageBuilder {
	return &chatMessageBuilder{}
}

// Repository represents an LLM backend.
type Repository interface {
	// Health verifies that the backend is available.
	Health(ctx context.Context) error
	// Models returns the available models.
	Models(ctx context.Context) ([]string, error)
	// Chat performs a chat completion.
	Chat(ctx context.Context, request ChatRequest) (ChatResponse, error)
}

// ChatRequestBuilder represents a chat request builder.
type ChatRequestBuilder interface {
	Model(model string) ChatRequestBuilder
	AddMessage(message ChatMessage) ChatRequestBuilder
	Temperature(temperature float32) ChatRequestBuilder
	TopP(topP float32) ChatRequestBuilder
	MaxTokens(tokens int) ChatRequestBuilder
	AddStop(stop string) ChatRequestBuilder
	Stream(stream bool) ChatRequestBuilder
	Build() (ChatRequest, error)
}

// ChatRequest represents a chat request
type ChatRequest interface {
	Model() string
	Messages() []ChatMessage
	Temperature() float32
	HasTemperature() bool
	TopP() float32
	HasTopP() bool
	MaxTokens() int
	HasMaxTokens() bool
	Stop() []string
	Stream() bool
}

// ChatResponseBuilder represents a chat response builder.
type ChatResponseBuilder interface {
	Model(model string) ChatResponseBuilder
	Message(message ChatMessage) ChatResponseBuilder
	PromptTokens(tokens int) ChatResponseBuilder
	CompletionTokens(tokens int) ChatResponseBuilder
	TotalTokens(tokens int) ChatResponseBuilder
	Build() (ChatResponse, error)
}

// ChatResponse represents a chat response
type ChatResponse interface {
	Model() string
	Message() ChatMessage
	PromptTokens() int
	CompletionTokens() int
	TotalTokens() int
}

// ChatMessageBuilder represents a chat message builder.
type ChatMessageBuilder interface {
	Role(role ChatRole) ChatMessageBuilder
	Content(content string) ChatMessageBuilder
	Build() (ChatMessage, error)
}

// ChatMessage represents a chat message
type ChatMessage interface {
	Role() ChatRole
	Content() string
}
