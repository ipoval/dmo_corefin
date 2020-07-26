package exchangenyse

var MarketIndicators = []struct {
	Symbol string
	Desc   string
}{
	{Symbol: "$SPX", Desc: "S&P500"},
	{Symbol: "$VIX", Desc: "Volatility index"},
	{Symbol: "$SPX:$GOLD", Desc: "S&P500/Gold"},
	{Symbol: "$CPCE", Desc: "Options equity Put/Call ratio"},
}