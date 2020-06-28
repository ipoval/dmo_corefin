## DMO Core Financials Clients

### Supported actions

Action | Exchange | Usage
:------|:------------|:-----------
stock chart img | NYSE | `#NYSE. Stock chart img`

###### NYSE. Stock chart img 
```go
provider := stockcharts.New()
img := New("SPY", provider)
encodedImg, err := img.GetImgBase64()
```
