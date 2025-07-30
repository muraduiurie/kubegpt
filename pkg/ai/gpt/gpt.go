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
		Host:              "https://api.openai.com",
		Token:             os.Getenv("GPT_API_TOKEN"),
		ResponsesEndpoint: "v1/responses",
		Log:               logger,
	}

	return client, nil
}

func (g *Client) AskAi(opts *helpers.RequestOpts) (Responser, error) {
	var request Requester
	var response Responser

	if opts.Message == nil {
		return nil, fmt.Errorf("message not specified")
	}

	var model, role string
	var requestType helpers.RequestType
	if opts.Model == nil {
		model = helpers.Gpt4_1
	}
	if opts.Role == nil {
		role = helpers.UserRole
	}
	if opts.RequestType == nil {
		requestType = helpers.TextRequestType
	}

	var gptEndpoint string
	switch requestType {
	case helpers.FileRequestType:
		gptEndpoint = g.ResponsesEndpoint
		g.Log.Info("FileInput request received", "message", opts.Message, "url", opts.FileUrl)
		fir := FileInputRequest{
			Model: model,
			Input: []FileInputRequestInput{
				{
					Role: role,
					Content: []FileInputRequestContent{
						{
							Type:    InputFile,
							FileUrl: *opts.FileUrl,
						},
						{
							Type: InputText,
							Text: *opts.Message,
						},
					},
				},
			},
		}
		request = &fir
		response = &FileInputResponse{}
	case helpers.ImageRequestType:
		gptEndpoint = g.ResponsesEndpoint
		g.Log.Info("ImageInput request received", "message", opts.Message, "url", opts.ImageUrl)
		iir := ImageInputRequest{
			Model: model,
			Input: []ImageInputRequestInput{
				{
					Role: role,
					Content: []ImageInputRequestContent{
						{
							Type:     InputImage,
							ImageUrl: *opts.ImageUrl,
						},
						{
							Type: InputText,
							Text: *opts.Message,
						},
					},
				},
			},
		}
		request = &iir
		response = &ImageInputResponse{}
	case helpers.TextRequestType:
		gptEndpoint = g.ResponsesEndpoint
		g.Log.Info("Chat request received", "message", opts.Message)
		tir := TextInputRequest{
			Model: model,
			Input: *opts.Message,
		}
		request = &tir
		response = &TextInputResponse{}
	default:
		return nil, fmt.Errorf("unknown reuqest type: %s", requestType)
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
