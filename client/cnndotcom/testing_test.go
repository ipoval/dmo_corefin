package cnndotcom

import (
	"os"
	"testing"

	"github.com/ipoval/dmocorefin/pkg/exchangenyse/img"
)

var testingChartImg img.ChartImage

// Called once for all tests
func TestMain(m *testing.M) {
	testingChartImg = New()
	os.Exit(m.Run())
}
