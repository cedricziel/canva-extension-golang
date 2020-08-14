package canva

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanDeserializeIncomingPublishRequest(t *testing.T) {
	const request = `{
		"user": "AUQ2RUzug9pEvgpK9lL2qlpRsIbn1Vy5GoEt1MaKRE=",
		"brand": "AUQ2RUxiRj966Wsvp7oGrz33BnaFmtq4ftBeLCSHf8=",
		"label": "PUBLISH",
		"assets": [
		  {
			"url": "https://s3.amazonaws.com/.../49-04fa92cbfbf8.jpg",
			"type": "JPG",
			"name": "0001-144954.jpg"
		  }
		],
		"parent": "12345"
	}`

	var publishRequest = PublishRequest{}
	err := json.Unmarshal([]byte(request), &publishRequest)

	assert.Nil(t, err)

	assert.EqualValues(t, "AUQ2RUzug9pEvgpK9lL2qlpRsIbn1Vy5GoEt1MaKRE=", publishRequest.User)
	assert.EqualValues(t, "AUQ2RUxiRj966Wsvp7oGrz33BnaFmtq4ftBeLCSHf8=", publishRequest.Brand)
	assert.EqualValues(t, "PUBLISH", publishRequest.Label)
	assert.EqualValues(t, "12345", publishRequest.Parent)

	assert.Len(t, publishRequest.Assets, 1)

	assert.EqualValues(t, "0001-144954.jpg", publishRequest.Assets[0].Name)
	assert.EqualValues(t, "JPG", publishRequest.Assets[0].Type)
	assert.EqualValues(t, "https://s3.amazonaws.com/.../49-04fa92cbfbf8.jpg", publishRequest.Assets[0].URL)
}
