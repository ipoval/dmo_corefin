// Package img:
package img

// ChartImage ....
type ChartImage interface {
  GetChartImgBase64(string) (string, error)
}
