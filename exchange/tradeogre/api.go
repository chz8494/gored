package tradeogre

// Copyright (c) 2015-2019 Bitontop Technologies Inc.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/chz8494/gored/coin"
	"github.com/chz8494/gored/exchange"
	"github.com/chz8494/gored/pair"
)

const (
	API_URL string = "https://tradeogre.com/api/v1"
)

/*API Base Knowledge
Path: API function. Usually after the base endpoint URL
Method:
	Get - Call a URL, API return a response
	Post - Call a URL & send a request, API return a response
Public API:
	It doesn't need authorization/signature , can be called by browser to get response.
	using exchange.HttpGetRequest/exchange.HttpPostRequest
Private API:
	Authorization/Signature is requried. The signature request should look at Exchange API Document.
	using ApiKeyGet/ApiKeyPost
Response:
	Response is a json structure.
	Copy the json to https://transform.now.sh/json-to-go/ convert to go Struct.
	Add the go Struct to model.go

ex. Get /api/v1/depth
Get - Method
/api/v1/depth - Path*/

/*************** Public API ***************/
/*Get Coins Information (If API provide)
Step 1: Change Instance Name    (e *<exchange Instance Name>)
Step 2: Add Model of API Response
Step 3: Modify API Path(strRequestUrl)*/
func (e *Tradeogre) GetCoinsData() error {
	pairsData := PairsData{}

	strRequestUrl := "/markets"
	strUrl := API_URL + strRequestUrl

	jsonCurrencyReturn := exchange.HttpGetRequest(strUrl, nil)
	if err := json.Unmarshal([]byte(jsonCurrencyReturn), &pairsData); err != nil {
		return fmt.Errorf("%s Get Coins Json Unmarshal Err: %v %v", e.GetName(), err, jsonCurrencyReturn)
	}

	for _, pairs := range pairsData {
		for key, _ := range pairs {
			coinStrs := strings.Split(key, "-")
			base := &coin.Coin{}
			target := &coin.Coin{}
			switch e.Source {
			case exchange.EXCHANGE_API:
				base = coin.GetCoin(coinStrs[0])
				if base == nil {
					base = &coin.Coin{}
					base.Code = coinStrs[0]
					coin.AddCoin(base)
				}
				target = coin.GetCoin(coinStrs[1])
				if target == nil {
					target = &coin.Coin{}
					target.Code = coinStrs[1]
					coin.AddCoin(target)
				}
			case exchange.JSON_FILE:
				base = e.GetCoinBySymbol(coinStrs[0])
				target = e.GetCoinBySymbol(coinStrs[1])
			}

			if base != nil {
				coinConstraint := e.GetCoinConstraint(base)
				if coinConstraint == nil {
					coinConstraint = &exchange.CoinConstraint{
						CoinID:       base.ID,
						Coin:         base,
						ExSymbol:     coinStrs[0],
						ChainType:    exchange.MAINNET,
						TxFee:        DEFAULT_TXFEE,
						Withdraw:     DEFAULT_WITHDRAW,
						Deposit:      DEFAULT_DEPOSIT,
						Confirmation: DEFAULT_CONFIRMATION,
						Listed:       DEFAULT_LISTED,
					}
				} else {
					coinConstraint.ExSymbol = coinStrs[0]
				}
				e.SetCoinConstraint(coinConstraint)
			}

			if target != nil {
				coinConstraint := e.GetCoinConstraint(target)
				if coinConstraint == nil {
					coinConstraint = &exchange.CoinConstraint{
						CoinID:       target.ID,
						Coin:         target,
						ExSymbol:     coinStrs[1],
						ChainType:    exchange.MAINNET,
						TxFee:        DEFAULT_TXFEE,
						Withdraw:     DEFAULT_WITHDRAW,
						Deposit:      DEFAULT_DEPOSIT,
						Confirmation: DEFAULT_CONFIRMATION,
						Listed:       DEFAULT_LISTED,
					}
				} else {
					coinConstraint.ExSymbol = coinStrs[1]
				}
				e.SetCoinConstraint(coinConstraint)
			}
		}
	}
	return nil
}

