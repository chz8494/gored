package main

// Copyright (c) 2015-2019 Bitontop Technologies Inc.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

import (
	"log"
	"os"
	"time"

	"github.com/chz8494/gored/coin"
	"github.com/chz8494/gored/exchange"
	"github.com/chz8494/gored/exchange/abcc"
	"github.com/chz8494/gored/exchange/bcex"
	"github.com/chz8494/gored/exchange/bgogo"
	"github.com/chz8494/gored/exchange/bibox"
	"github.com/chz8494/gored/exchange/bigone"
	"github.com/chz8494/gored/exchange/biki"
	"github.com/chz8494/gored/exchange/binance"
	"github.com/chz8494/gored/exchange/bitbay"
	"github.com/chz8494/gored/exchange/bitbns"
	"github.com/chz8494/gored/exchange/bitforex"
	"github.com/chz8494/gored/exchange/bithumb"
	"github.com/chz8494/gored/exchange/bitmart"
	"github.com/chz8494/gored/exchange/bitmax"
	"github.com/chz8494/gored/exchange/bitpie"
	"github.com/chz8494/gored/exchange/bitrue"
	"github.com/chz8494/gored/exchange/bitstamp"
	"github.com/chz8494/gored/exchange/bittrex"
	"github.com/chz8494/gored/exchange/bitz"
	"github.com/chz8494/gored/exchange/bkex"
	"github.com/chz8494/gored/exchange/blocktrade"
	"github.com/chz8494/gored/exchange/bw"
	"github.com/chz8494/gored/exchange/bybit"
	"github.com/chz8494/gored/exchange/coinbase"
	"github.com/chz8494/gored/exchange/coinbene"
	"github.com/chz8494/gored/exchange/coindeal"
	"github.com/chz8494/gored/exchange/coineal"
	"github.com/chz8494/gored/exchange/coinex"
	"github.com/chz8494/gored/exchange/cointiger"
	"github.com/chz8494/gored/exchange/dcoin"
	"github.com/chz8494/gored/exchange/deribit"
	"github.com/chz8494/gored/exchange/digifinex"
	"github.com/chz8494/gored/exchange/dragonex"
	"github.com/chz8494/gored/exchange/ftx"
	"github.com/chz8494/gored/exchange/gateio"
	"github.com/chz8494/gored/exchange/goko"
	"github.com/chz8494/gored/exchange/hibitex"
	"github.com/chz8494/gored/exchange/hitbtc"
	"github.com/chz8494/gored/exchange/homiex"
	"github.com/chz8494/gored/exchange/hoo"
	"github.com/chz8494/gored/exchange/huobi"
	"github.com/chz8494/gored/exchange/huobidm"
	"github.com/chz8494/gored/exchange/ibankdigital"
	"github.com/chz8494/gored/exchange/idcm"
	"github.com/chz8494/gored/exchange/kraken"
	"github.com/chz8494/gored/exchange/kucoin"
	"github.com/chz8494/gored/exchange/latoken"
	"github.com/chz8494/gored/exchange/lbank"
	"github.com/chz8494/gored/exchange/liquid"
	"github.com/chz8494/gored/exchange/mxc"
	"github.com/chz8494/gored/exchange/newcapital"
	"github.com/chz8494/gored/exchange/okex"
	"github.com/chz8494/gored/exchange/okexdm"
	"github.com/chz8494/gored/exchange/otcbtc"
	"github.com/chz8494/gored/exchange/poloniex"
	"github.com/chz8494/gored/exchange/probit"
	"github.com/chz8494/gored/exchange/stex"
	"github.com/chz8494/gored/exchange/switcheo"
	"github.com/chz8494/gored/exchange/tagz"
	"github.com/chz8494/gored/exchange/tokok"
	"github.com/chz8494/gored/exchange/tradeogre"
	"github.com/chz8494/gored/exchange/tradesatoshi"
	"github.com/chz8494/gored/exchange/txbit"
	"github.com/chz8494/gored/exchange/virgocx"
	"github.com/chz8494/gored/exchange/zebitex"
	"github.com/chz8494/gored/pair"
	"github.com/chz8494/gored/test/conf"
	"github.com/chz8494/gored/utils"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	exMan := exchange.CreateExchangeManager()

	if len(os.Args) > 1 {
		switch os.Args[1] {
		/* case "export":
		Init(exchange.EXCHANGE_API, "")
		utils.ConvertBaseDataToJson("./data")
		for _, ex := range exMan.GetExchanges() {
			utils.ConvertExchangeDataToJson("./data", ex)
		}
		break */
		case "json":
			Init(exchange.JSON_FILE, "./data")
			for _, ex := range exMan.GetExchanges() {
				for _, coin := range ex.GetCoins() {
					log.Printf("%s Coin %+v", ex.GetName(), coin)
				}
				for _, pair := range ex.GetPairs() {
					log.Printf("%s Pair %+v", ex.GetName(), pair)
				}
			}
			break
		case "renew":
			Init(exchange.JSON_FILE, "./data")
			updateConfig := &exchange.Update{
				ExNames: exMan.GetSupportExchanges(),
				Method:  exchange.TIME_TIGGER,
				Time:    10 * time.Second,
			}
			exMan.UpdateExData(updateConfig)
			break
		case "test":
			base := coin.Coin{
				Code: "BTC",
			}
			target := coin.Coin{
				Code: "ETH",
			}
			pair := pair.Pair{
				Base:   &base,
				Target: &target,
			}
			log.Println(pair)

			// okex.Socket(&pair)
			// stex.Socket()
			// bitfinex.Socket()
		}
	}
}

