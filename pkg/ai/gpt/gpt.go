package gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-logr/logr"
	"github.com/muraduiurie/kubegpt/pkg/ai/helpers"
	"io"
	"net/http"
	"net/url"
	"os"
)

func GetGptConfig(logger logr.Logger) (*Client, error) {
	if os.Getenv("GPT_API_TOKEN") == "" {
		return nil, fmt.Errorf("GPT_API_TOKEN is not set")
	}

	client := &Client{
		Host:            "https://api.openai.com",
		Token:           os.Getenv("GPT_API_TOKEN"),
		ChatEndpoint:    "v1/chat/completions",
		FileUrlEndpoint: "v1/responses",
		Log:             logger,
	}

	return client, nil
}

func (g *Client) AskAi(rt helpers.RequestType, opts helpers.RequestOpts) (Responser, error) {
	var request Requester
	var response Responser

	var gptEndpoint string
	switch rt {
	case helpers.FileRequestType:
		gptEndpoint = g.FileUrlEndpoint
		g.Log.Info("FileInput request received", "message", opts.Message, "url", opts.FileUrl)
		fir := FileInputRequest{
			Model: opts.Model,
			Input: []struct {
				Role    string `json:"role"`
				Content []struct {
					Type    string `json:"type"`
					Text    string `json:"text,omitempty"`
					FileUrl string `json:"file_url,omitempty"`
				} `json:"content"`
			}{
				{
					Role: opts.Role,
					Content: []struct {
						Type    string `json:"type"`
						Text    string `json:"text,omitempty"`
						FileUrl string `json:"file_url,omitempty"`
					}{
						{
							Type:    InputFile,
							FileUrl: opts.FileUrl,
						},
						{
							Type: InputText,
							Text: opts.Message,
						},
					},
				},
			},
		}
		request = &fir
		response = &FileInputResponse{}
	case helpers.ImageRequestType:
		gptEndpoint = g.FileUrlEndpoint
		g.Log.Info("ImageInput request received", "message", opts.Message, "url", opts.ImageUrl)
		iir := ImageInputRequest{
			Model: opts.Model,
			Input: []struct {
				Role    string `json:"role"`
				Content []struct {
					Type     string `json:"type"`
					Text     string `json:"text,omitempty"`
					ImageUrl string `json:"image_url,omitempty"`
				} `json:"content"`
			}{
				{
					Role: opts.Role,
					Content: []struct {
						Type     string `json:"type"`
						Text     string `json:"text,omitempty"`
						ImageUrl string `json:"image_url,omitempty"`
					}{
						{
							Type:     InputImage,
							ImageUrl: opts.ImageUrl,
						},
						{
							Type: InputText,
							Text: opts.Message,
						},
					},
				},
			},
		}
		request = &iir
		response = &ImageInputResponse{}
	case helpers.TextRequestType:
		gptEndpoint = g.ChatEndpoint
		g.Log.Info("Chat request received", "message", opts.Message)
		tir := TextInputRequest{
			Model: opts.Model,
			Input: opts.Message,
		}
		request = &tir
		response = &TextInputResponse{}
	}

	jsonRequestBody, err := request.Marshal()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}

	host, err := url.JoinPath(g.Host, gptEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to join url")
	}
	g.Log.Info("request", "json", string(jsonRequestBody), "host", host)
	req, err := http.NewRequest(http.MethodPost, host, bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", g.Token))
	req.Header.Set("Content-Type", "application/json")

	// create a new client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}

	// handle response
	defer resp.Body.Close()

	// read the response body
	jsonResponseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code %d: %s", resp.StatusCode, string(jsonResponseBody))
	}

	err = json.Unmarshal(jsonResponseBody, response)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %v", err)
	}

	g.Log.Info("Got response", "response", response)

	return response, nil
}
