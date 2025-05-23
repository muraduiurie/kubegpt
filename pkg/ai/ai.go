package ai

import (
	"fmt"
	"github.com/go-logr/logr"
	"github.com/muraduiurie/kubegpt/pkg/ai/deepseek"
	"github.com/muraduiurie/kubegpt/pkg/ai/gpt"
)

type AiClient interface {
	AskAi(message string, role string, model string) (string, error)
}

func InitAiClient(client string, logger logr.Logger) (AiClient, error) {
	switch client {
	case Gpt:
		gptClient, err := gpt.GetGptConfig(logger.WithName("chatgpt"))
		if err != nil {
			return nil, fmt.Errorf("unable to get gpt configuration: %v", err)
		}
		return gptClient, nil
	case DeepSeek:
		deepSeekClient, err := deepseek.GetDeepSeekConfig(logger.WithName("deepseek"))
		if err != nil {
			return nil, fmt.Errorf("unable to get deepseek configuration: %v", err)
		}
		return deepSeekClient, nil
	case "":
		return nil, fmt.Errorf("AI client not specified")
	default:
		return nil, fmt.Errorf("unsupported AI client: %s", client)
	}
}
