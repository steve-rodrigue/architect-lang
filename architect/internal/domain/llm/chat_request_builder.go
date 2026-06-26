package llm

import "fmt"

type chatRequestBuilder struct {
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

func (b *chatRequestBuilder) Model(model string) ChatRequestBuilder {
	b.model = model
	return b
}

func (b *chatRequestBuilder) AddMessage(message ChatMessage) ChatRequestBuilder {
	b.messages = append(b.messages, message)
	return b
}

func (b *chatRequestBuilder) Temperature(temperature float32) ChatRequestBuilder {
	b.temperature = temperature
	b.hasTemperature = true
	return b
}

func (b *chatRequestBuilder) TopP(topP float32) ChatRequestBuilder {
	b.topP = topP
	b.hasTopP = true
	return b
}

func (b *chatRequestBuilder) MaxTokens(tokens int) ChatRequestBuilder {
	b.maxTokens = tokens
	b.hasMaxTokens = true
	return b
}

func (b *chatRequestBuilder) AddStop(stop string) ChatRequestBuilder {
	b.stop = append(b.stop, stop)
	return b
}

func (b *chatRequestBuilder) Stream(stream bool) ChatRequestBuilder {
	b.stream = stream
	return b
}

func (b *chatRequestBuilder) Build() (ChatRequest, error) {
	if b.model == "" {
		return nil, fmt.Errorf("chat request model is required")
	}

	if len(b.messages) == 0 {
		return nil, fmt.Errorf("at least one chat message is required")
	}

	return &chatRequest{
		model:          b.model,
		messages:       b.messages,
		temperature:    b.temperature,
		hasTemperature: b.hasTemperature,
		topP:           b.topP,
		hasTopP:        b.hasTopP,
		maxTokens:      b.maxTokens,
		hasMaxTokens:   b.hasMaxTokens,
		stop:           b.stop,
		stream:         b.stream,
	}, nil
}
