package canva

type CanvaAsset struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Url  string `json:"url"`
}

type CanvaMessage struct {
	User    string
	Brand   string
	Label   string
	Message string
	Assets  []CanvaAsset `json:"assets"`
}

type CanvaResponse struct {
	Type string `json:"type"`
}
