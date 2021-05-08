package img

import (
	"github.com/ipoval/dmocorefin/client/cnndotcom"
	"github.com/ipoval/dmocorefin/client/stockchartsdotcom"
	"os"
	"testing"
)

var (
	testChartImgProviderCnnDotCom         ChartImageProvider
	testChartImgProviderStockChartsDotCom ChartImageProvider
)

// Called once for all tests
func TestMain(m *testing.M) {
	testChartImgProviderStockChartsDotCom = stockchartsdotcom.New()
	testChartImgProviderCnnDotCom = cnndotcom.New()

	os.Exit(m.Run())
}
