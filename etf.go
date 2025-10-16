package coinglassclient

import (
	"context"
)

type ETFBitcoinItem struct {
	Ticker               string  `json:"ticker"`
	FundName             string  `json:"fund_name"`
	Region               string  `json:"region"`
	MarketStatus         string  `json:"market_status"`
	PrimaryExchange      string  `json:"primary_exchange"`
	CikCode              string  `json:"cik_code"`
	FundType             string  `json:"fund_type"`
	ListDate             int64   `json:"list_date"`
	SharesOutstanding    string  `json:"shares_outstanding"`
	AumUsd               string  `json:"aum_usd"`
	ManagementFeePercent string  `json:"management_fee_percent"`
	LastTradeTime        int64   `json:"last_trade_time"`
	LastQuoteTime        int64   `json:"last_quote_time"`
	VolumeQuantity       int     `json:"volume_quantity"`
	VolumeUsd            float64 `json:"volume_usd"`
	PriceChangeUsd       float64 `json:"price_change_usd"`
	PriceChangePercent   float64 `json:"price_change_percent"`
	AssetDetails         struct {
		NetAssetValueUsd       float64 `json:"net_asset_value_usd"`
		PremiumDiscountPercent float64 `json:"premium_discount_percent"`
		BtcHolding             float64 `json:"btc_holding"`
		BtcChangePercent24H    int     `json:"btc_change_percent_24h"`
		BtcChange24H           float64 `json:"btc_change_24h"`
		BtcChangePercent7D     float64 `json:"btc_change_percent_7d"`
		BtcChange7D            float64 `json:"btc_change_7d"`
		UpdateDate             string  `json:"update_date"`
	} `json:"asset_details"`
	UpdateTimestamp int64 `json:"update_timestamp"`
}

func (c *Client) BitcoinETFList(ctx context.Context) (result []ETFBitcoinItem, err error) {
	request := NewRequest(c.Url("/etf/bitcoin/list"), nil)
	_, err = c.doCall(ctx, request, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

type ETFEthereumItem struct {
	Ticker               string  `json:"ticker"`
	Name                 string  `json:"name"`
	Region               string  `json:"region"`
	MarketStatus         string  `json:"market_status"`
	PrimaryExchange      string  `json:"primary_exchange"`
	CikCode              string  `json:"cik_code"`
	Type                 string  `json:"type"`
	MarketCap            string  `json:"market_cap"`
	ListDate             int64   `json:"list_date"`
	SharesOutstanding    string  `json:"shares_outstanding"`
	Aum                  string  `json:"aum"`
	ManagementFeePercent string  `json:"management_fee_percent"`
	LastTradeTime        int64   `json:"last_trade_time"`
	LastQuoteTime        int64   `json:"last_quote_time"`
	VolumeQuantity       int     `json:"volume_quantity"`
	VolumeUsd            float64 `json:"volume_usd"`
	Price                float64 `json:"price"`
	PriceChange          float64 `json:"price_change"`
	PriceChangePercent   float64 `json:"price_change_percent"`
	AssetInfo            struct {
		Nav              float64 `json:"nav"`
		PremiumDiscount  float64 `json:"premium_discount"`
		HoldingQuantity  float64 `json:"holding_quantity"`
		ChangePercent1D  int     `json:"change_percent_1d"`
		ChangeQuantity1D int     `json:"change_quantity_1d"`
		ChangePercent7D  float64 `json:"change_percent_7d"`
		ChangeQuantity7D float64 `json:"change_quantity_7d"`
		Date             string  `json:"date"`
	} `json:"asset_info"`
	UpdateTime int64 `json:"update_time"`
}

func (c *Client) EthereumETFList(ctx context.Context) (result []ETFEthereumItem, err error) {
	request := NewRequest(c.Url("/etf/ethereum/list"), nil)
	_, err = c.doCall(ctx, request, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
