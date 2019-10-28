package api

import (
	"time"
	"encoding/json"
	"net/http"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/potix/gobitflyer/client"
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/api/public"
	"github.com/potix/gobitflyer/api/private"
)

type RealTickerCallback func(getTickerResponse *public.GetTickerResponse, callbackData interface{})
type RealBoardSnapshotCallback func(getBoardResponse *public.GetBoardResponse, callbackData interface{})

type realTickerChannel struct {
	wsClient      *client.WSClient
	callback      RealTickerCallback
	callbackData  interface{}
	subscribed    bool
	writeDataChan chan *JsonRPC2Subscribe
}

type realBoardSnapshotChannel struct {
	wsClient      *client.WSClient
	callback      RealBoardSnapshotCallback
	callbackData  interface{}
	subscribed    bool
	writeDataChan chan *JsonRPC2Subscribe
}

type JsonRPC2Subscribe struct {
	JsonRpc string                  `json:"jsonrpc"`
	Method  string                  `json:"method"`
	Params  JsonRPC2SubscribeParams `json:"params"`
}

type JsonRPC2SubscribeParams struct {
	Channel string `json:"channel"`
}

type JsonRPC2TickerNotify struct {
	JsonRpc string                      `json:"jsonrpc"`
	Method  string                      `json:"method"`
	Params  *JsonRPC2TickerNotifyParams `json:"params"`
}

type JsonRPC2TickerNotifyParams struct {
	Channel string                    `json:"channel"`
	Message *public.GetTickerResponse `json:"message"`
}

type JsonRPC2BoardSnapshotNotify struct {
	JsonRpc string                             `json:"jsonrpc"`
	Method  string                             `json:"method"`
	Params  *JsonRPC2BoardSnapshotNotifyParams `json:"params"`
}

type JsonRPC2BoardSnapshotNotifyParams struct {
	Channel string                   `json:"channel"`
	Message *public.GetBoardResponse `json:"message"`
}

type APIClient struct {
	endpoint                  string
	httpClient                *client.HTTPClient
	realTickerChannels        map[types.ProductCode]*realTickerChannel
	realBoardSnapshotChannels map[types.ProductCode]*realBoardSnapshotChannel
	authenticator             Authenticator
}

func (c *APIClient) containsStatus(candidates []int, statusCode int) (bool) {
	for _, c := range candidates {
		if c == statusCode {
			return true
		}
	}
	return false
}

