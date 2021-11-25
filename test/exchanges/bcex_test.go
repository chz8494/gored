package test

import (
	"testing"

	"github.com/chz8494/gored/exchange"
	"github.com/chz8494/gored/pair"
	// "../exchange/bcex"
	// "./conf"
)

// Copyright (c) 2015-2019 Bitontop Technologies Inc.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

/********************Public API********************/

func Test_Bcex(t *testing.T) {
	e := InitEx(exchange.BCEX)

	pair := pair.GetPairByKey("BTC|ETH")

	Test_Coins(e)
	Test_Pairs(e)
	Test_Pair(e, pair)
	Test_Orderbook(e, pair)
	Test_ConstraintFetch(e, pair)
	Test_Constraint(e, pair)

	// Test_Balance(e, pair)
	// Test_Trading(e, pair, 0.00000001, 100)
	// Test_OrderStatus(e, pair, "2637263726372637263726372637")
	// Test_CancelOrder(e, pair, "2637263726372637263726372637")
	// Test_Withdraw(e, pair.Base, 1, "ADDRESS")
}
