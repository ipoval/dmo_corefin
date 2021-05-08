package stockchartsdotcom

import (
	"fmt"

	. "github.com/ipoval/dmocorefin/client"
)

const (
	Referer = "https://stockcharts.dot.com"
)

type Client struct {
	Referer string
}

// Constructor
func New() *Client {
	return &Client{Referer: Referer}
}

// Get base64 encoded bytes for chart image by security ticker
func (c *Client) GetChartImgBase64(ticker string) (string, error) {
	return RemoteImgToBase64(c.getChartImgURL(ticker), c.Referer)
}

func (c *Client) getChartImgURL(ticker string) string {
	return fmt.Sprintf(
		`https://stockcharts.com/c-sc/sc?s=%s&p=W&b=5&g=0&i=p77260315374&r=1593988736282`,
		ticker,
	)
}
