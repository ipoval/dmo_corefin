package cnndotcom

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gocolly/colly/v2"
	gasrest "github.com/ipoval/goactivesupport/rest"

	. "github.com/ipoval/dmocorefin/client"
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
	var localImgUrl string

	// Scrape image url from the CNN page
	webScraper := colly.NewCollector(colly.UserAgent(gasrest.UserAgent.Chrome()))
	webScraper.OnHTML("#needleChart", func(e *colly.HTMLElement) {
		imgUrl := e.Attr("style")
		imgUrl = strings.TrimPrefix(imgUrl, "background-image:url('")
		imgUrl = strings.TrimSuffix(imgUrl, "');")
		localImgUrl = imgUrl
	})
	if err := webScraper.Visit(FearAndGreedIndexUrl); err != nil {
		fmt.Println("error webScraper: %s", err.Error())
		return "", ErrChartImgStatusCode
	}

	if err := validateFearAndGreedIndexImgUrl(localImgUrl); err != nil {
		fmt.Println("invalid url: ", localImgUrl)
		return "", ErrChartImgStatusCode
	}

	return RemoteImgToBase64(localImgUrl, c.Referer)
}

func validateFearAndGreedIndexImgUrl(imgUrl string) error {
	validate := validator.New()
	errs := validate.Var(imgUrl, "required,url")
	return errs
}