/* GetPairsData - Get Pairs Information (If API provide)
Step 1: Change Instance Name    (e *<exchange Instance Name>)
Step 2: Add Model of API Response
Step 3: Modify API Path(strRequestUrl)*/
func (e *Tradeogre) GetPairsData() error {
	pairsData := PairsData{}

	strRequestUrl := "/markets"
	strUrl := API_URL + strRequestUrl

	jsonSymbolsReturn := exchange.HttpGetRequest(strUrl, nil)
	if err := json.Unmarshal([]byte(jsonSymbolsReturn), &pairsData); err != nil {
		return fmt.Errorf("%s Get Pairs Json Unmarshal Err: %v %v", e.GetName(), err, jsonSymbolsReturn)
	}
	for _, pairs := range pairsData {
		for key, _ := range pairs {
			coinStrs := strings.Split(key, "-")
			p := &pair.Pair{}
			switch e.Source {
			case exchange.EXCHANGE_API:
				base := coin.GetCoin(coinStrs[0])
				target := coin.GetCoin(coinStrs[1])
				if base != nil && target != nil {
					p = pair.GetPair(base, target)
				}
			case exchange.JSON_FILE:
				p = e.GetPairBySymbol(key)
			}
			if p != nil {
				pairConstraint := e.GetPairConstraint(p)
				if pairConstraint == nil {
					pairConstraint = &exchange.PairConstraint{
						PairID:      p.ID,
						Pair:        p,
						ExSymbol:    key,
						MakerFee:    DEFAULT_MAKER_FEE,
						TakerFee:    DEFAULT_TAKER_FEE,
						LotSize:     DEFAULT_LOT_SIZE,
						PriceFilter: DEFAULT_PRICE_FILTER,
						Listed:      DEFAULT_LISTED,
					}
				} else {
					pairConstraint.ExSymbol = key
				}
				e.SetPairConstraint(pairConstraint)
			}
		}
	}
	return nil
}

/*Get Pair Market Depth
Step 1: Change Instance Name    (e *<exchange Instance Name>)
Step 2: Add Model of API Response
Step 3: Get Exchange Pair Code ex. symbol := e.GetPairCode(p)
Step 4: Modify API Path(strRequestUrl)
Step 5: Add Params - Depend on API request
Step 6: Convert the response to Standard Maker struct*/
/*
func (e *Tradeogre) GetPrice(pair *pair.Pair) (string, error) {
	pairsDetail := PairsDetail{}

	strRequestUrl := "/ticker/"
	strUrl := API_URL + strRequestUrl + pair.Name

	jsonSymbolsReturn := exchange.HttpGetRequest(strUrl, nil)
	if err := json.Unmarshal([]byte(jsonSymbolsReturn), &pairsDetail); err != nil {
		return "error",fmt.Errorf("%s Get Coin detail Json Unmarshal Err: %v %v", e.GetName(), err, jsonSymbolsReturn)
	}
  fmt.Println(pairsDetail.Price)
  return pairsDetail.Price,nil
}
*/

func (e *Tradeogre) OrderBook(pair *pair.Pair) (*exchange.Maker, error) {

	orderBook := OrderBook{}
	symbol := e.GetSymbolByPair(pair)

	strRequestUrl := "/orders/" + symbol
	strUrl := API_URL + strRequestUrl

	maker := &exchange.Maker{
		WorkerIP:        exchange.GetExternalIP(),
		Source:          exchange.EXCHANGE_API,
		BeforeTimestamp: float64(time.Now().UnixNano() / 1e6),
	}

	jsonOrderbook := exchange.HttpGetRequest(strUrl, nil)
	jsonOrderbook = strings.Replace(jsonOrderbook, `"buy":[]`, `"buy":{}`, -1)
	jsonOrderbook = strings.Replace(jsonOrderbook, `"sell":[]`, `"sell":{}`, -1)
	if err := json.Unmarshal([]byte(jsonOrderbook), &orderBook); err != nil {
		return nil, fmt.Errorf("%s Get Orderbook Json Unmarshal Err: %v %v", e.GetName(), err, jsonOrderbook)
	} else if orderBook.Success != "true" {
		return nil, fmt.Errorf("%s Get Orderbook Failed: %v", e.GetName(), jsonOrderbook)
	}

	maker.AfterTimestamp = float64(time.Now().UnixNano() / 1e6)

	var err error
	var buyRates []string
	for rate, _ := range orderBook.Buy {
		buyRates = append(buyRates, rate)
	}
	sort.Strings(buyRates)
	for i := len(buyRates) - 1; i >= 0; i-- {
		var buydata exchange.Order

		buydata.Rate, err = strconv.ParseFloat(buyRates[i], 64)
		if err != nil {
			return nil, err
		}
		buydata.Quantity, err = strconv.ParseFloat(orderBook.Buy[buyRates[i]], 64)
		if err != nil {
			return nil, err
		}

		maker.Bids = append(maker.Bids, buydata)
	}

	var sellRates []string
	for rate, _ := range orderBook.Sell {
		sellRates = append(sellRates, rate)
	}
	sort.Strings(sellRates)
	for _, rate := range sellRates {
		var selldata exchange.Order

		selldata.Rate, err = strconv.ParseFloat(rate, 64)
		if err != nil {
			return nil, err
		}
		selldata.Quantity, err = strconv.ParseFloat(orderBook.Sell[rate], 64)
		if err != nil {
			return nil, err
		}

		maker.Asks = append(maker.Asks, selldata)
	}

	return maker, nil
}

