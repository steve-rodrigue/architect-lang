package llm

type chatRequest struct {
	model string

	messages []ChatMessage

	temperature    float32
	hasTemperature bool

	topP    float32
	hasTopP bool

	maxTokens    int
	hasMaxTokens bool

	stop []string

	stream bool
}

func (r *chatRequest) Model() string {
	return r.model
}

func (r *chatRequest) Messages() []ChatMessage {
	return r.messages
}

func (r *chatRequest) Temperature() float32 {
	return r.temperature
}

func (r *chatRequest) HasTemperature() bool {
	return r.hasTemperature
}

func (r *chatRequest) TopP() float32 {
	return r.topP
}

func (r *chatRequest) HasTopP() bool {
	return r.hasTopP
}

func (r *chatRequest) MaxTokens() int {
	return r.maxTokens
}

func (r *chatRequest) HasMaxTokens() bool {
	return r.hasMaxTokens
}

func (r *chatRequest) Stop() []string {
	return r.stop
}

func (r *chatRequest) Stream() bool {
	return r.stream
}
