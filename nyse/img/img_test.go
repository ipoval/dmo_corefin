package img

import (
  "testing"

  "github.com/stretchr/testify/assert"
)

func TestGetImgBase64(t *testing.T) {
  is := assert.New(t)
  img := New("SPY", testProviderChartImg)
  encodedImg, err := img.GetImgBase64()

  is.NoError(err)
  is.NotEmpty(encodedImg)
  // t.Log(encodedImg)
}
