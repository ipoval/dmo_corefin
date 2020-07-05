package exchange_nyse

var MarketTickersIndicators = []struct {
  Symbol string
  Desc   string
}{
  {Symbol: "$SPX", Desc: "S&P500"},
  {Symbol: "$VIX", Desc: "Volatility index"},
}
