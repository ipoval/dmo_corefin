package img

type Img struct {
	Ticker string
	ChartImage
}

// New ...
func New(ticker string, chartImg ChartImage) *Img {
	return &Img{
		Ticker:     ticker,
		ChartImage: chartImg,
	}
}

// GetImgBase64 ...
func (i *Img) GetImgBase64() (string, error) {
	return i.GetChartImgBase64(i.Ticker)
}
