package canva

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanDeserializeContainerRequest(t *testing.T) {
	const request = `{
		"user": "AVnWYDnv7TtPvDSMSvngaIJP6plu860IYLU3rbCfGJ0=",
		"brand": "AVnWYDnOvjpbzV2-Rax9bSeEAIw3cEI6onnyvyDf5ac=",
		"label": "CONTENT",
		"limit": 8,
		"locale": "en-GB",
		"type": "CONTAINER",
		"containerId": "myContainer"
	}`

	var contentRequest = ContentRequest{}
	err := json.Unmarshal([]byte(request), &contentRequest)

	assert.Nil(t, err)
	assert.EqualValues(t, "AVnWYDnv7TtPvDSMSvngaIJP6plu860IYLU3rbCfGJ0=", contentRequest.User)
	assert.EqualValues(t, "AVnWYDnOvjpbzV2-Rax9bSeEAIw3cEI6onnyvyDf5ac=", contentRequest.Brand)
	assert.EqualValues(t, "CONTENT", contentRequest.Label)
	assert.EqualValues(t, 8, contentRequest.Limit)
	assert.EqualValues(t, "en-GB", contentRequest.Locale)
	assert.EqualValues(t, "CONTAINER", contentRequest.Type)
	assert.EqualValues(t, "myContainer", contentRequest.ContainerID)
}
