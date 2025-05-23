package deepseek

import "github.com/go-logr/logr"

const (
	User      string = "user"
	Assistant string = "assistant"
	System    string = "system"

	// models
	DeepSeekChat     string = "deepseek-chat"
	DeepSeekReasoner string = "deepseek-reasoner"

	// tempertaure types
	NoVariability       float32 = 0.0
	DefaultVariability  float32 = 0.7
	ModerateVariability float32 = 1.0
	HighVariability     float32 = 1.5
	MaximumVariability  float32 = 2.0
)

type Client struct {
	Host         string
	Token        string
	ChatEndpoint string
	Log          logr.Logger
}

type ChatRequest struct {
	Model       string           `json:"model"`
	Temperature float32          `json:"temperature"`
	Messages    []RequestMessage `json:"messages"`
	MaxTokens   int              `json:"max_tokens,omitempty"`
}

type RequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatResponse struct {
	Id      string            `json:"id"`
	Object  string            `json:"object"`
	Created int64             `json:"created"`
	Model   string            `json:"model"`
	Usage   ResponseUsage     `json:"usage"`
	Choices []ResponseChoices `json:"choices"`
}

type ResponseUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ResponseChoices struct {
	Message      ResponseMessage `json:"message"`
	FinishReason string          `json:"finish_reason"`
	Index        int             `json:"index"`
}

type ResponseMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
