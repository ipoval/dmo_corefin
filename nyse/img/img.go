package img

type ProviderChartImage interface {
  GetChartImgBase64(string) (string, error)
}

type Img struct {
  Ticker string
  ProviderChartImage
}

func New(ticker string, chartImgProvider ProviderChartImage) *Img {
  return &Img{
    Ticker:             ticker,
    ProviderChartImage: chartImgProvider,
  }
}

func (i *Img) GetImgBase64() (string, error) {
  return i.GetChartImgBase64(i.Ticker)
}
