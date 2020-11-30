package client

import (
	"encoding/base64"
	"net/http"
	"time"

	gasrest "github.com/ipoval/goactivesupport/rest"
	"github.com/lucperkins/rek"
)

var (
	httpTimeout = rek.Timeout(5 * time.Second)
)

func RemoteImgToBase64(url, referer string) (string, error) {
	resp, err := rek.Get(
		url,
		rek.Headers(map[string]string{
			"Referer": referer,
		}),
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
