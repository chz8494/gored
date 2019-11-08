package test

import (
	"log"
	"testing"

	"github.com/bitontop/gored/coin"
	"github.com/bitontop/gored/exchange"
	"github.com/bitontop/gored/pair"

	"github.com/bitontop/gored/exchange/bitz"
	"github.com/bitontop/gored/test/conf"
	// "../../exchange/bitz"
	// "../conf"
)

// Copyright (c) 2015-2019 Bitontop Technologies Inc.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

/********************Public API********************/

func Test_Bitz(t *testing.T) {
	e := InitBitz()

	pair := pair.GetPairByKey("BTC|ETH")

	// Test_Coins(e)
	// Test_Pairs(e)
	// Test_Pair(e, pair)
	// Test_Orderbook(e, pair)
	// Test_ConstraintFetch(e, pair)
	// Test_Constraint(e, pair)

	Test_Balance(e, pair)
	// Test_Trading(e, pair, 0.00000001, 100)
	// Test_Withdraw(e, pair.Base, 1, "ADDRESS")

	// Test Withdraw
	opWithdraw := &exchange.AccountOperation{
		Type:            exchange.Withdraw,
		Coin:            pair.Target,
		WithdrawAmount:  "1",
		WithdrawAddress: "addr",
		DebugMode:       true,
	}
	err := e.DoAccoutOperation(opWithdraw)
	if err != nil {
		log.Printf("%v", err)
	}
	log.Printf("WithdrawID: %v, err: %v", opWithdraw.WithdrawID, opWithdraw.Error)
}

func InitBitz() exchange.Exchange {
	coin.Init()
	pair.Init()
	config := &exchange.Config{}
	config.Source = exchange.EXCHANGE_API
	conf.Exchange(exchange.BITZ, config)

	ex := bitz.CreateBitz(config)
	log.Printf("Initial [ %v ] ", ex.GetName())

	config = nil
	return ex
}
