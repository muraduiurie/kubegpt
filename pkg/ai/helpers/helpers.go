package helpers

type AiOpts struct {
	Message    string
	Role       string
	Model      string
	FileUpload *FileUpload
	FileUrl    *Url
	ImageUrl   *Url
}

type Url struct {
	Hostname string
}

type FileUpload struct{}
