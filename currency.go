package coinbasecommerce

// Currency also known as currency code; may or may not follow ISO 4217.
type Currency string

// Currency constants.
const (
	CurrencyBitcoin     Currency = "BTC"
	CurrencyBitcoinCash Currency = "BCH"
	CurrencyEthereum    Currency = "ETH"
	CurrencyLitecoin    Currency = "LTC"
	CurrencyUSDCoin     Currency = "USDC"
	CurrencyDai         Currency = "DAI"
)
