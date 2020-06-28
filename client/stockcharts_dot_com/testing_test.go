package stockcharts_dot_com

import (
  "os"
  "testing"

  "github.com/ipoval/dmo_corefin/nyse/img"
)

var testProviderChartImg img.ProviderChartImage

// Called once for all tests
func TestMain(m *testing.M) {
  testProviderChartImg = New()
  os.Exit(m.Run())
}
