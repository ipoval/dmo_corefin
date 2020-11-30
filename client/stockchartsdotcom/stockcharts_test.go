package stockchartsdotcom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetImgBase64(t *testing.T) {
	is := assert.New(t)
	sp500 := "SPY"
	encodedImg, err := testingChartImg.GetChartImgBase64(sp500)

	is.NoError(err)
	is.NotEmpty(encodedImg)
	// t.Log(encodedImg)
}
