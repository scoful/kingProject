package coingecko

type GetPrice struct {
	CoinIds     string `v:"required#请输入币id,如果有多个用,分隔"`
	VsCurrencie string `v:"required#请输入货币代码，如cny，usd"`
}
