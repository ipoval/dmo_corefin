## DMO Core Financials Clients

[![Maintainability](https://api.codeclimate.com/v1/badges/5c9a9a905fae29c790b5/maintainability)](https://codeclimate.com/github/ipoval/dmo_corefin/maintainability)

[goreportcard.com](https://goreportcard.com/report/github.com/ipoval/dmo_corefin)

### Supported actions

Action | Exchange | Usage
:------|:------------|:-----------
stock chart img | NYSE | [link](https://github.com/ipoval/dmo_corefin/tree/master#nyse-stock-chart-img)

###### NYSE. Stock chart img 
```go
provider := stockchartsDotCom.New()
img := New("SPY", provider)
encodedImg, err := img.GetImgBase64()
```
