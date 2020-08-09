package canva

import (
	"encoding/json"
	"net/http"
)

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

func NewSuccessResponse() CanvaResponse {
	return CanvaResponse{Type: "SUCCESS"}
}

// WriteSuccessResponse will write a correct success response
// Canva is really picky about response codes and expects
// a status of 200 and an inline success.
func WriteSuccessResponse(w http.ResponseWriter) {
	json, _ := json.Marshal(NewSuccessResponse())

	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	w.Write(json)
}
