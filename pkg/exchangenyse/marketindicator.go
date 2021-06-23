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
	{
		Symbol: "$SPX",
		Desc:   "S&P500",
		Origin: OriginStockCharts,
	},
	{
		Symbol: "$VIX",
		Desc:   "Volatility index",
		Origin: OriginStockCharts,
	},
	{
		Symbol: "$SPX:$GOLD",
		Desc:   "S&P500/Gold",
		Origin: OriginStockCharts,
	},
	{
		Symbol: "GLD",
		Desc:   "Gold shares",
		Origin: OriginStockCharts,
	},
	{
		Symbol: "!GT50GDX",
		Desc:   "% of GDX above 50dma",
		Origin: OriginStockCharts,
	},
	{
		Symbol: "$CPCE",
		Desc:   "Options equity Put/Call ratio",
		Origin: OriginStockCharts,
	},
	{
		Symbol: "CNN_FEAR_AND_GREED_INDEX",
		Desc:   "CNN fear and greed indicator",
		Origin: OriginCnnFearAndGreed,
	},
	{
		Symbol: "$SPXA200R",
		Desc:   "NYSE Percent of stocks above 200DMA",
		Origin: OriginStockCharts,
		Source: "https://stockcharts.com/h-sc/ui?s=$SPXA200R",
	},
	{
		Symbol: "$COPPER",
		Desc:   "Copper continuous contract price",
		Origin: OriginStockCharts,
		Source: "https://stockcharts.com/h-sc/ui?s=$COPPER",
	},
	{
		Symbol: "$PLAT",
		Desc:   "Platinum continuous contract price",
		Origin: OriginStockCharts,
		Source: "https://stockcharts.com/h-sc/ui?s=$PLAT",
	},
	{
		Symbol: "$PALL",
		Desc:   "Palladium continuous contract price",
		Origin: OriginStockCharts,
		Source: "https://stockcharts.com/h-sc/ui?s=$PALL",
	},
	{
		Symbol: "$BRENT",
		Desc:   "Brent crude oil continuous contract price",
		Origin: OriginStockCharts,
		Source: "https://stockcharts.com/h-sc/ui?s=$BRENT",
	},
	{
		Symbol: "$NATGAS",
		Desc:   "Natural gas continuous contract price",
		Origin: OriginStockCharts,
		Source: "https://stockcharts.com/h-sc/ui?s=$NATGAS",
	},
	{
		Symbol: "$LUMBER",
		Desc:   "Lumber continuous contract price",
		Origin: OriginStockCharts,
		Source: "https://stockcharts.com/h-sc/ui?s=$LUMBER",
	},
	{
		Symbol: "$SUGAR",
		Desc:   "Sugar continuous contract price",
		Origin: OriginStockCharts,
		Source: "https://stockcharts.com/h-sc/ui?s=$SUGAR",
	},
	{
		Symbol: "$CORN",
		Desc:   "Corn continuous contract price",
		Origin: OriginStockCharts,
		Source: "https://stockcharts.com/h-sc/ui?s=$CORN",
	},
	{
		Symbol: "$WHEAT",
		Desc:   "Wheat continuous contract price",
		Origin: OriginStockCharts,
		Source: "https://stockcharts.com/h-sc/ui?s=$WHEAT",
	},
}
