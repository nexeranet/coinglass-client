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
	CikCode              string   `json:"cik_code"`
	FundType             string   `json:"fund_type"`
	ListDate             int64    `json:"list_date"`
	SharesOutstanding    string   `json:"shares_outstanding"`
	AumUsd               string   `json:"aum_usd"`
	ManagementFeePercent string   `json:"management_fee_percent"`
	LastTradeTime        int64    `json:"last_trade_time"`
	LastQuoteTime        int64    `json:"last_quote_time"`
	VolumeQuantity       float64  `json:"volume_quantity"`
	VolumeUsd            float64  `json:"volume_usd"`
	Price                *float64 `json:"price,omitempty"`
	PriceChange          *float64 `json:"price_change,omitempty"`
	PriceChangeUsd       *float64 `json:"price_change_usd,omitempty"`
	PriceChangePercent   *float64 `json:"price_change_percent,omitempty"`
	AssetDetails         struct {
		NetAssetValueUsd       *float64 `json:"net_asset_value_usd,omitempty"`
		PremiumDiscountPercent *float64 `json:"premium_discount_percent,omitempty"`
		BtcHolding             *float64 `json:"btc_holding,omitempty"`
		BtcChangePercent24H    *float64 `json:"btc_change_percent_24h,omitempty"`
		BtcChange24H           *float64 `json:"btc_change_24h,omitempty"`
		BtcChangePercent7D     *float64 `json:"btc_change_percent_7d,omitempty"`
		BtcChange7D            *float64 `json:"btc_change_7d,omitempty"`
		HoldingQuantity        *float64 `json:"holding_quantity,omitempty"`
		ChangePercent24H       *float64 `json:"change_percent_24h,omitempty"`
		ChangeQuantity24H      *float64 `json:"change_quantity_24h,omitempty"`
		ChangePercent7D        *float64 `json:"change_percent_7d,omitempty"`
		ChangeQuantity7D       *float64 `json:"change_quantity_7d,omitempty"`
		UpdateDate             string   `json:"update_date"`
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
	Name                 string   `json:"name"`
	Region               string   `json:"region"`
	MarketStatus         string   `json:"market_status"`
	PrimaryExchange      string   `json:"primary_exchange"`
	CikCode              string   `json:"cik_code"`
	Type                 string   `json:"type"`
	MarketCap            string   `json:"market_cap"`
	ListDate             int64    `json:"list_date"`
	SharesOutstanding    string   `json:"shares_outstanding"`
	Aum                  string   `json:"aum"`
	ManagementFeePercent string   `json:"management_fee_percent"`
	LastTradeTime        int64    `json:"last_trade_time"`
	LastQuoteTime        int64    `json:"last_quote_time"`
	VolumeQuantity       float64  `json:"volume_quantity"`
	VolumeUsd            *float64 `json:"volume_usd"`
	Price                *float64 `json:"price,omitempty"`
	PriceChange          *float64 `json:"price_change,omitempty"`
	PriceChangePercent   *float64 `json:"price_change_percent,omitempty"`
	PriceUsd             *float64 `json:"price_usd,omitempty"`
	PriceChangeUsd       *float64 `json:"price_change_usd,omitempty"`
	AssetInfo            struct {
		Nav              *float64 `json:"nav,omitempty"`
		PremiumDiscount  *float64 `json:"premium_discount,omitempty"`
		HoldingQuantity  *float64 `json:"holding_quantity,omitempty"`
		ChangePercent1D  *float64 `json:"change_percent_1d,omitempty"`
		ChangeQuantity1D *float64 `json:"change_quantity_1d,omitempty"`
		ChangePercent7D  *float64 `json:"change_percent_7d,omitempty"`
		ChangeQuantity7D *float64 `json:"change_quantity_7d,omitempty"`
		Date             string   `json:"date"`
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
