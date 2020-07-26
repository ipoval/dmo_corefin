package stockchartsdotcom

import (
	"time"

	"github.com/lucperkins/rek"
)

const (
	Referer = "https://stockcharts.dot.com"
)

var (
	httpTimeout = rek.Timeout(5 * time.Second)
)
