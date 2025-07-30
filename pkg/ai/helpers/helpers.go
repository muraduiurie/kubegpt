package helpers

const (
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

	// roles
	User      string = "user"
	Assistant string = "assistant"
	System    string = "system"

	// request types
	FileRequestType  RequestType = "file"
	ImageRequestType RequestType = "image"
	TextRequestType  RequestType = "text"
)

type RequestType string

type RequestOpts struct {
	Message  string
	ImageUrl string
	FileUrl  string
	Model    string
	Role     string
}
