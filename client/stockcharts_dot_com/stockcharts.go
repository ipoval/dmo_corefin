package stockcharts_dot_com

import (
	"encoding/base64"
	"fmt"
	"net/http"

	gasrest "github.com/ipoval/goactivesupport/rest"
	"github.com/lucperkins/rek"
)

// Initialize new client
func New() *Client {
	return &Client{
		Referer: Referer,
	}
}

type Client struct {
	Referer string
}

// Get base64 encoded bytes for chart image by security ticker
func (c *Client) GetChartImgBase64(ticker string) (string, error) {
	httpHeaders := map[string]string{"Referer": Referer}
	url := c.getChartImgURL(ticker)

	resp, err := rek.Get(
		url,
		rek.Headers(httpHeaders),
		rek.UserAgent(gasrest.UserAgent.Chrome()),
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

func (c *Client) getChartImgURL(ticker string) string {
	return fmt.Sprintf(
		`https://stockcharts.com/c-sc/sc?s=%s&p=W&b=5&g=0&i=p77260315374&r=1593988736282`,
		ticker,
	)
}