/*
func (e *Tradeogre) LoadPublicData(operation *exchange.PublicOperation) error {
	return fmt.Errorf("LoadPublicData :: Operation type invalid: %+v", operation.Type)
}
*/
/*************** Private API ***************/
func (e *Tradeogre) DoAccountOperation(operation *exchange.AccountOperation) error {
	switch operation.Type {
	case exchange.GetOrderHistory:
		switch operation.Wallet {
		case exchange.SpotWallet:
			return e.doAccountHistory(operation)
		}
	case exchange.Balance:
		switch operation.Wallet {
		case exchange.SpotWallet:
			return e.doAccountBalance(operation)
		}
	case exchange.CancelOrderID:
		switch operation.Wallet {
		case exchange.SpotWallet:
			return e.CancelOrderID(operation)
		}
	//default: fmt.Println("yes")
	}

	return fmt.Errorf("LoadPublicData :: Operation type invalid: %+v", operation.Type)
}

func (e *Tradeogre) doAccountHistory(operation *exchange.AccountOperation) error {
	if e.API_KEY == "" || e.API_SECRET == "" {
		return fmt.Errorf("%s API Key or Secret Key are nil", e.GetName())
	}
	tickerPrice := PairsData{}
	strRequest := "/account/market/history"
	mapParams := make(map[string]string)
	mapParams["market"] = "BTC-WOW"
	jsonOrderHistory := e.ApiKeyRequest("POST", strRequest, mapParams)
	fmt.Println(jsonOrderHistory)
	if err := json.Unmarshal([]byte(jsonOrderHistory), &tickerPrice); err != nil {
		return fmt.Errorf("%s doTickerPrice Result Unmarshal Err: %v %s", e.GetName(), err, jsonOrderHistory)
	}
	return nil
}

func (e *Tradeogre) doAccountBalance(operation *exchange.AccountOperation) error {
	if e.API_KEY == "" || e.API_SECRET == "" {
		return fmt.Errorf("%s API Key or Secret Key are nil", e.GetName())
	}
	accountBalance := AccountBalances{}
	strRequest := "/account/balances"

	jsonBalanceReturn := e.ApiKeyRequest("GET", strRequest, make(map[string]string))
	if err := json.Unmarshal([]byte(jsonBalanceReturn), &accountBalance); err != nil {
		return fmt.Errorf("%s UpdateAllBalances Json Unmarshal Err: %v %v", e.GetName(), err, jsonBalanceReturn)
	} else if !accountBalance.Success {
		return fmt.Errorf("%s UpdateAllBalances Failed: %v", e.GetName(), jsonBalanceReturn)
	}
	for key, data := range accountBalance.Balances {
		c := e.GetCoinBySymbol(key)
		//fmt.Println(c)
		if c != nil {
			freeBalance, err := strconv.ParseFloat(data, 64)
			if err != nil {
				return fmt.Errorf("%s balance parse error: %v, %v", e.GetName(), err, data)
			}
			balanceMap.Set(c.Code, freeBalance)
		}
	}
	if operation.DebugMode {
		operation.RequestURI = API_URL
		operation.CallResponce = jsonBalanceReturn
	}
	operation.Balances = accountBalance.Balances
	//fmt.Println(accountBalance.Balances)
	return nil
}

func (e *Tradeogre) CancelOrderID(operation *exchange.AccountOperation) error {
	if e.API_KEY == "" || e.API_SECRET == "" {
		return fmt.Errorf("%s API Key or Secret Key are nil", e.GetName())
	}

	cancelOrder := CancelOrder{}
	strRequest := "/order/cancel"
	mapParams := make(map[string]string)
	mapParams["uuid"] = operation.CancelOrderID

	jsonCancelOrder := e.ApiKeyRequest("POST", strRequest, mapParams)
	if err := json.Unmarshal([]byte(jsonCancelOrder), &cancelOrder); err != nil {
		return fmt.Errorf("%s CancelOrder Json Unmarshal Err: %v %v", e.GetName(), err, jsonCancelOrder)
	} else if !cancelOrder.Success {
		return fmt.Errorf("%s CancelOrder Failed: %v", e.GetName(), jsonCancelOrder)
	}
	return nil
}

