package llm

import (
	"net/http"
	"strings"

	"github.com/steve-rodrigue/architect-lang/architect/internal/domain/llm"
)

// NewOllamaRepository creates a new Ollama-backed LLM repository.
func NewOllamaRepository(baseURL string, httpClient *http.Client) llm.Repository {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &ollamaRepository{
		baseURL:    strings.TrimRight(baseURL, "/"),
		httpClient: httpClient,
	}
}
