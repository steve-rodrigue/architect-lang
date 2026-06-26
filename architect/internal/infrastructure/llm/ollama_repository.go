package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/llm"
)

type ollamaRepository struct {
	baseURL    string
	httpClient *http.Client
}

type ollamaTagsResponse struct {
	Models []ollamaModel `json:"models"`
}

type ollamaModel struct {
	Name string `json:"name"`
}

type ollamaChatRequest struct {
	Model    string          `json:"model"`
	Messages []ollamaMessage `json:"messages"`
	Stream   bool            `json:"stream"`
	Options  map[string]any  `json:"options,omitempty"`
}

type ollamaMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ollamaChatResponse struct {
	Model           string        `json:"model"`
	Message         ollamaMessage `json:"message"`
	Done            bool          `json:"done"`
	PromptEvalCount int           `json:"prompt_eval_count"`
	EvalCount       int           `json:"eval_count"`
}

func (r *ollamaRepository) Health(ctx context.Context) error {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, r.baseURL+"/api/tags", nil)
	if err != nil {
		return err
	}

	response, err := r.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		body, _ := io.ReadAll(response.Body)
		return fmt.Errorf("ollama health failed with status %d: %s", response.StatusCode, string(body))
	}

	return nil
}

func (r *ollamaRepository) Models(ctx context.Context) ([]string, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, r.baseURL+"/api/tags", nil)
	if err != nil {
		return nil, err
	}

	response, err := r.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		body, _ := io.ReadAll(response.Body)
		return nil, fmt.Errorf("ollama models failed with status %d: %s", response.StatusCode, string(body))
	}

	var payload ollamaTagsResponse
	if err := json.NewDecoder(response.Body).Decode(&payload); err != nil {
		return nil, fmt.Errorf("decode ollama models response: %w", err)
	}

	models := make([]string, 0, len(payload.Models))
	for _, model := range payload.Models {
		models = append(models, model.Name)
	}

	return models, nil
}

func (r *ollamaRepository) Chat(ctx context.Context, request llm.ChatRequest) (llm.ChatResponse, error) {
	if request == nil {

		return nil, fmt.Errorf("chat request is required")
	}
	if request.Stream() {
		return nil, fmt.Errorf("ollama streaming chat is not supported by this repository")
	}
	payload := ollamaChatRequest{
		Model:    request.Model(),
		Messages: toOllamaMessages(request.Messages()),
		Stream:   false,
		Options:  toOllamaOptions(request),
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("encode ollama chat request: %w", err)
	}

	httpRequest, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		r.baseURL+"/api/chat",
		bytes.NewReader(body),
	)
	if err != nil {
		return nil, err
	}

	httpRequest.Header.Set("Content-Type", "application/json")

	response, err := r.httpClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		body, _ := io.ReadAll(response.Body)
		return nil, fmt.Errorf("ollama chat failed with status %d: %s", response.StatusCode, string(body))
	}

	if request.Stream() {
		return nil, fmt.Errorf("ollama streaming chat is not supported by this repository")
	}

	var payloadResponse ollamaChatResponse
	if err := json.NewDecoder(response.Body).Decode(&payloadResponse); err != nil {
		return nil, fmt.Errorf("decode ollama chat response: %w", err)
	}

	message, err := llm.NewChatMessageBuilder().
		Role(llm.ChatRole(payloadResponse.Message.Role)).
		Content(payloadResponse.Message.Content).
		Build()
	if err != nil {
		return nil, err
	}

	totalTokens := payloadResponse.PromptEvalCount + payloadResponse.EvalCount

	return llm.NewChatResponseBuilder().
		Model(payloadResponse.Model).
		Message(message).
		PromptTokens(payloadResponse.PromptEvalCount).
		CompletionTokens(payloadResponse.EvalCount).
		TotalTokens(totalTokens).
		Build()
}

func toOllamaMessages(messages []llm.ChatMessage) []ollamaMessage {
	values := make([]ollamaMessage, 0, len(messages))

	for _, message := range messages {
		values = append(values, ollamaMessage{
			Role:    string(message.Role()),
			Content: message.Content(),
		})
	}

	return values
}

func toOllamaOptions(request llm.ChatRequest) map[string]any {
	options := make(map[string]any)

	if request.HasTemperature() {
		options["temperature"] = request.Temperature()
	}

	if request.HasTopP() {
		options["top_p"] = request.TopP()
	}

	if request.HasMaxTokens() {
		options["num_predict"] = request.MaxTokens()
	}

	if len(request.Stop()) > 0 {
		options["stop"] = request.Stop()
	}

	if len(options) == 0 {
		return nil
	}

	return options
}