func (e *Tradeogre) LoadPublicData(operation *exchange.PublicOperation) error {
	switch operation.Type {
	case exchange.GetTickerPrice:
		switch operation.Wallet {
		case exchange.SpotWallet:
			return e.doTickerPrice(operation)
		}
	//default: fmt.Println("yes")
	}
	return fmt.Errorf("LoadPublicData :: Operation type invalid: %+v", operation.Type)
}
func (e *Tradeogre) doTickerPrice(operation *exchange.PublicOperation) error {
	tickerPrice := PairsData{}
	strRequestUrl := "/markets"
	strUrl := API_URL + strRequestUrl
	jsonTickerPrice := exchange.HttpGetRequest(strUrl, nil)
	if err := json.Unmarshal([]byte(jsonTickerPrice), &tickerPrice); err != nil {
		return fmt.Errorf("%s doTickerPrice Result Unmarshal Err: %v %s", e.GetName(), err, jsonTickerPrice)
	}
	//fmt.Println(tickerPrice)
	if operation.DebugMode {
		operation.RequestURI = API_URL
		operation.CallResponce = jsonTickerPrice
	}
	operation.TickerPrice = []*exchange.TickerPriceDetail{}
	for _, ticker := range tickerPrice {
		for name,tp := range ticker {
			price64,_ := strconv.ParseFloat(tp.Price, 64)
			tpd := &exchange.TickerPriceDetail{
				Pair:  e.GetPairBySymbol(name),
				Price: price64,
				Initialprice: tp.Initialprice,
				High:  tp.High,
				Low:   tp.Low,
				Volume: tp.Volume,
			}
			operation.TickerPrice = append(operation.TickerPrice, tpd)
		}
		//operation.TickerPrice = append(operation.TickerPrice, tpd)
	}
	/*for _,i := range 	operation.TickerPrice{
		fmt.Println(operation.TickerPrice)
	}*/

	/* method to find specific ticker price
	operation.TickerPrice = []*exchange.TickerPriceDetail{}
	pairsDetail := PairsDetail{}

	strRequestUrl := "/ticker/BTC-WOW"
	strUrl := API_URL + strRequestUrl
	fmt.Println(strUrl)
	jsonSymbolsReturn := exchange.HttpGetRequest(strUrl, nil)
	fmt.Println(jsonSymbolsReturn)
	if err := json.Unmarshal([]byte(jsonSymbolsReturn), &pairsDetail); err != nil {
		return fmt.Errorf("%s Get Coin detail Json Unmarshal Err: %v %v", e.GetName(), err, jsonSymbolsReturn)
	}
	y,_ := strconv.ParseFloat(pairsDetail.Price, 64)
	x := &exchange.TickerPriceDetail{
		Pair: pair.GetPairByKey("BTC|WOW"),
		Price: y,
	}
	operation.TickerPrice = append(operation.TickerPrice, x)
	//fmt.Println(operation.TickerPrice[0])
	//operation.TickerPrice[0].Price,_ = strconv.ParseFloat(pairsDetail.Price, 64)
  fmt.Println(pairsDetail.Price)
	*/
  return nil
}

func (e *Tradeogre) UpdateAllBalances() {
	if e.API_KEY == "" || e.API_SECRET == "" {
	  fmt.Errorf("%s API Key or Secret Key are nil", e.GetName())
	}

	accountBalance := AccountBalances{}
	strRequest := "/account/balances"

	jsonBalanceReturn := e.ApiKeyRequest("GET", strRequest, make(map[string]string))
	if err := json.Unmarshal([]byte(jsonBalanceReturn), &accountBalance); err != nil {
		fmt.Errorf("%s UpdateAllBalances Json Unmarshal Err: %v %v", e.GetName(), err, jsonBalanceReturn)
	} else if !accountBalance.Success {
		fmt.Errorf("%s UpdateAllBalances Failed: %v", e.GetName(), jsonBalanceReturn)
	}
	for key, data := range accountBalance.Balances {
		c := e.GetCoinBySymbol(key)
		if c != nil {
			freeBalance, err := strconv.ParseFloat(data, 64)
			if err != nil {
				fmt.Errorf("%s balance parse error: %v, %v", e.GetName(), err, data)
			}
			balanceMap.Set(c.Code, freeBalance)
		}
	}
}

