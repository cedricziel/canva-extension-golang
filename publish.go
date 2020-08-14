package canva

// PublishRequest is an incoming publish message
// see https://www.canva.com/developers/docs/publish-extensions/api/post-publish-resources-upload/
type PublishRequest struct {
	User    string         `json:"user"`
	Brand   string         `json:"brand"`
	Label   string         `json:"label"`
	Message string         `json:"message"`
	Parent  string         `json:"parent"`
	Assets  []PublishAsset `json:"assets"`
}

// PublishAsset represents an asset on the Canva side that's to be published
type PublishAsset struct {
	Name string `json:"name"`
	Type string `json:"type"`
	URL  string `json:"url"`
}
