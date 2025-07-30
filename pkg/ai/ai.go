package ai

import (
	"fmt"
	"github.com/go-logr/logr"
	"github.com/muraduiurie/kubegpt/pkg/ai/gpt"
	"github.com/muraduiurie/kubegpt/pkg/ai/helpers"
)

type AiClient interface {
	AskAi(rt gpt.RequestType, opts helpers.RequestOpts) (gpt.Responser, error)
}

func InitAiClient(client string, logger logr.Logger) (AiClient, error) {
	switch client {
	case Gpt:
		gptClient, err := gpt.GetGptConfig(logger.WithName("chatgpt"))
		if err != nil {
			return nil, fmt.Errorf("unable to get gpt configuration: %v", err)
		}
		return gptClient, nil
	case "":
		return nil, fmt.Errorf("AI client not specified")
	default:
		return nil, fmt.Errorf("unsupported AI client: %s", client)
	}
}
