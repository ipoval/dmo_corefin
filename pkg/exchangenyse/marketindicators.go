package exchangenyse

var MarketIndicators = []struct {
	Symbol     string
	Desc       string
	HostOrigin string
	Source     string
}{
	{Symbol: "$SPX", Desc: "S&P500", HostOrigin: "stockcharts.com"},
	{Symbol: "$VIX", Desc: "Volatility index", HostOrigin: "stockcharts.com"},
	{Symbol: "$SPX:$GOLD", Desc: "S&P500/Gold", HostOrigin: "stockcharts.com"},
	{Symbol: "GLD", Desc: "Gold shares", HostOrigin: "stockcharts.com"},
	{Symbol: "!GT50GDX", Desc: "% of GDX above 50dma", HostOrigin: "stockcharts.com"},
	{Symbol: "$CPCE", Desc: "Options equity Put/Call ratio", HostOrigin: "stockcharts.com"},
	{Symbol: "CNN_FEAR_AND_GREED_INDEX", Desc: "CNN fear and greed indicator", HostOrigin: "money.cnn.com"},
	{Symbol: "$SPXA200R", Desc: "NYSE Percent of stocks above 200DMA", HostOrigin: "stockcharts.com", Source: "https://stockcharts.com/h-sc/ui?s=$SPXA200R"},
}
