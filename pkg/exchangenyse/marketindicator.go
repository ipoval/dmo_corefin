package exchangenyse

const (
	OriginStockCharts     = "stockcharts.com"
	OriginCnnFearAndGreed = "money.cnn.com"
)

type Indicator interface {
	GetSymbol() string
	GetOrigin() string
}

type MarketIndicator struct {
	Symbol string
	Desc   string
	Origin string
	Source string
}

func (i MarketIndicator) GetSymbol() string {
	return i.Symbol
}

func (i MarketIndicator) GetOrigin() string {
	return i.Origin
}

var MarketIndicators = []MarketIndicator{
	{Symbol: "$SPX", Desc: "S&P500", Origin: OriginStockCharts},
	{Symbol: "$VIX", Desc: "Volatility index", Origin: OriginStockCharts},
	{Symbol: "$SPX:$GOLD", Desc: "S&P500/Gold", Origin: OriginStockCharts},
	{Symbol: "GLD", Desc: "Gold shares", Origin: OriginStockCharts},
	{Symbol: "!GT50GDX", Desc: "% of GDX above 50dma", Origin: OriginStockCharts},
	{Symbol: "$CPCE", Desc: "Options equity Put/Call ratio", Origin: OriginStockCharts},
	{Symbol: "CNN_FEAR_AND_GREED_INDEX", Desc: "CNN fear and greed indicator", Origin: OriginCnnFearAndGreed},
	{Symbol: "$SPXA200R", Desc: "NYSE Percent of stocks above 200DMA", Origin: OriginStockCharts, Source: "https://stockcharts.com/h-sc/ui?s=$SPXA200R"},
}