func Init(source exchange.DataSource, sourceURI string) {
	coin.Init()
	pair.Init()
	if source == exchange.JSON_FILE {
		utils.GetCommonDataFromJSON(sourceURI)
	}
	config := &exchange.Config{}
	config.Source = source
	config.SourceURI = sourceURI

	InitBinance(config)
	InitBittrex(config)
	InitCoinex(config)
	InitStex(config)
	InitKucoin(config)
	InitBitmax(config)
	InitBitstamp(config)
	InitOtcbtc(config)
	InitHuobi(config)
	InitBibox(config)
	InitOkex(config)
	InitBitz(config)
	InitHitbtc(config)
	InitDragonex(config)
	InitBigone(config)
	InitGateio(config)
	InitLiquid(config)
	InitBitforex(config)
	InitTokok(config)
	InitMxc(config)
	InitBitrue(config)
	InitTradeSatoshi(config)
	InitKraken(config)
	InitPoloniex(config)
	InitCoineal(config)
	InitTradeogre(config)
	InitCoinbene(config)
	InitIbankdigital(config)
	InitLbank(config)
	InitBitmart(config)
	InitDcoin(config)
	InitBiki(config)
	InitCointiger(config)
	InitHuobidm(config)
	InitBw(config)
	InitBitbay(config)
	InitDeribit(config)
	InitOkexdm(config)
	InitGoko(config)
	InitBcex(config)
	InitDigifinex(config)
	InitLatoken(config)
	InitVirgocx(config)
	InitAbcc(config)
	// InitBybit(config)
	InitZebitex(config)
	InitBithumb(config)
	InitSwitcheo(config)
	InitBlocktrade(config)
	InitBkex(config)
	InitNewcapital(config)
	InitCoindeal(config)
	// InitHibitex(config)
	InitBgogo(config)
	InitFtx(config)
	InitTxbit(config)
	InitProbit(config)
	InitBitpie(config)
	InitTagz(config)
	InitIdcm(config)
	InitHoo(config)
	InitHomiex(config)
	InitCoinbase(config)
	// InitBitbns(config)
}

