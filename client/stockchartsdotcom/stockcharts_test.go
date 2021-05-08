package stockchartsdotcom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetImgBase64(t *testing.T) {
	is := assert.New(t)
	encodedImg, err := New().GetChartImgBase64("SPY")

	is.NoError(err)
	is.NotEmpty(encodedImg)
	// t.Log(encodedImg)
}
