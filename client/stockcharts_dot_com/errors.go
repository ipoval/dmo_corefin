package stockcharts_dot_com

import (
	"github.com/pkg/errors"
)

var ErrChartImgStatusCode = errors.New("not 2xx status code fetching chart image")
