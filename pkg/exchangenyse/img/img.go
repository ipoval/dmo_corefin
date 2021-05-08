package img

import (
	"github.com/ipoval/dmocorefin/client/cnndotcom"
	"github.com/ipoval/dmocorefin/client/stockchartsdotcom"
	"github.com/ipoval/dmocorefin/pkg/exchangenyse"
)

type Img struct {
	Symbol string
	ChartImageProvider
}

// New ...
func New(indicator exchangenyse.Indicator) *Img {
	img := &Img{Symbol: indicator.GetSymbol()}
	img.setChartImageProvider(indicator)
	return img
}

// GetImgBase64 ...
func (i *Img) GetImgBase64() (string, error) {
	return i.GetChartImgBase64(i.Symbol)
}

func (i *Img) setChartImageProvider(indicator exchangenyse.Indicator) {
	switch indicator.GetOrigin() {
	case exchangenyse.OriginStockCharts:
		i.ChartImageProvider = stockchartsdotcom.New()
	case exchangenyse.OriginCnnFearAndGreed:
		i.ChartImageProvider = cnndotcom.New()
	default:
		panic("unexpected chart image provider type")
	}
}
