package okex

// Copyright (c) 2015-2019 Bitontop Technologies Inc.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

import (
	"time"
)

type ErrorMsg struct {
	Code string `json:"code"`
	Msg  string `json:"message"`
}

type CoinsData []struct {
	CanDeposit    string `json:"can_deposit"`
	CanWithdraw   string `json:"can_withdraw"`
	Currency      string `json:"currency"`
	MinWithdrawal string `json:"min_withdrawal"`
	Name          string `json:"name"`
}

type PairsData []struct {
	BaseCurrency  string `json:"base_currency"`
	InstrumentID  string `json:"instrument_id"`
	MinSize       string `json:"min_size"`
	QuoteCurrency string `json:"quote_currency"`
	SizeIncrement string `json:"size_increment"`
	TickSize      string `json:"tick_size"`
}

type OrderBook struct {
	Asks      [][]string `json:"asks"`
	Bids      [][]string `json:"bids"`
	Timestamp time.Time  `json:"timestamp"`
}

type AccountBalances []struct {
	Frozen    string `json:"frozen"`
	Hold      string `json:"hold"`
	ID        string `json:"id"`
	Currency  string `json:"currency"`
	Balance   string `json:"balance"`
	Available string `json:"available"`
	Holds     string `json:"holds"`
}

type WithdrawResponse struct {
	Result       bool   `json:"result"`
	Amount       string `json:"amount"`
	WithdrawalID string `json:"withdrawal_id"`
	Currency     string `json:"currency"`
	Code         string `json:"code"`
	Message      string `json:"message"`
}

type WithdrawFee []struct {
	Currency string `json:"currency"`
	MaxFee   string `json:"max_fee"`
	MinFee   string `json:"min_fee"`
}

type PlaceOrder struct {
	OrderID   string `json:"order_id"`
	ClientOid string `json:"client_oid"`
	Result    bool   `json:"result"`
	Code      string `json:"code"`
	Message   string `json:"message"`
}

type OrderStatus struct {
	OrderID        string    `json:"order_id"`
	Notional       string    `json:"notional"`
	Price          string    `json:"price"`
	Size           string    `json:"size"`
	InstrumentID   string    `json:"instrument_id"`
	Side           string    `json:"side"`
	Type           string    `json:"type"`
	Timestamp      time.Time `json:"timestamp"`
	FilledSize     string    `json:"filled_size"`
	FilledNotional string    `json:"filled_notional"`
	Status         string    `json:"status"`
	Code           string    `json:"code"`
	Message        string    `json:"message"`
}

type WSOrderBook struct {
	Table  string `json:"table"`
	Action string `json:"action"`
	Data   []struct {
		InstrumentID string     `json:"instrument_id"`
		Asks         [][]string `json:"asks"`
		Bids         [][]string `json:"bids"`
		Timestamp    time.Time  `json:"timestamp"`
		Checksum     int        `json:"checksum"`
	} `json:"data"`
}

// type Transfer struct {
// 	TransferID int     `json:"transfer_id"`
// 	Currency   string  `json:"currency"`
// 	From       int     `json:"from"`
// 	Amount     float64 `json:"amount"`
// 	To         int     `json:"to"`
// 	Result     bool    `json:"result"`
// 	Code       string     `json:"code"`
// 	Message    string  `json:"message"`
// }

type Transfer struct {
	TransferID string `json:"transfer_id"`
	Currency   string `json:"currency"`
	From       string `json:"from"`
	Amount     string `json:"amount"`
	To         string `json:"to"`
	Result     bool   `json:"result"`
}

type AssetBalance []struct {
	ID        string `json:"id"`
	Balance   string `json:"balance"`
	Available string `json:"available"`
	Currency  string `json:"currency"`
	Hold      string `json:"hold"`
}

// type AutoGenerated []struct {
// 	Frozen    string `json:"frozen"`
// 	Hold      string `json:"hold"`
// 	ID        string `json:"id"`
// 	Currency  string `json:"currency"`
// 	Balance   string `json:"balance"`
// 	Available string `json:"available"`
// 	Holds     string `json:"holds"`
// }

type TradeHistory []struct {
	Time      time.Time `json:"time"`
	Timestamp time.Time `json:"timestamp"`
	TradeID   string    `json:"trade_id"`
	Price     string    `json:"price"`
	Size      string    `json:"size"`
	Side      string    `json:"side"`
}

type TickerPrice []struct {
	BestAsk        string    `json:"best_ask"`
	BestBid        string    `json:"best_bid"`
	InstrumentID   string    `json:"instrument_id"`
	ProductID      string    `json:"product_id"`
	Last           string    `json:"last"`
	LastQty        string    `json:"last_qty"`
	Ask            string    `json:"ask"`
	BestAskSize    string    `json:"best_ask_size"`
	Bid            string    `json:"bid"`
	BestBidSize    string    `json:"best_bid_size"`
	Open24H        string    `json:"open_24h"`
	High24H        string    `json:"high_24h"`
	Low24H         string    `json:"low_24h"`
	BaseVolume24H  string    `json:"base_volume_24h"`
	Timestamp      time.Time `json:"timestamp"`
	QuoteVolume24H string    `json:"quote_volume_24h"`
}

type TransferHistory []struct {
	Amount    string    `json:"amount"`
	Balance   string    `json:"balance"`
	Currency  string    `json:"currency"`
	Fee       string    `json:"fee"`
	LedgerID  string    `json:"ledger_id"`
	Timestamp time.Time `json:"timestamp"`
	Typename  string    `json:"typename"`
}

type OpenOrders []struct {
	ClientOid      string    `json:"client_oid"`
	CreatedAt      time.Time `json:"created_at"`
	FilledNotional string    `json:"filled_notional"`
	FilledSize     string    `json:"filled_size"`
	Funds          string    `json:"funds"`
	InstrumentID   string    `json:"instrument_id"`
	Notional       string    `json:"notional"`
	OrderID        string    `json:"order_id"`
	OrderType      string    `json:"order_type"`
	Price          string    `json:"price"`
	PriceAvg       string    `json:"price_avg"`
	ProductID      string    `json:"product_id"`
	Side           string    `json:"side"`
	Size           string    `json:"size"`
	Status         string    `json:"status"`
	State          string    `json:"state"`
	Timestamp      time.Time `json:"timestamp"`
	Type           string    `json:"type"`
}
