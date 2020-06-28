package img

import (
  "os"
  "testing"

  stockcharts "github.com/ipoval/dmo_corefin/client/stockcharts_dot_com"
)

var testProviderChartImg ProviderChartImage

// Called once for all tests
func TestMain(m *testing.M) {
  testProviderChartImg = stockcharts.NewClientStockchartsDotCom()
  os.Exit(m.Run())
}