func (e *Tradeogre) Withdraw(coin *coin.Coin, quantity float64, addr, tag string) bool {

	return false
}

func (e *Tradeogre) LimitSell(pair *pair.Pair, quantity, rate float64) (*exchange.Order, error) {
	if e.API_KEY == "" || e.API_SECRET == "" {
		return nil, fmt.Errorf("%s API Key or Secret Key are nil", e.GetName())
	}

	placeOrder := PlaceOrder{}
	strRequest := "/order/sell"

	mapParams := make(map[string]string)
	mapParams["quantity"] = strconv.FormatFloat(quantity, 'f', -1, 64)
	mapParams["price"] = strconv.FormatFloat(rate, 'f', -1, 64)
	mapParams["market"] = e.GetSymbolByPair(pair)

	jsonPlaceReturn := e.ApiKeyRequest("POST", strRequest, mapParams)
	if err := json.Unmarshal([]byte(jsonPlaceReturn), &placeOrder); err != nil {
		return nil, fmt.Errorf("%s LimitSell Json Unmarshal Err: %v %v", e.GetName(), err, jsonPlaceReturn)
	} else if !placeOrder.Success {
		return nil, fmt.Errorf("%s LimitSell Failed: %v", e.GetName(), jsonPlaceReturn)
	}

	order := &exchange.Order{
		Pair:         pair,
		OrderID:      placeOrder.UUID,
		Rate:         rate,
		Quantity:     quantity,
		Direction:    exchange.Sell,
		Status:       exchange.New,
		JsonResponse: jsonPlaceReturn,
	}

	return order, nil
}

func (e *Tradeogre) LimitBuy(pair *pair.Pair, quantity, rate float64) (*exchange.Order, error) {
	if e.API_KEY == "" || e.API_SECRET == "" {
		return nil, fmt.Errorf("%s API Key or Secret Key are nil", e.GetName())
	}

	placeOrder := PlaceOrder{}
	strRequest := "/order/buy"

	mapParams := make(map[string]string)
	mapParams["quantity"] = strconv.FormatFloat(quantity, 'f', -1, 64)
	mapParams["price"] = strconv.FormatFloat(rate, 'f', -1, 64)
	mapParams["market"] = e.GetSymbolByPair(pair)

	jsonPlaceReturn := e.ApiKeyRequest("POST", strRequest, mapParams)
	if err := json.Unmarshal([]byte(jsonPlaceReturn), &placeOrder); err != nil {
		return nil, fmt.Errorf("%s LimitBuy Json Unmarshal Err: %v %v", e.GetName(), err, jsonPlaceReturn)
	} else if !placeOrder.Success {
		return nil, fmt.Errorf("%s LimitBuy Failed: %v", e.GetName(), jsonPlaceReturn)
	}

	order := &exchange.Order{
		Pair:         pair,
		OrderID:      placeOrder.UUID,
		Rate:         rate,
		Quantity:     quantity,
		Direction:    exchange.Buy,
		Status:       exchange.New,
		JsonResponse: jsonPlaceReturn,
	}
// example json response: {"success":true,"uuid":"c07ea117-33ea-4f45-6c18-a6d195ae4bdb","bnewbalavail":"0.00040764","snewbalavail":"1850.37346666"}

	return order, nil
}

func (e *Tradeogre) OrderStatus(order *exchange.Order) error {
	if e.API_KEY == "" || e.API_SECRET == "" {
		return fmt.Errorf("%s API Key or Secret Key are nil", e.GetName())
	}

	orderStatus := OrderStatus{}
	strRequest := fmt.Sprintf("/account/order/%v", order.OrderID)

	jsonOrderStatus := e.ApiKeyRequest("GET", strRequest, make(map[string]string))
	if err := json.Unmarshal([]byte(jsonOrderStatus), &orderStatus); err != nil {
		return fmt.Errorf("%s OrderStatus Json Unmarshal Err: %v %v", e.GetName(), err, jsonOrderStatus)
	} else if !orderStatus.Success {
		log.Printf("orderStatus.Error:%+v", orderStatus.Error)
		if orderStatus.Error == "Order not found" { //temperory solution , tradeogre's bug: when filled, the order uuid can't be traced. this solution will cause all not found orders shown as filled.
			order.Status = exchange.Filled
			return nil
		} else {
			return fmt.Errorf("%s OrderStatus Failed: %v", e.GetName(), jsonOrderStatus)
		}
	}

	order.StatusMessage = jsonOrderStatus

	fulfilledFloat, err := strconv.ParseFloat(orderStatus.Fulfilled, 64)
	if err != nil {
		return fmt.Errorf("%s orderStatus parse error: %v, %v", e.GetName(), err, orderStatus.Fulfilled)
	}
	if fulfilledFloat == order.Quantity {
		order.Status = exchange.Filled
	} else if fulfilledFloat > 0 && fulfilledFloat < order.Quantity {
		order.Status = exchange.Partial
	} else if fulfilledFloat == 0 {
		order.Status = exchange.New
	} else {
		order.Status = exchange.Other
	}

	return nil
}

