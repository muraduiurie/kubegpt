package gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-logr/logr"
	"io"
	"net/http"
	"os"
	"strings"
)

func GetGptConfig(logger logr.Logger) (*Client, error) {
	if os.Getenv("GPT_API_TOKEN") == "" {
		return nil, fmt.Errorf("GPT_API_TOKEN is not set")
	}

	client := &Client{
		Host:         "https://api.openai.com",
		Token:        os.Getenv("GPT_API_TOKEN"),
		ChatEndpoint: "v1/chat/completions",
		Log:          logger,
	}

	return client, nil
}

func (g *Client) AskAi(message string, role string, model string) (string, error) {
	g.Log.Info("AskAi", "message", message, "role", role, "model", model)
	request := ChatRequest{
		Model:       model,
		Temperature: ModerateVariability,
		Messages: []RequestMessage{
			{
				Role:    role,
				Content: message,
			},
		},
	}

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, strings.Join([]string{g.Host, g.ChatEndpoint}, "/"), bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", g.Token))
	req.Header.Set("Content-Type", "application/json")

	// create a new client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}

	// handle response
	defer resp.Body.Close()

	// read the response body
	jsonBodyResponse, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request failed with status code %d: %s", resp.StatusCode, string(jsonBodyResponse))
	}

	newResponse := ChatResponse{}
	err = json.Unmarshal(jsonBodyResponse, &newResponse)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal response: %v", err)
	}

	g.Log.Info("Got response", "response", newResponse)

	return newResponse.Choices[0].Message.Content, nil
}