func InitBinance(config *exchange.Config) {
	conf.Exchange(exchange.BINANCE, config)
	ex := binance.CreateBinance(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitBittrex(config *exchange.Config) {
	conf.Exchange(exchange.BITTREX, config)
	ex := bittrex.CreateBittrex(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitCoinex(config *exchange.Config) {
	conf.Exchange(exchange.COINEX, config)
	ex := coinex.CreateCoinex(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitStex(config *exchange.Config) {
	conf.Exchange(exchange.STEX, config)
	ex := stex.CreateStex(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitKucoin(config *exchange.Config) {
	conf.Exchange(exchange.KUCOIN, config)
	ex := kucoin.CreateKucoin(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitBitmax(config *exchange.Config) {
	conf.Exchange(exchange.BITMAX, config)
	ex := bitmax.CreateBitmax(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitBitstamp(config *exchange.Config) {
	conf.Exchange(exchange.BITSTAMP, config)
	ex := bitstamp.CreateBitstamp(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitOtcbtc(config *exchange.Config) {
	conf.Exchange(exchange.OTCBTC, config)
	ex := otcbtc.CreateOtcbtc(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitHuobi(config *exchange.Config) {
	conf.Exchange(exchange.HUOBI, config)
	ex := huobi.CreateHuobi(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitBibox(config *exchange.Config) {
	conf.Exchange(exchange.BIBOX, config)
	ex := bibox.CreateBibox(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitOkex(config *exchange.Config) {
	conf.Exchange(exchange.OKEX, config)
	ex := okex.CreateOkex(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitBitz(config *exchange.Config) {
	conf.Exchange(exchange.BITZ, config)
	ex := bitz.CreateBitz(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitHitbtc(config *exchange.Config) {
	conf.Exchange(exchange.HITBTC, config)
	ex := hitbtc.CreateHitbtc(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitDragonex(config *exchange.Config) {
	conf.Exchange(exchange.DRAGONEX, config)
	ex := dragonex.CreateDragonex(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitBigone(config *exchange.Config) {
	conf.Exchange(exchange.BIGONE, config)
	ex := bigone.CreateBigone(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitGateio(config *exchange.Config) {
	conf.Exchange(exchange.GATEIO, config)
	ex := gateio.CreateGateio(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitLiquid(config *exchange.Config) {
	conf.Exchange(exchange.LIQUID, config)
	ex := liquid.CreateLiquid(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitBitforex(config *exchange.Config) {
	conf.Exchange(exchange.BITFOREX, config)
	ex := bitforex.CreateBitforex(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitTokok(config *exchange.Config) {
	conf.Exchange(exchange.TOKOK, config)
	ex := tokok.CreateTokok(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitMxc(config *exchange.Config) {
	conf.Exchange(exchange.MXC, config)
	ex := mxc.CreateMxc(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitBitrue(config *exchange.Config) {
	conf.Exchange(exchange.BITRUE, config)
	ex := bitrue.CreateBitrue(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitTradeSatoshi(config *exchange.Config) {
	conf.Exchange(exchange.TRADESATOSHI, config)
	ex := tradesatoshi.CreateTradeSatoshi(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitKraken(config *exchange.Config) {
	conf.Exchange(exchange.KRAKEN, config)
	ex := kraken.CreateKraken(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitPoloniex(config *exchange.Config) {
	conf.Exchange(exchange.POLONIEX, config)
	ex := poloniex.CreatePoloniex(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitCoineal(config *exchange.Config) {
	conf.Exchange(exchange.COINEAL, config)
	ex := coineal.CreateCoineal(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitTradeogre(config *exchange.Config) {
	conf.Exchange(exchange.TRADEOGRE, config)
	ex := tradeogre.CreateTradeogre(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitCoinbene(config *exchange.Config) {
	conf.Exchange(exchange.COINBENE, config)
	ex := coinbene.CreateCoinbene(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitIbankdigital(config *exchange.Config) {
	conf.Exchange(exchange.IBANKDIGITAL, config)
	ex := ibankdigital.CreateIbankdigital(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitLbank(config *exchange.Config) {
	conf.Exchange(exchange.LBANK, config)
	ex := lbank.CreateLbank(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitBitmart(config *exchange.Config) {
	conf.Exchange(exchange.BITMART, config)
	ex := bitmart.CreateBitmart(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitBiki(config *exchange.Config) {
	conf.Exchange(exchange.BIKI, config)
	ex := biki.CreateBiki(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitDcoin(config *exchange.Config) {
	conf.Exchange(exchange.DCOIN, config)
	ex := dcoin.CreateDcoin(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitCointiger(config *exchange.Config) {
	conf.Exchange(exchange.COINTIGER, config)
	ex := cointiger.CreateCointiger(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitHuobidm(config *exchange.Config) {
	conf.Exchange(exchange.HUOBIDM, config)
	ex := huobidm.CreateHuobidm(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitBw(config *exchange.Config) {
	conf.Exchange(exchange.BW, config)
	ex := bw.CreateBw(config)

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitBitbay(config *exchange.Config) {
	conf.Exchange(exchange.BITBAY, config)
	ex := bitbay.CreateBitbay(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitDeribit(config *exchange.Config) {
	conf.Exchange(exchange.DERIBIT, config)
	ex := deribit.CreateDeribit(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitOkexdm(config *exchange.Config) {
	conf.Exchange(exchange.OKEXDM, config)
	ex := okexdm.CreateOkexdm(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitGoko(config *exchange.Config) {
	conf.Exchange(exchange.GOKO, config)
	ex := goko.CreateGoko(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitBcex(config *exchange.Config) {
	conf.Exchange(exchange.BCEX, config)
	ex := bcex.CreateBcex(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitDigifinex(config *exchange.Config) {
	conf.Exchange(exchange.DIGIFINEX, config)
	ex := digifinex.CreateDigifinex(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitLatoken(config *exchange.Config) {
	conf.Exchange(exchange.LATOKEN, config)
	ex := latoken.CreateLatoken(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitVirgocx(config *exchange.Config) {
	conf.Exchange(exchange.VIRGOCX, config)
	ex := virgocx.CreateVirgocx(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitAbcc(config *exchange.Config) {
	conf.Exchange(exchange.ABCC, config)
	ex := abcc.CreateAbcc(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitBybit(config *exchange.Config) {
	conf.Exchange(exchange.BYBIT, config)
	ex := bybit.CreateBybit(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitZebitex(config *exchange.Config) {
	conf.Exchange(exchange.ZEBITEX, config)
	ex := zebitex.CreateZebitex(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitBithumb(config *exchange.Config) {
	conf.Exchange(exchange.BITHUMB, config)
	ex := bithumb.CreateBithumb(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitSwitcheo(config *exchange.Config) {
	conf.Exchange(exchange.SWITCHEO, config)
	ex := switcheo.CreateSwitcheo(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitBlocktrade(config *exchange.Config) {
	conf.Exchange(exchange.BLOCKTRADE, config)
	ex := blocktrade.CreateBlocktrade(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitBkex(config *exchange.Config) {
	conf.Exchange(exchange.BKEX, config)
	ex := bkex.CreateBkex(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitNewcapital(config *exchange.Config) {
	conf.Exchange(exchange.NEWCAPITAL, config)
	ex := newcapital.CreateNewcapital(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitCoindeal(config *exchange.Config) {
	conf.Exchange(exchange.COINDEAL, config)
	ex := coindeal.CreateCoindeal(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitHibitex(config *exchange.Config) {
	conf.Exchange(exchange.HIBITEX, config)
	ex := hibitex.CreateHibitex(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitBgogo(config *exchange.Config) {
	conf.Exchange(exchange.BGOGO, config)
	ex := bgogo.CreateBgogo(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitFtx(config *exchange.Config) {
	conf.Exchange(exchange.FTX, config)
	ex := ftx.CreateFtx(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitTxbit(config *exchange.Config) {
	conf.Exchange(exchange.TXBIT, config)
	ex := txbit.CreateTxbit(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitProbit(config *exchange.Config) {
	conf.Exchange(exchange.PROBIT, config)
	ex := probit.CreateProbit(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitBitpie(config *exchange.Config) {
	conf.Exchange(exchange.BITPIE, config)
	ex := bitpie.CreateBitpie(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitTagz(config *exchange.Config) {
	conf.Exchange(exchange.TAGZ, config)
	ex := tagz.CreateTagz(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitIdcm(config *exchange.Config) {
	conf.Exchange(exchange.IDCM, config)
	ex := idcm.CreateIdcm(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitHoo(config *exchange.Config) {
	conf.Exchange(exchange.HOO, config)
	ex := hoo.CreateHoo(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitHomiex(config *exchange.Config) {
	conf.Exchange(exchange.HOMIEX, config)
	ex := homiex.CreateHomiex(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitCoinbase(config *exchange.Config) {
	conf.Exchange(exchange.COINBASE, config)
	ex := coinbase.CreateCoinbase(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}

func InitBitbns(config *exchange.Config) {
	conf.Exchange(exchange.BITBNS, config)
	ex := bitbns.CreateBitbns(config)
	log.Printf("Initial [ %12v ] ", ex.GetName())

	exMan := exchange.CreateExchangeManager()
	exMan.Add(ex)
}
