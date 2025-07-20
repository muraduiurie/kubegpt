package helpers

type AiOpts struct {
	Message    string
	Role       string
	Model      string
	FileUpload *FileUpload
	FileUrl    *FileUrl
}

type FileUrl struct {
	Url string
}

type FileUpload struct{}
