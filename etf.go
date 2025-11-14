package coinglassclient

import (
	"context"
)

type ETFBitcoinItem struct {
	Ticker               string   `json:"ticker"`
	FundName             string   `json:"fund_name"`
	Region               string   `json:"region"`
	MarketStatus         string   `json:"market_status"`
	PrimaryExchange      string   `json:"primary_exchange"`
	FundType             string   `json:"fund_type"`
	MarketCapUsd         string   `json:"market_cap_usd"`
	CikCode              *string  `json:"cik_code,omitempty"`
	ListDate             int64    `json:"list_date"`
	SharesOutstanding    string   `json:"shares_outstanding"`
	AumUsd               string   `json:"aum_usd"`
	ManagementFeePercent string   `json:"management_fee_percent"`
	LastTradeTime        int64    `json:"last_trade_time"`
	LastQuoteTime        int64    `json:"last_quote_time"`
	VolumeQuantity       float64  `json:"volume_quantity"`
	VolumeUsd            float64  `json:"volume_usd"`
	Price                *float64 `json:"price,omitempty"`
	PriceUsd             *float64 `json:"price_usd,omitempty"`
	PriceChange          *float64 `json:"price_change,omitempty"`
	PriceChangeUsd       *float64 `json:"price_change_usd,omitempty"`
	PriceChangePercent   *float64 `json:"price_change_percent,omitempty"`
	AssetDetails         struct {
		NetAssetValueUsd       float64 `json:"net_asset_value_usd"`
		PremiumDiscountPercent float64 `json:"premium_discount_percent"`
		HoldingQuantity        float64 `json:"holding_quantity"`
		ChangePercent24H       float64 `json:"change_percent_24h"`
		ChangeQuantity24H      float64 `json:"change_quantity_24h"`
		ChangePercent7D        float64 `json:"change_percent_7d"`
		ChangeQuantity7D       float64 `json:"change_quantity_7d"`
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
	Ticker               string   `json:"ticker"`
	FundName             string   `json:"fund_name,omitempty"`
	Region               string   `json:"region"`
	MarketStatus         string   `json:"market_status"`
	PrimaryExchange      string   `json:"primary_exchange"`
	CikCode              string   `json:"cik_code"`
	FundType             string   `json:"fund_type,omitempty"`
	MarketCapUsd         string   `json:"market_cap_usd,omitempty"`
	ListDate             int64    `json:"list_date"`
	SharesOutstanding    string   `json:"shares_outstanding"`
	AumUsd               string   `json:"aum_usd,omitempty"`
	ManagementFeePercent string   `json:"management_fee_percent"`
	LastTradeTime        int64    `json:"last_trade_time"`
	LastQuoteTime        int64    `json:"last_quote_time"`
	VolumeQuantity       float64  `json:"volume_quantity"`
	VolumeUsd            float64  `json:"volume_usd"`
	Price                *float64 `json:"price,omitempty"`
	PriceUsd             *float64 `json:"price_usd,omitempty"`
	PriceChange          *float64 `json:"price_change,omitempty"`
	PriceChangeUsd       *float64 `json:"price_change_usd,omitempty"`
	PriceChangePercent   *float64 `json:"price_change_percent,omitempty"`
	AssetDetails         struct {
		NetAssetValueUsd       *float64 `json:"net_asset_value_usd,omitempty"`
		PremiumDiscountPercent *float64 `json:"premium_discount_percent,omitempty"`
		HoldingQuantity        *float64 `json:"holding_quantity,omitempty"`
		ChangePercent24H       *float64 `json:"change_percent_24h,omitempty"`
		ChangeQuantity24H      *float64 `json:"change_quantity_24h,omitempty"`
		ChangePercent7D        *float64 `json:"change_percent_7d,omitempty"`
		ChangeQuantity7D       *float64 `json:"change_quantity_7d,omitempty"`
		UpdateDate             string   `json:"update_date,omitempty"`
	} `json:"asset_details"`
	UpdateTimestamp *int64 `json:"update_timestamp,omitempty"`
}

func (c *Client) EthereumETFList(ctx context.Context) (result []ETFEthereumItem, err error) {
	request := NewRequest(c.Url("/etf/ethereum/list"), nil)
	_, err = c.doCall(ctx, request, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}
