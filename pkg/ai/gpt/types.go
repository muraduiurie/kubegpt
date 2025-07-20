package gpt

import "github.com/go-logr/logr"

const (
	User      string = "user"
	Assistant string = "assistant"
	System    string = "system"

	// chatgpt models
	Gpt3_5_turbo_0301 string = "gpt-3.5-turbo-0301"
	Gpt3_5_turbo_1106 string = "gpt-3.5-turbo-1106"
	Gpt3_5_turbo      string = "gpt-3.5-turbo"
	Gpt4_1            string = "gpt-4.1"
	Gpt4o             string = "gpt-4o"
	Gpt4oMini         string = "gpt-4o-mini"
	Gpt4o_turbo       string = "gpt-4-turbo"
	TTS1              string = "tts-1"
	TTS1_HD           string = "tts-1-hd"

	// tempertaure types
	NoVariability       float32 = 0.0
	DefaultVariability  float32 = 0.7
	ModerateVariability float32 = 1.0
	HighVariability     float32 = 1.5
	MaximumVariability  float32 = 2.0

	// content types
	InputFile string = "input_file"
	InputText string = "input_text"
)

type Client struct {
	Host            string
	Token           string
	ChatEndpoint    string
	FileUrlEndpoint string
	Log             logr.Logger
}

type AutoToTextResponse struct {
	Text string `json:"text"`
}

type ChatRequest struct {
	Model       string           `json:"model"`
	Temperature float32          `json:"temperature"`
	Messages    []RequestMessage `json:"messages,omitempty"`
	MaxTokens   int              `json:"max_tokens,omitempty"`
	Input       []InputUrl       `json:"input"`
}

type ContentUrl struct {
	Type    string `json:"type"`
	FileUrl string `json:"file_url,omitempty"`
	Text    string `json:"text,omitempty"`
}

type InputUrl struct {
	Role    string       `json:"role"`
	Content []ContentUrl `json:"content"`
}

type RequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type TextToAudioRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
	Voice string `json:"voice"`
}

type ResponseUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ResponseMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ResponseChoices struct {
	Message      ResponseMessage `json:"message"`
	FinishReason string          `json:"finish_reason"`
	Index        int             `json:"index"`
}

type ChatResponse struct {
	Id      string            `json:"id"`
	Object  string            `json:"object"`
	Created int64             `json:"created"`
	Model   string            `json:"model"`
	Usage   ResponseUsage     `json:"usage"`
	Choices []ResponseChoices `json:"choices"`
}

type ClientResponse struct {
	Message string `json:"message"`
	Contact bool   `json:"did client mention to be contacted?"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Product string `json:"desired product"`
}