func (e *Tradeogre) ListOrders() ([]*exchange.Order, error) {
	//!- only Testing for orders
	if e.API_KEY == "" || e.API_SECRET == "" {
		return nil, fmt.Errorf("%s API Key or Secret Key are nil", e.GetName())
	}

	strRequest := "/account/orders"

	mapParams := make(map[string]string)
	// mapParams["market"] = "BTC-RVN"
  var listOrders []Order
	jsonListOrders := e.ApiKeyRequest("POST", strRequest, mapParams)

	if err := json.Unmarshal([]byte(jsonListOrders), &listOrders); err != nil {
		return nil, fmt.Errorf("%s ListOrders Json Unmarshal Err: %v %v\n", e.GetName(), err, jsonListOrders)
	}
	var res []*exchange.Order
	for _, orderItem := range listOrders {
		pair := e.GetPairBySymbol(orderItem.Market)
		rate, _ := strconv.ParseFloat(orderItem.Price, 64)
		quantity, _ := strconv.ParseFloat(orderItem.Quantity, 64)
		Timestamp := int64(orderItem.Date)

		order := &exchange.Order{
			Pair:         pair,
			OrderID:      string(orderItem.UUID),
			Timestamp:    Timestamp,
			Rate:         rate,
			Quantity:     quantity,
			Status:       exchange.New,
			JsonResponse: jsonListOrders,
		}

		switch orderItem.Type {
		case "buy":
			order.Direction = exchange.Buy
		case "sell":
			order.Direction = exchange.Sell
		}

		res = append(res, order)
	}
  //fmt.Println(res)
  //fmt.Println(listOrders)
//	log.Printf("json:%+v", json)
//        return jsonListOrders, nil
	return res, nil
}

func (e *Tradeogre) CancelOrder(order *exchange.Order) error {
	if e.API_KEY == "" || e.API_SECRET == "" {
		return fmt.Errorf("%s API Key or Secret Key are nil", e.GetName())
	}

	cancelOrder := CancelOrder{}
	strRequest := "/order/cancel"

	mapParams := make(map[string]string)
	mapParams["uuid"] = order.OrderID

	jsonCancelOrder := e.ApiKeyRequest("POST", strRequest, mapParams)
	if err := json.Unmarshal([]byte(jsonCancelOrder), &cancelOrder); err != nil {
		return fmt.Errorf("%s CancelOrder Json Unmarshal Err: %v %v", e.GetName(), err, jsonCancelOrder)
	} else if !cancelOrder.Success {
		return fmt.Errorf("%s CancelOrder Failed: %v", e.GetName(), jsonCancelOrder)
	}

	order.Status = exchange.Canceling
	order.CancelStatus = jsonCancelOrder

	return nil
}

func (e *Tradeogre) CancelAllOrder() error {
	return nil
}

/*************** Signature Http Request ***************/
/*Method: API Request and Signature is required
Step 1: Change Instance Name    (e *<exchange Instance Name>)
Step 2: Create mapParams Depend on API Signature request
Step 3: Add HttpGetRequest below strUrl if API has different requests*/
func (e *Tradeogre) ApiKeyRequest(strMethod, strRequestPath string, mapParams map[string]string) string {

	strUrl := "https://" + e.API_KEY + ":" + e.API_SECRET + "@tradeogre.com/api/v1" + strRequestPath

	if strMethod == "GET" {
		return exchange.HttpGetRequest(strUrl, mapParams)
	}

	httpClient := &http.Client{}
	req, err := http.NewRequest(strMethod, strUrl, strings.NewReader(exchange.Map2UrlQuery(mapParams)))
	if err != nil {
		return err.Error()
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	response, err := httpClient.Do(req)
	if err != nil {
		log.Printf("err=%v", err)
		return err.Error()
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err.Error()
	}
	return string(body)
}
