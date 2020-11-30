package exchangenyse

var MarketIndicators = []struct {
	Symbol     string
	Desc       string
	HostOrigin string
}{
	{Symbol: "$SPX", Desc: "S&P500", HostOrigin: "stockcharts.com"},
	{Symbol: "$VIX", Desc: "Volatility index", HostOrigin: "stockcharts.com"},
	{Symbol: "$SPX:$GOLD", Desc: "S&P500/Gold", HostOrigin: "stockcharts.com"},
	{Symbol: "$CPCE", Desc: "Options equity Put/Call ratio", HostOrigin: "stockcharts.com"},
	{Symbol: "CNN_FEAR_AND_GREED_INDEX", Desc: "CNN fear and greed indicator", HostOrigin: "money.cnn.com"},
}
