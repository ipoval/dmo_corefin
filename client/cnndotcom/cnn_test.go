package cnndotcom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetChartImgBase64GetImgBase64(t *testing.T) {
	is := assert.New(t)
	encodedImg, err := New().GetChartImgBase64("CNN_FEAR_AND_GREED_INDEX")

	is.NoError(err)
	is.NotEmpty(encodedImg)
	// t.Log(encodedImg)
}
