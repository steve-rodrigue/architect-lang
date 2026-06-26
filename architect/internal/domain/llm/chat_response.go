package llm

type chatResponse struct {
	model string

	message ChatMessage

	promptTokens int

	completionTokens int

	totalTokens int
}

func (r *chatResponse) Model() string {
	return r.model
}

func (r *chatResponse) Message() ChatMessage {
	return r.message
}

func (r *chatResponse) PromptTokens() int {
	return r.promptTokens
}

func (r *chatResponse) CompletionTokens() int {
	return r.completionTokens
}

func (r *chatResponse) TotalTokens() int {
	return r.totalTokens
}
