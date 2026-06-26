package llm

import "fmt"

type chatResponseBuilder struct {
	model string

	message ChatMessage

	promptTokens int

	completionTokens int

	totalTokens int
}

func (b *chatResponseBuilder) Model(model string) ChatResponseBuilder {
	b.model = model
	return b
}

func (b *chatResponseBuilder) Message(message ChatMessage) ChatResponseBuilder {
	b.message = message
	return b
}

func (b *chatResponseBuilder) PromptTokens(tokens int) ChatResponseBuilder {
	b.promptTokens = tokens
	return b
}

func (b *chatResponseBuilder) CompletionTokens(tokens int) ChatResponseBuilder {
	b.completionTokens = tokens
	return b
}

func (b *chatResponseBuilder) TotalTokens(tokens int) ChatResponseBuilder {
	b.totalTokens = tokens
	return b
}

func (b *chatResponseBuilder) Build() (ChatResponse, error) {
	if b.model == "" {
		return nil, fmt.Errorf("chat response model is required")
	}

	if b.message == nil {
		return nil, fmt.Errorf("chat response message is required")
	}

	return &chatResponse{
		model:            b.model,
		message:          b.message,
		promptTokens:     b.promptTokens,
		completionTokens: b.completionTokens,
		totalTokens:      b.totalTokens,
	}, nil
}
