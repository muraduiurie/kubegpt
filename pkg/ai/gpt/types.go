package gpt

import (
	"encoding/json"
	"github.com/go-logr/logr"
)

const (
	// tempertaure types
	NoVariability       float32 = 0.0
	DefaultVariability  float32 = 0.7
	ModerateVariability float32 = 1.0
	HighVariability     float32 = 1.5
	MaximumVariability  float32 = 2.0

	// content types
	InputFile  string = "input_file"
	InputImage string = "input_image"
	InputText  string = "input_text"
)

type Client struct {
	Host              string
	Token             string
	ResponsesEndpoint string
	Log               logr.Logger
}

type Requester interface {
	Marshal() ([]byte, error)
}

type Responser interface {
	Unmarshal(b []byte) error
}

// requests
type TextInputRequest struct {
	Model string `json:"model"`
	Input string `json:"input"`
}

type ImageInputRequest struct {
	Model string `json:"model"`
	Input []struct {
		Role    string `json:"role"`
		Content []struct {
			Type     string `json:"type"`
			Text     string `json:"text,omitempty"`
			ImageUrl string `json:"image_url,omitempty"`
		} `json:"content"`
	} `json:"input"`
}

type FileInputRequest struct {
	Model string `json:"model"`
	Input []struct {
		Role    string `json:"role"`
		Content []struct {
			Type    string `json:"type"`
			Text    string `json:"text,omitempty"`
			FileUrl string `json:"file_url,omitempty"`
		} `json:"content"`
	} `json:"input"`
}

// responses
type ImageInputResponse struct {
	Id                string      `json:"id"`
	Object            string      `json:"object"`
	CreatedAt         int         `json:"created_at"`
	Status            string      `json:"status"`
	Error             interface{} `json:"error"`
	IncompleteDetails interface{} `json:"incomplete_details"`
	Instructions      interface{} `json:"instructions"`
	MaxOutputTokens   interface{} `json:"max_output_tokens"`
	Model             string      `json:"model"`
	Output            []struct {
		Type    string `json:"type"`
		Id      string `json:"id"`
		Status  string `json:"status"`
		Role    string `json:"role"`
		Content []struct {
			Type        string        `json:"type"`
			Text        string        `json:"text"`
			Annotations []interface{} `json:"annotations"`
		} `json:"content"`
	} `json:"output"`
	ParallelToolCalls  bool        `json:"parallel_tool_calls"`
	PreviousResponseId interface{} `json:"previous_response_id"`
	Reasoning          struct {
		Effort  interface{} `json:"effort"`
		Summary interface{} `json:"summary"`
	} `json:"reasoning"`
	Store       bool    `json:"store"`
	Temperature float64 `json:"temperature"`
	Text        struct {
		Format struct {
			Type string `json:"type"`
		} `json:"format"`
	} `json:"text"`
	ToolChoice string        `json:"tool_choice"`
	Tools      []interface{} `json:"tools"`
	TopP       float64       `json:"top_p"`
	Truncation string        `json:"truncation"`
	Usage      struct {
		InputTokens        int `json:"input_tokens"`
		InputTokensDetails struct {
			CachedTokens int `json:"cached_tokens"`
		} `json:"input_tokens_details"`
		OutputTokens        int `json:"output_tokens"`
		OutputTokensDetails struct {
			ReasoningTokens int `json:"reasoning_tokens"`
		} `json:"output_tokens_details"`
		TotalTokens int `json:"total_tokens"`
	} `json:"usage"`
	User     interface{} `json:"user"`
	Metadata struct {
	} `json:"metadata"`
}

