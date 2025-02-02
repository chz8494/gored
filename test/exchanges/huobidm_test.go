package test

import (
	"testing"

	//"../exchange/huobidm"
	//"./conf"

	"github.com/chz8494/gored/exchange"
	"github.com/chz8494/gored/pair"
)

// Copyright (c) 2015-2019 Bitontop Technologies Inc.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

/********************Public API********************/

func Test_Huobidm(t *testing.T) {
	e := InitEx(exchange.HUOBIDM)
	//CW FOR THIS WEEK, NW FOR NEXT WEEK, CQ FOR QUARTER
	//BASE IS ALWAYS USD
	pair := pair.GetPairByKey("USD|NWTRX")

	Test_Coins(e)
	Test_Pairs(e)
	Test_Pair(e, pair)
	Test_Orderbook(e, pair)
	//Test_ConstraintFetch(e, pair)
	//Test_Constraint(e, pair)

	//Test_Balance(e, pair)
	//Test_Trading(e, pair, 0.0001, 100)
	//Test_Withdraw(e, pair.Base, 1, "ADDRESS")
}
