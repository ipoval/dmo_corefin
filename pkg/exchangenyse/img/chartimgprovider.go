package img

// ChartImageProvider ....
type ChartImageProvider interface {
  GetChartImgBase64(string) (string, error)
}