type TextInputResponse struct {
	Id                string      `json:"id"`
	Object            string      `json:"object"`
	CreatedAt         int         `json:"created_at"`
	Status            string      `json:"status"`
	Error             interface{} `json:"error"`
	IncompleteDetails interface{} `json:"incomplete_details"`
	Instructions      interface{} `json:"instructions"`
	MaxOutputTokens   interface{} `json:"max_output_tokens"`
	Model             string      `json:"model"`
	Output            []struct {
		Type    string `json:"type"`
		Id      string `json:"id"`
		Status  string `json:"status"`
		Role    string `json:"role"`
		Content []struct {
			Type        string        `json:"type"`
			Text        string        `json:"text"`
			Annotations []interface{} `json:"annotations"`
		} `json:"content"`
	} `json:"output"`
	ParallelToolCalls  bool        `json:"parallel_tool_calls"`
	PreviousResponseId interface{} `json:"previous_response_id"`
	Reasoning          struct {
		Effort  interface{} `json:"effort"`
		Summary interface{} `json:"summary"`
	} `json:"reasoning"`
	Store       bool    `json:"store"`
	Temperature float64 `json:"temperature"`
	Text        struct {
		Format struct {
			Type string `json:"type"`
		} `json:"format"`
	} `json:"text"`
	ToolChoice string        `json:"tool_choice"`
	Tools      []interface{} `json:"tools"`
	TopP       float64       `json:"top_p"`
	Truncation string        `json:"truncation"`
	Usage      struct {
		InputTokens        int `json:"input_tokens"`
		InputTokensDetails struct {
			CachedTokens int `json:"cached_tokens"`
		} `json:"input_tokens_details"`
		OutputTokens        int `json:"output_tokens"`
		OutputTokensDetails struct {
			ReasoningTokens int `json:"reasoning_tokens"`
		} `json:"output_tokens_details"`
		TotalTokens int `json:"total_tokens"`
	} `json:"usage"`
	User     interface{} `json:"user"`
	Metadata struct {
	} `json:"metadata"`
}

type FileInputResponse struct {
	Id                string      `json:"id"`
	Object            string      `json:"object"`
	CreatedAt         int         `json:"created_at"`
	Status            string      `json:"status"`
	Background        bool        `json:"background"`
	Error             interface{} `json:"error"`
	IncompleteDetails interface{} `json:"incomplete_details"`
	Instructions      interface{} `json:"instructions"`
	MaxOutputTokens   interface{} `json:"max_output_tokens"`
	MaxToolCalls      interface{} `json:"max_tool_calls"`
	Model             string      `json:"model"`
	Output            []struct {
		Id      string `json:"id"`
		Type    string `json:"type"`
		Status  string `json:"status"`
		Content []struct {
			Type        string        `json:"type"`
			Annotations []interface{} `json:"annotations"`
			Logprobs    []interface{} `json:"logprobs"`
			Text        string        `json:"text"`
		} `json:"content"`
		Role string `json:"role"`
	} `json:"output"`
	ParallelToolCalls  bool        `json:"parallel_tool_calls"`
	PreviousResponseId interface{} `json:"previous_response_id"`
	Reasoning          struct {
		Effort  interface{} `json:"effort"`
		Summary interface{} `json:"summary"`
	} `json:"reasoning"`
	ServiceTier string  `json:"service_tier"`
	Store       bool    `json:"store"`
	Temperature float64 `json:"temperature"`
	Text        struct {
		Format struct {
			Type string `json:"type"`
		} `json:"format"`
	} `json:"text"`
	ToolChoice  string        `json:"tool_choice"`
	Tools       []interface{} `json:"tools"`
	TopLogprobs int           `json:"top_logprobs"`
	TopP        float64       `json:"top_p"`
	Truncation  string        `json:"truncation"`
	Usage       struct {
		InputTokens        int `json:"input_tokens"`
		InputTokensDetails struct {
			CachedTokens int `json:"cached_tokens"`
		} `json:"input_tokens_details"`
		OutputTokens        int `json:"output_tokens"`
		OutputTokensDetails struct {
			ReasoningTokens int `json:"reasoning_tokens"`
		} `json:"output_tokens_details"`
		TotalTokens int `json:"total_tokens"`
	} `json:"usage"`
	User     interface{} `json:"user"`
	Metadata struct {
	} `json:"metadata"`
}

func (r *TextInputRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *ImageInputRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *FileInputRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

func (r *TextInputResponse) Unmarshal(b []byte) error {
	return json.Unmarshal(b, r)
}

func (r *ImageInputResponse) Unmarshal(b []byte) error {
	return json.Unmarshal(b, r)
}

func (r *FileInputResponse) Unmarshal(b []byte) error {
	return json.Unmarshal(b, r)
}