func (c *APIClient) PubGetMarkets() (*http.Response, public.GetMarketsResponse, error) {
	getMarketsRequest := public.NewGetMarketsRequest()
	getMarketsResponse := make(public.GetMarketsResponse, 0)
	httpRequest, err := getMarketsRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of get markets")
	}
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of get markets (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of get markets (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, &getMarketsResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of get markets (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, getMarketsResponse, nil
}

func (c *APIClient) PubGetBoard(productCode types.ProductCode) (*http.Response, *public.GetBoardResponse, error) {
	getBoardRequest := public.NewGetBoardRequest(productCode)
	getBoardResponse := new(public.GetBoardResponse)
	httpRequest, err := getBoardRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of get board")
	}
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of get board (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of get board (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, getBoardResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of get board (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, getBoardResponse, nil
}

func (c *APIClient) PubGetTicker(productCode types.ProductCode) (*http.Response, *public.GetTickerResponse, error) {
	getTickerRequest := public.NewGetTickerRequest(productCode)
	getTickerResponse := new(public.GetTickerResponse)
	httpRequest, err := getTickerRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of get ticker")
	}
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of get ticker (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of get ticker (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, getTickerResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of get ticker (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, getTickerResponse, nil
}

func (c *APIClient) PubGetExecutions(productCode types.ProductCode, count int64, before int64, after int64) (*http.Response, public.GetExecutionsResponse, error) {
	getExecutionsRequest := public.NewGetExecutionsRequest(productCode, count, before, after)
	getExecutionsResponse := make(public.GetExecutionsResponse, 0)
	httpRequest, err := getExecutionsRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of get  executions")
	}
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of get executions (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of get executions (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, &getExecutionsResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of get executions (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, getExecutionsResponse, nil
}

func (c *APIClient) PubGetBoardState(productCode types.ProductCode) (*http.Response, *public.GetBoardStateResponse, error) {
	getBoardStateRequest := public.NewGetBoardStateRequest(productCode)
	getBoardStateResponse := new(public.GetBoardStateResponse)
	httpRequest, err := getBoardStateRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of get board state")
	}
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of get board state (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of get board state (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, getBoardStateResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of get board state (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, getBoardStateResponse, nil
}

func (c *APIClient) PubGetHealth(productCode types.ProductCode) (*http.Response, *public.GetHealthResponse, error) {
	getHealthRequest := public.NewGetHealthRequest(productCode)
	getHealthResponse := new(public.GetHealthResponse)
	httpRequest, err := getHealthRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of get health")
	}
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of get health (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of get health (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, getHealthResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of get health (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, getHealthResponse, nil
}

func (c *APIClient) PubGetChats(fromDate int64) (*http.Response, *public.GetChatsResponse, error) {
	getChatsRequest := public.NewGetChatsRequest(fromDate)
	getChatsResponse := new(public.GetChatsResponse)
	httpRequest, err := getChatsRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of get chats")
	}
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of get chats (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of get chats (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, getChatsResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of get chats (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, getChatsResponse, nil
}

func (c *APIClient) PriGetPermissions() (*http.Response, *private.GetPermissionsResponse, error) {
	getPermissionsRequest := private.NewGetPermissionsRequest()
	getPermissionsResponse := new(private.GetPermissionsResponse)
	httpRequest, err := getPermissionsRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of get permissions")
	}
	c.authenticator.SetAuthHeaders(httpRequest.Headers, time.Now(), httpRequest.Method, httpRequest.PathQuery, httpRequest.Body)
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of get permissions (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of get permissions (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, getPermissionsResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of get permissions (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, getPermissionsResponse, nil
}

func (c *APIClient) PriGetBalance() (*http.Response, private.GetBalanceResponse, error) {
	getBalanceRequest := private.NewGetBalanceRequest()
	getBalanceResponse := make(private.GetBalanceResponse, 0)
	httpRequest, err := getBalanceRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of get balance")
	}
	c.authenticator.SetAuthHeaders(httpRequest.Headers, time.Now(), httpRequest.Method, httpRequest.PathQuery, httpRequest.Body)
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of get balance (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code ofaget  balance (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, &getBalanceResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of get balance (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, getBalanceResponse, nil
}

func (c *APIClient) PriGetCollateral() (*http.Response, *private.GetCollateralResponse, error) {
	getCollateralRequest := private.NewGetCollateralRequest()
	getCollateralResponse := new(private.GetCollateralResponse)
	httpRequest, err := getCollateralRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of get collateral")
	}
	c.authenticator.SetAuthHeaders(httpRequest.Headers, time.Now(), httpRequest.Method, httpRequest.PathQuery, httpRequest.Body)
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of get collateral (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of get collateral (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, getCollateralResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of get collateral (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, getCollateralResponse, nil
}

func (c *APIClient) PriGetCollateralAccounts() (*http.Response, private.GetCollateralAccountsResponse, error) {
	getCollateralAccountsRequest := private.NewGetCollateralAccountsRequest()
	getCollateralAccountsResponse := make(private.GetCollateralAccountsResponse, 0)
	httpRequest, err := getCollateralAccountsRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of get collateral accounts")
	}
	c.authenticator.SetAuthHeaders(httpRequest.Headers, time.Now(), httpRequest.Method, httpRequest.PathQuery, httpRequest.Body)
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of get collateral accounts (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of get collateral accounts (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, &getCollateralAccountsResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of get collateral accounts (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, getCollateralAccountsResponse, nil
}

func (c *APIClient) PriSendChildOrder(productCode types.ProductCode,
                                           childOrderType types.OrderType,
                                           side types.Side,
                                           price float64,
                                           size float64,
                                           minuteToExpire int64,
                                           timeInForce types.TimeInForce) (*http.Response, *private.SendChildOrderResponse, error) {
	sendChildOrderRequest := private.NewSendChildOrderRequest(productCode, childOrderType, side, price, size, minuteToExpire, timeInForce)
	sendChildOrderResponse := new(private.SendChildOrderResponse)
	httpRequest, err := sendChildOrderRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of send child order")
	}
	c.authenticator.SetAuthHeaders(httpRequest.Headers, time.Now(), httpRequest.Method, httpRequest.PathQuery, httpRequest.Body)
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of send child order (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of send child order (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, sendChildOrderResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of send child order (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, sendChildOrderResponse, nil
}

func (c *APIClient) PriCancelChildOrder(productCode types.ProductCode, idType types.IdType, orderId string) (*http.Response, error) {
	cancelChildOrderRequest, err := private.NewCancelChildOrderRequest(productCode, idType, orderId)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create cancel child order request")
	}
	httpRequest, err := cancelChildOrderRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create http request of cancel child order")
	}
	c.authenticator.SetAuthHeaders(httpRequest.Headers, time.Now(), httpRequest.Method, httpRequest.PathQuery, httpRequest.Body)
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, errors.Wrapf(err, "can not request of cancel child order (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, errors.Errorf("unexpected status code of cancel child order (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	return httpResponse, nil
}

func (c *APIClient) PriGetChildOrders(productCode types.ProductCode, count int64, before int64, after int64, orderState types.OrderState) (*http.Response, private.GetChildOrdersResponse, error) {
	getChildOrdersRequest := private.NewGetChildOrdersRequest(productCode, count, before, after, orderState)
	getChildOrdersResponse := make(private.GetChildOrdersResponse, 0)
	httpRequest, err := getChildOrdersRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of get child orders")
	}
	c.authenticator.SetAuthHeaders(httpRequest.Headers, time.Now(), httpRequest.Method, httpRequest.PathQuery, httpRequest.Body)
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of get child orders (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of get child orders (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, &getChildOrdersResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of get child orders (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, getChildOrdersResponse, nil
}

func (c *APIClient) PriGetChildOrdersById(productCode types.ProductCode, idType types.IdType, orderId string) (*http.Response, private.GetChildOrdersResponse, error) {
	getChildOrdersRequest, err := private.NewGetChildOrdersRequestById(productCode, idType, orderId)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create get child orders by id request")
	}
	getChildOrdersResponse := make(private.GetChildOrdersResponse, 0)
	httpRequest, err := getChildOrdersRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of get child orders by id")
	}
	c.authenticator.SetAuthHeaders(httpRequest.Headers, time.Now(), httpRequest.Method, httpRequest.PathQuery, httpRequest.Body)
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of get child orders by id (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of get child orders by id (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, &getChildOrdersResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of get child orders by id (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, getChildOrdersResponse, nil
}

func (c *APIClient) PriCancelAllChildOrders(productCode types.ProductCode) (*http.Response, error) {
	cancelAllChildOrdersRequest := private.NewCancelAllChildOrdersRequest(productCode)
	httpRequest, err := cancelAllChildOrdersRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create http request of cancel all child order")
	}
	c.authenticator.SetAuthHeaders(httpRequest.Headers, time.Now(), httpRequest.Method, httpRequest.PathQuery, httpRequest.Body)
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, errors.Wrapf(err, "can not request of cancel all child order (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, errors.Errorf("unexpected status code of cancel all child order (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	return httpResponse, nil
}

func (c *APIClient) PriGetExecutions(productCode types.ProductCode, count int64, before int64, after int64) (*http.Response, private.GetExecutionsResponse, error) {
	getExecutionsRequest := private.NewGetExecutionsRequest(productCode, count, before, after)
	getExecutionsResponse := make(private.GetExecutionsResponse, 0)
	httpRequest, err := getExecutionsRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of get executions")
	}
	c.authenticator.SetAuthHeaders(httpRequest.Headers, time.Now(), httpRequest.Method, httpRequest.PathQuery, httpRequest.Body)
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of get executions (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of get executions (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, &getExecutionsResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of get executions (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, getExecutionsResponse, nil
}

func (c *APIClient) PriGetExecutionsById(productCode types.ProductCode, idType types.IdType, orderId string) (*http.Response, private.GetExecutionsResponse, error) {
	getExecutionsRequest, err := private.NewGetExecutionsRequestById(productCode, idType, orderId)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create get executions by id request")
	}
	getExecutionsResponse := make(private.GetExecutionsResponse, 0)
	httpRequest, err := getExecutionsRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of get executions by id")
	}
	c.authenticator.SetAuthHeaders(httpRequest.Headers, time.Now(), httpRequest.Method, httpRequest.PathQuery, httpRequest.Body)
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of get executions by id (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of get executions by id (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, &getExecutionsResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of get executions by id (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, getExecutionsResponse, nil
}

func (c *APIClient) PriGetBalanceHistory(currencyCode types.CurrencyCode, count int64, before int64, after int64) (*http.Response, private.GetBalanceHistoryResponse, error) {
	getBalanceHistoryRequest := private.NewGetBalanceHistoryRequest(currencyCode, count, before, after)
	getBalanceHistoryResponse := make(private.GetBalanceHistoryResponse, 0)
	httpRequest, err := getBalanceHistoryRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of get balance history")
	}
	c.authenticator.SetAuthHeaders(httpRequest.Headers, time.Now(), httpRequest.Method, httpRequest.PathQuery, httpRequest.Body)
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of get balance history (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of get balance history (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, &getBalanceHistoryResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of get balance history (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, getBalanceHistoryResponse, nil
}

func (c *APIClient) PriGetPositions() (*http.Response, private.GetPositionsResponse, error) {
	getPositionsRequest := private.NewGetPositionsRequest()
	getPositionsResponse := make(private.GetPositionsResponse, 0)
	httpRequest, err := getPositionsRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of get positions")
	}
	c.authenticator.SetAuthHeaders(httpRequest.Headers, time.Now(), httpRequest.Method, httpRequest.PathQuery, httpRequest.Body)
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of get positions (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of get positions (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, &getPositionsResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of get positions (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, getPositionsResponse, nil
}

func (c *APIClient) PriGetCollateralHistory(count int64, before int64, after int64) (*http.Response, private.GetCollateralHistoryResponse, error) {
	getCollateralHistoryRequest := private.NewGetCollateralHistoryRequest(count, before, after)
	getCollateralHistoryResponse := make(private.GetCollateralHistoryResponse, 0)
	httpRequest, err := getCollateralHistoryRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of get collateral history")
	}
	c.authenticator.SetAuthHeaders(httpRequest.Headers, time.Now(), httpRequest.Method, httpRequest.PathQuery, httpRequest.Body)
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of get collateral history (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of get collateral history (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, &getCollateralHistoryResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of get collateral history (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, getCollateralHistoryResponse, nil
}

func (c *APIClient) PriGetTradingCommission(productCode types.ProductCode) (*http.Response, *private.GetTradingCommissionResponse, error) {
	getTradingCommissionRequest := private.NewGetTradingCommissionRequest(productCode)
	getTradingCommissionResponse := new(private.GetTradingCommissionResponse)
	httpRequest, err := getTradingCommissionRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of get trading commission")
	}
	c.authenticator.SetAuthHeaders(httpRequest.Headers, time.Now(), httpRequest.Method, httpRequest.PathQuery, httpRequest.Body)
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of get trading commission (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of get trading commission (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, getTradingCommissionResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of get trading commission (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, getTradingCommissionResponse, nil
}

func (c *APIClient) PriSendParentOrder(orderMethod types.OrderMethod,
				       minuteToRxpire int64,
				       timeInForce types.TimeInForce,
				       parameters ...*private.SendParentOrderParameter) (*http.Response, *private.SendParentOrderResponse, error) {
	sendParentOrderRequest := private.NewSendParentOrderRequest(orderMethod, minuteToRxpire, timeInForce, parameters...)
	sendParentOrderResponse := new(private.SendParentOrderResponse)
	httpRequest, err := sendParentOrderRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of send parent order")
	}
	c.authenticator.SetAuthHeaders(httpRequest.Headers, time.Now(), httpRequest.Method, httpRequest.PathQuery, httpRequest.Body)
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of send parent order (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of send parent order (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, sendParentOrderResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of send parent order (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, sendParentOrderResponse, nil
}

func (c *APIClient) PriCancelParentOrder(productCode types.ProductCode, idType types.IdType, orderId string) (*http.Response, error) {
	cancelParentOrderRequest, err := private.NewCancelParentOrderRequest(productCode, idType, orderId)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create cancel parent order request")
	}
	httpRequest, err := cancelParentOrderRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, errors.Wrapf(err, "can not create http request of cancel parent order")
	}
	c.authenticator.SetAuthHeaders(httpRequest.Headers, time.Now(), httpRequest.Method, httpRequest.PathQuery, httpRequest.Body)
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, errors.Wrapf(err, "can not request of cancel parent order (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, errors.Errorf("unexpected status code of cancel parent order (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	return httpResponse, nil
}

func (c *APIClient) PriGetParentOrders(productCode types.ProductCode, count int64, before int64, after int64, orderState types.OrderState) (*http.Response, private.GetParentOrdersResponse, error) {
	getParentOrdersRequest := private.NewGetParentOrdersRequest(productCode, count, before, after, orderState)
	getParentOrdersResponse := make(private.GetParentOrdersResponse, 0)
	httpRequest, err := getParentOrdersRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of get parent orders")
	}
	c.authenticator.SetAuthHeaders(httpRequest.Headers, time.Now(), httpRequest.Method, httpRequest.PathQuery, httpRequest.Body)
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of get parent orders (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of get parent orders (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, &getParentOrdersResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of get parent orders (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, getParentOrdersResponse, nil
}

func (c *APIClient) PriGetParentOrder(IdType types.IdType, orderId string) (*http.Response, *private.GetParentOrderResponse, error) {
	getParentOrderRequest, err := private.NewGetParentOrderRequest(IdType, orderId)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create get parent order request")
	}
	getParentOrderResponse := new(private.GetParentOrderResponse)
	httpRequest, err := getParentOrderRequest.CreateHTTPRequest(c.endpoint)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not create http request of get parent order")
	}
	c.authenticator.SetAuthHeaders(httpRequest.Headers, time.Now(), httpRequest.Method, httpRequest.PathQuery, httpRequest.Body)
	httpResponse, body, err := c.httpClient.DoRequest(httpRequest)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "can not request of get parent order (request = %v)", httpRequest.ToString())
	}
	if !c.containsStatus([]int{200}, httpResponse.StatusCode) {
		return nil, nil, errors.Errorf("unexpected status code of get parent order (request = %v, status = %v, body = %v)", httpRequest.ToString(), httpResponse.Status, string(body))
	}
	err = json.Unmarshal(body, getParentOrderResponse)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "unmarshal data of get parent order (request = %v, body = %v)", httpRequest.ToString(), string(body))
	}
	return httpResponse, getParentOrderResponse, nil
}

func (c *APIClient) RealTickerCallback(conn *websocket.Conn, calbackData interface{}) (error) {
	rtc := (calbackData).(*realTickerChannel)
	select {
	case d := <-rtc.writeDataChan:
		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		err := conn.WriteJSON(d)
		if err != nil {
			return errors.Wrapf(err, "can not write subscribe or unsubscribed")
		}
		if d.Method == "subscribe" {
			rtc.subscribed = true
		} else if d.Method == "unsubscribe" {
			rtc.subscribed = false
		}
		return nil
	default:
		if rtc.subscribed {
			notify := new(JsonRPC2TickerNotify)
			err := conn.ReadJSON(notify);
			if err != nil {
				return errors.Wrapf(err, "can not read message")
			}
			rtc.callback(notify.Params.Message, rtc.callbackData)
		}
		return nil
	}
}

func (c *APIClient) RealTickerStart(productCode types.ProductCode, callback RealTickerCallback, callbackData interface{}) (error){
	wsRequest := &client.WSRequest {
		URL: "wss://ws.lightstream.bitflyer.com/json-rpc",
		Headers: make(map[string]string),
	}
	wsClient := client.NewWSClient(0, 0, 60, 1)
	rc := &realTickerChannel{
		wsClient:     wsClient,
		callback:     callback,
		callbackData: callbackData,
		subscribed:   false,
		writeDataChan: make(chan *JsonRPC2Subscribe),
	}
	err := wsClient.Start(wsRequest, c.RealTickerCallback, rc)
	if err != nil {
		return errors.Wrapf(err, "can not connect realtime api")
	}
	rc.writeDataChan <- &JsonRPC2Subscribe{JsonRpc: "2.0", Method: "subscribe", Params: JsonRPC2SubscribeParams{Channel: "lightning_ticker_" + string(productCode)}}
	c.realTickerChannels[productCode] = rc
	return nil
}

func (c *APIClient) RealTickerStop(productCode types.ProductCode) (error) {
	rc, ok := c.realTickerChannels[productCode]
	if !ok {
		return errors.Errorf("not found realtime api connection")
	}
	rc.writeDataChan <- &JsonRPC2Subscribe{JsonRpc: "2.0", Method: "unsubscribe", Params: JsonRPC2SubscribeParams{Channel: "lightning_ticker_" + string(productCode)}}
	rc.wsClient.Stop()
	close(rc.writeDataChan)
	delete(c.realTickerChannels, productCode)
	return nil
}

func (c *APIClient) RealBoardSnapshotCallback(conn *websocket.Conn, calbackData interface{}) (error) {
	rtc := (calbackData).(*realBoardSnapshotChannel)
	select {
	case d := <-rtc.writeDataChan:
		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		err := conn.WriteJSON(d)
		if err != nil {
			return errors.Wrapf(err, "can not write subscribe or unsubscribed")
		}
		if d.Method == "subscribe" {
			rtc.subscribed = true
		} else if d.Method == "unsubscribe" {
			rtc.subscribed = false
		}
		return nil
	default:
		if rtc.subscribed {
			notify := new(JsonRPC2BoardSnapshotNotify)
			err := conn.ReadJSON(notify);
			if err != nil {
				return errors.Wrapf(err, "can not read message")
			}
			rtc.callback(notify.Params.Message, rtc.callbackData)
		}
		return nil
	}
}

func (c *APIClient) RealBoardSnapshotStart(productCode types.ProductCode, callback RealBoardSnapshotCallback, callbackData interface{}) (error){
	wsRequest := &client.WSRequest {
		URL: "wss://ws.lightstream.bitflyer.com/json-rpc",
		Headers: make(map[string]string),
	}
	wsClient := client.NewWSClient(0, 0, 60, 1)
	rc := &realBoardSnapshotChannel{
		wsClient:     wsClient,
		callback:     callback,
		callbackData: callbackData,
		subscribed:   false,
		writeDataChan: make(chan *JsonRPC2Subscribe),
	}
	err := wsClient.Start(wsRequest, c.RealBoardSnapshotCallback, rc)
	if err != nil {
		return errors.Wrapf(err, "can not connect realtime api")
	}
	rc.writeDataChan <- &JsonRPC2Subscribe{JsonRpc: "2.0", Method: "subscribe", Params: JsonRPC2SubscribeParams{Channel: "lightning_board_snapshot_" + string(productCode)}}
	c.realBoardSnapshotChannels[productCode] = rc
	return nil
}

func (c *APIClient) RealBoardSnapshotStop(productCode types.ProductCode) (error) {
	rc, ok := c.realBoardSnapshotChannels[productCode]
	if !ok {
		return errors.Errorf("not found realtime api connection")
	}
	rc.writeDataChan <- &JsonRPC2Subscribe{JsonRpc: "2.0", Method: "unsubscribe", Params: JsonRPC2SubscribeParams{Channel: "lightning_board_snapshot_" + string(productCode)}}
	rc.wsClient.Stop()
	close(rc.writeDataChan)
	delete(c.realBoardSnapshotChannels, productCode)
	return nil
}

func NewAPIClient(httpClient *client.HTTPClient, authenticator Authenticator) (*APIClient) {
	return &APIClient{
		endpoint:                 "https://api.bitflyer.jp",
		httpClient:               httpClient,
		realTickerChannels:       make(map[types.ProductCode]*realTickerChannel),
		realBoardSnapshotChannels: make(map[types.ProductCode]*realBoardSnapshotChannel),
		authenticator:            authenticator,
	}
}

