package stockcharts_dot_com

import (
  "encoding/base64"
  "net/http"

  "github.com/lucperkins/rek"
)

func NewClientStockchartsDotCom() *ClientStockchartsDotCom {
  return &ClientStockchartsDotCom{
    Referer: Referer,
  }
}

type ClientStockchartsDotCom struct {
  Referer string
}

// Get base64 encoded bytes for chart image by security ticker
func (c *ClientStockchartsDotCom) GetChartImgBase64(ticker string) (string, error) {
  httpHeaders := map[string]string{"Referer": Referer,}
  url := c.getChartImgURL(ticker)

  resp, err := rek.Get(
    url,
    rek.Headers(httpHeaders),
    rek.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.106 Safari/537.36"),
    httpTimeout,
  )

  if err != nil {
    return "", err
  }

  if resp.StatusCode() != http.StatusOK {
    return "", ErrChartImgStatusCode
  }

  body, err := rek.BodyAsBytes(resp.Body())

  if err != nil {
    return "", err
  }

  return base64.StdEncoding.EncodeToString(body), nil
}

func (c *ClientStockchartsDotCom) getChartImgURL(ticker string) string {
  return `https://stockcharts.com/c-sc/sc?s=` + ticker + `&p=W&b=5&g=0&i=t1041368475c`
}
