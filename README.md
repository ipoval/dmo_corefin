## DMO Core Financials Clients

[![Maintainability](https://api.codeclimate.com/v1/badges/8b1f1cb35c53faa6c4b5/maintainability)](https://codeclimate.com/github/ipoval/dmocorefin/maintainability)

[goreportcard.com](https://goreportcard.com/report/github.com/ipoval/dmocorefin)

### Supported actions

Action | Exchange | Usage
:------|:------------|:-----------
stock chart img | NYSE | [link](https://github.com/ipoval/dmocorefin/tree/master#nyse-stock-chart-img)

###### NYSE. Stock chart img 
```go
provider := stockchartsDotCom.New()
img := New("SPY", provider)
encodedImg, err := img.GetImgBase64()
```
