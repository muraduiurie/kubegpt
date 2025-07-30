package helpers

type AiOpts struct {
	Message    string
	Role       string
	Model      string
	FileInput  *Url
	ImageInput *Url
	TextInput  *TextMessage
}

type Url struct {
	Hostname string
}

type TextMessage struct {
	Message string
}
