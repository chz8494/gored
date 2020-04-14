package conf

// Copyright (c) 2015-2019 Bitontop Technologies Inc.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

import (
	"github.com/bitontop/gored/exchange"
)

func Exchange(name exchange.ExchangeName, config *exchange.Config) {
	config.ExName = name
	switch name {
	case exchange.BINANCE:
		config.API_KEY = "d9e0KQDxRvbKseG4wrmGKQWsM3dtZl4R5gerOXeNcpXOPYDuiVIlqNngKn5Fa01m"
		config.API_SECRET = "HOhmtTcrEkDXHf6Kla7lDDh0kBawl5crKX3qjxL8oKQWCFbk4jWZg8odxjFYTPes"
		break

	case exchange.BITTREX:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.COINEX:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.STEX:
		config.API_KEY = " "
		config.API_SECRET = "S2S Token"
		config.ExpireTS = 1573081822
		break

	case exchange.BITMEX:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.KUCOIN:
		config.API_KEY = ""
		config.API_SECRET = ""
		config.Passphrase = ""
		break

	case exchange.BITMAX:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.BITSTAMP:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.OTCBTC:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.HUOBI:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.BIBOX:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.OKEX:
		config.API_KEY = ""
		config.API_SECRET = ""
		config.Passphrase = ""
		config.TradePassword = ""
		break

	case exchange.BITZ:
		config.API_KEY = ""
		config.API_SECRET = ""
		config.TradePassword = ""
		break

	case exchange.HITBTC:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.DRAGONEX:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.BIGONE:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.BITFINEX:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.GATEIO:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.IDEX:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.LIQUID:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.BITFOREX:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.TOKOK:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.MXC:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.BITRUE:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.TRADESATOSHI:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.KRAKEN:
		config.API_KEY = ""
		config.API_SECRET = ""
		config.Two_Factor = ""
		break

	case exchange.POLONIEX:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.COINEAL:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.TRADEOGRE:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.COINBENE:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.IBANKDIGITAL:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.LBANK:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.BITMART:
		config.API_KEY = ""
		config.API_SECRET = ""
		config.Passphrase = "" // key name
		break

	case exchange.BIKI:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.BITATM:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.DCOIN:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.GEMINI:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.COINTIGER:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.HUOBIDM:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.BW:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.BITBAY:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.DERIBIT:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.OKEXDM:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.GOKO:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.BCEX:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.DIGIFINEX:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.LATOKEN:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.VIRGOCX:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.ABCC:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.BYBIT:
		config.API_KEY = ""
		config.API_SECRET = ""
		break

	case exchange.ZEBITEX:
		config.API_KEY = ""
		config.API_SECRET = ""

	case exchange.BITHUMB:
		config.API_KEY = ""
		config.API_SECRET = ""

	case exchange.SWITCHEO:
		config.API_KEY = ""
		config.API_SECRET = ""

	case exchange.BLOCKTRADE:
		config.API_KEY = ""
		config.API_SECRET = ""

	case exchange.BKEX:
		config.API_KEY = ""
		config.API_SECRET = ""

	case exchange.NEWCAPITAL:
		config.API_KEY = ""
		config.API_SECRET = ""

	case exchange.COINDEAL:
		config.API_KEY = ""
		config.API_SECRET = ""

	case exchange.HIBITEX:
		config.API_KEY = ""
		config.API_SECRET = ""

	case exchange.BGOGO:
		config.API_KEY = ""
		config.API_SECRET = ""

	case exchange.FTX:
		config.API_KEY = ""
		config.API_SECRET = ""

	case exchange.TXBIT:
		config.API_KEY = ""
		config.API_SECRET = ""

	case exchange.PROBIT:
		config.API_KEY = ""
		config.API_SECRET = ""

	case exchange.BITPIE:
		config.API_KEY = ""
		config.API_SECRET = ""

	case exchange.TAGZ:
		config.API_KEY = ""
		config.API_SECRET = ""

	case exchange.IDCM:
		config.API_KEY = ""
		config.API_SECRET = ""

	case exchange.HOO:
		config.API_KEY = ""
		config.API_SECRET = ""

	case exchange.HOMIEX:
		config.API_KEY = ""
		config.API_SECRET = ""

	case exchange.COINBASE:
		config.API_KEY = "34c04e7f575e5e8a1115af459f38d245"
		config.API_SECRET = "ImgRDl2pfw2W1M4CY2GJ/e3WEr0pI4T/wwrIAII4S4zaYDt6na2jTBrOuOlSSYHH1b1tOK7FG7Ar5xx5H7l/OQ=="
		config.Passphrase = "tanboo."
	}
}
