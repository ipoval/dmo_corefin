## DMO Core Financials Clients

### Supported actions

Action | Exchange | Usage
:------|:------------|:-----------
stock chart img | NYSE | [link](https://github.com/ipoval/dmo_corefin/tree/master#nyse-stock-chart-img)

###### NYSE. Stock chart img 
```go
provider := stockcharts.New()
img := New("SPY", provider)
encodedImg, err := img.GetImgBase64()
```
