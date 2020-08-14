package canva

const (
	ContentContainer string = "CONTAINER"
	ContentImage     string = "IMAGE"
	ContentEmbed     string = "EMBED"
)

type ContentThumbnail struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type Content struct {
	Type        string           `json:"type"`
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Thumbnail   ContentThumbnail `json:"thumbnail"`
	URL         string           `json:"url"`
	ContentType string           `json:"contentType"`
}

func NewImageContent() Content {
	return Content{Type: ContentImage}
}

func NewContainerContent(name string, id string) Content {
	return Content{Type: ContentContainer, Name: name, ID: id}
}

// ContentRequest is an incoming request from canva for you to provide Content
// for the in-product search
type ContentRequest struct {
	Brand        string `json:"brand"`
	Label        string `json:"label"`
	Limit        int    `json:"limit"`
	Type         string `json:"type"`
	User         string `json:"user"`
	ContainerID  string `json:"containerId"`
	Continuation string `json:"continuation"`
	Locale       string `json:"locale"`
	Query        string `json:"query"`
}

type ContentResponse struct {
	Type         string `json:"type"`
	Resources    string `json:"resources"`
	Continuation string `json:"continuation"`
}
