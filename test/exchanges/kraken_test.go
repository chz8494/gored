package test

import (
	"testing"

	"github.com/chz8494/gored/exchange"
	"github.com/chz8494/gored/pair"
	// "../../exchange/kraken"
	// "../../utils"
	// "../conf"
)

// Copyright (c) 2015-2019 Bitontop Technologies Inc.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

/********************Public API********************/

func Test_Kraken(t *testing.T) {
	e := InitEx(exchange.KRAKEN)
	pair := pair.GetPairByKey("USDT|ETH")

	// Test_Coins(e)
	// Test_Pairs(e)
	Test_Pair(e, pair)
	// Test_Orderbook(e, pair)
	// Test_ConstraintFetch(e, pair)
	// Test_Constraint(e, pair)

	// Test_Balance(e, pair)
	// Test_Trading(e, pair, 0.0001, 100)
	// Test_Trading_Sell(e, pair, 0.05, 0.04)
	// Test_Withdraw(e, pair.Base, 1, "ADDRESS")

	// Test_DoWithdraw(e, pair.Target, "1", "0x37E0Fc27C6cDB5035B2a3d0682B4E7C05A4e6C46", "tag")

	Test_TradeHistory(e, pair)
	Test_NewOrderBook(e, pair)
}
