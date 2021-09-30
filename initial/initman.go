package initial

import (
	"sync"

	"github.com/chz8494/gored/exchange"
	"github.com/chz8494/gored/exchange/abcc"
	"github.com/chz8494/gored/exchange/bcex"
	"github.com/chz8494/gored/exchange/bgogo"
	"github.com/chz8494/gored/exchange/bibox"
	"github.com/chz8494/gored/exchange/bigone"
	"github.com/chz8494/gored/exchange/biki"
	"github.com/chz8494/gored/exchange/binance"
	"github.com/chz8494/gored/exchange/binancedex"
	"github.com/chz8494/gored/exchange/bitbay"
	"github.com/chz8494/gored/exchange/bitbns"
	"github.com/chz8494/gored/exchange/bitfinex"
	"github.com/chz8494/gored/exchange/bitforex"
	"github.com/chz8494/gored/exchange/bithumb"
	"github.com/chz8494/gored/exchange/bitmart"
	"github.com/chz8494/gored/exchange/bitmax"
	"github.com/chz8494/gored/exchange/bitmex"
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
	"github.com/chz8494/gored/exchange/huobiotc"
	"github.com/chz8494/gored/exchange/ibankdigital"
	"github.com/chz8494/gored/exchange/idcm"
	"github.com/chz8494/gored/exchange/idex"
	"github.com/chz8494/gored/exchange/kraken"
	"github.com/chz8494/gored/exchange/kucoin"
	"github.com/chz8494/gored/exchange/latoken"
	"github.com/chz8494/gored/exchange/lbank"
	"github.com/chz8494/gored/exchange/liquid"
	"github.com/chz8494/gored/exchange/mxc"
	"github.com/chz8494/gored/exchange/newcapital"
	"github.com/chz8494/gored/exchange/okex"
	"github.com/chz8494/gored/exchange/okexdm"
	"github.com/chz8494/gored/exchange/oksim"
	"github.com/chz8494/gored/exchange/otcbtc"
	"github.com/chz8494/gored/exchange/poloniex"
	"github.com/chz8494/gored/exchange/probit"
	"github.com/chz8494/gored/exchange/stex"
	"github.com/chz8494/gored/exchange/switcheo"
	"github.com/chz8494/gored/exchange/tagz"
	"github.com/chz8494/gored/exchange/tokok"
	"github.com/chz8494/gored/exchange/tradeogre"
	"github.com/chz8494/gored/exchange/txbit"
	"github.com/chz8494/gored/exchange/virgocx"
	"github.com/chz8494/gored/exchange/zebitex"
)

var instance *InitManager
var once sync.Once

type InitManager struct {
	exMan *exchange.ExchangeManager
}

func CreateInitManager() *InitManager {
	once.Do(func() {
		instance = &InitManager{
			exMan: exchange.CreateExchangeManager(),
		}
	})
	return instance
}

func (e *InitManager) Init(config *exchange.Config) exchange.Exchange {
	switch config.ExName {
	case exchange.BINANCE:
		ex := binance.CreateBinance(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITTREX:
		ex := bittrex.CreateBittrex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.COINEX:
		ex := coinex.CreateCoinex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.STEX:
		ex := stex.CreateStex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITMEX:
		ex := bitmex.CreateBitmex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.KUCOIN:
		ex := kucoin.CreateKucoin(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.HUOBIOTC:
		ex := huobiotc.CreateHuobiOTC(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITMAX:
		ex := bitmax.CreateBitmax(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITSTAMP:
		ex := bitstamp.CreateBitstamp(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.OTCBTC:
		ex := otcbtc.CreateOtcbtc(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.HUOBI:
		ex := huobi.CreateHuobi(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BIBOX:
		ex := bibox.CreateBibox(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.OKEX:
		ex := okex.CreateOkex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITZ:
		ex := bitz.CreateBitz(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.HITBTC:
		ex := hitbtc.CreateHitbtc(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.DRAGONEX:
		ex := dragonex.CreateDragonex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BIGONE:
		ex := bigone.CreateBigone(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITFINEX:
		ex := bitfinex.CreateBitfinex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.GATEIO:
		ex := gateio.CreateGateio(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.IDEX:
		ex := idex.CreateIdex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.LIQUID:
		ex := liquid.CreateLiquid(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITFOREX:
		ex := bitforex.CreateBitforex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.TOKOK:
		ex := tokok.CreateTokok(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.MXC:
		ex := mxc.CreateMxc(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITRUE:
		ex := bitrue.CreateBitrue(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	// case exchange.TRADESATOSHI:
	// 	ex := tradesatoshi.CreateTradeSatoshi(config)
	// 	if ex != nil {
	// 		e.exMan.Add(ex)
	// 	}
	// 	return ex

	case exchange.KRAKEN:
		ex := kraken.CreateKraken(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.POLONIEX:
		ex := poloniex.CreatePoloniex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.COINEAL:
		ex := coineal.CreateCoineal(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.TRADEOGRE:
		ex := tradeogre.CreateTradeogre(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.COINBENE:
		ex := coinbene.CreateCoinbene(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.IBANKDIGITAL:
		ex := ibankdigital.CreateIbankdigital(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.LBANK:
		ex := lbank.CreateLbank(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BINANCEDEX:
		ex := binancedex.CreateBinanceDex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITMART:
		ex := bitmart.CreateBitmart(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BIKI:
		ex := biki.CreateBiki(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.DCOIN:
		ex := dcoin.CreateDcoin(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.COINTIGER:
		ex := cointiger.CreateCointiger(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.HUOBIDM:
		ex := huobidm.CreateHuobidm(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BW:
		ex := bw.CreateBw(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITBAY:
		ex := bitbay.CreateBitbay(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.DERIBIT:
		ex := deribit.CreateDeribit(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.OKEXDM:
		ex := okexdm.CreateOkexdm(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.GOKO:
		ex := goko.CreateGoko(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BCEX:
		ex := bcex.CreateBcex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.DIGIFINEX:
		ex := digifinex.CreateDigifinex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.LATOKEN:
		ex := latoken.CreateLatoken(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.VIRGOCX:
		ex := virgocx.CreateVirgocx(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.ABCC:
		ex := abcc.CreateAbcc(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BYBIT:
		ex := bybit.CreateBybit(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.ZEBITEX:
		ex := zebitex.CreateZebitex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITHUMB:
		ex := bithumb.CreateBithumb(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.SWITCHEO:
		ex := switcheo.CreateSwitcheo(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BLOCKTRADE:
		ex := blocktrade.CreateBlocktrade(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BKEX:
		ex := bkex.CreateBkex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.NEWCAPITAL:
		ex := newcapital.CreateNewcapital(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.COINDEAL:
		ex := coindeal.CreateCoindeal(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.HIBITEX:
		ex := hibitex.CreateHibitex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BGOGO:
		ex := bgogo.CreateBgogo(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.FTX:
		ex := ftx.CreateFtx(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.TXBIT:
		ex := txbit.CreateTxbit(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.PROBIT:
		ex := probit.CreateProbit(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITPIE:
		ex := bitpie.CreateBitpie(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.TAGZ:
		ex := tagz.CreateTagz(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.IDCM:
		ex := idcm.CreateIdcm(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.HOO:
		ex := hoo.CreateHoo(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.HOMIEX:
		ex := homiex.CreateHomiex(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.COINBASE:
		ex := coinbase.CreateCoinbase(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.BITBNS:
		ex := bitbns.CreateBitbns(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	case exchange.OKSIM:
		ex := oksim.CreateOksim(config)
		if ex != nil {
			e.exMan.Add(ex)
		}
		return ex

	}
	return nil
}
