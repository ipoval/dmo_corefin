package stockchartsdotcom

import (
	"os"
	"testing"

	"github.com/ipoval/dmocorefin/pkg/exchangenyse/img"
)

var testProviderChartImg img.ChartImage

// Called once for all tests
func TestMain(m *testing.M) {
	testProviderChartImg = New()
	os.Exit(m.Run())
}
