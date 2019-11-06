package api

import (
	"strings"
	"time"
	"sort"
	"encoding/json"
	"net/http"
	"sync/atomic"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/potix/gobitflyer/client"
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/api/public"
	"github.com/potix/gobitflyer/api/private"
	"github.com/potix/gobitflyer/api/realtime"
)

const (
        BFCallableAPISpanSeconds int64 = 300
        BFCallableAPICount       int64 = 500
)

const (
	apiEndpoint         string = "https://api.bitflyer.jp"
	realtimeApiEndpoint string = "wss://ws.lightstream.bitflyer.com/json-rpc"
)

type APIClient struct {
	endpoint                  string
	httpClient                *client.HTTPClient
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

func NewAPIClient(httpClient *client.HTTPClient, authenticator Authenticator) (*APIClient) {
	return &APIClient{
		endpoint:                  apiEndpoint,
		httpClient:                httpClient,
		authenticator:             authenticator,
	}
}

type RealAPIClient struct {
	endpoint                  string
	wsClient                  *client.WSClient
	apiClient                 *APIClient
	realtimeChannel           *realtime.RealtimeChannel
}

func (c *RealAPIClient) RealBoardSnapshotCallback(conn *websocket.Conn, calbackData interface{}) (error) {
	rc := (calbackData).(*realtime.RealtimeChannel)
	select {
	case d := <-rc.UnsubscribeChan:
		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		err := conn.WriteJSON(d)
		if err != nil {
			atomic.StoreUint32(&rc.Subscribed, 0)
			return errors.Wrapf(err, "can not write unsubscribed")
		}
		atomic.StoreUint32(&rc.Subscribed, 0)
		return nil
	default:
		if atomic.LoadUint32(&rc.Subscribed) == 0 {
			conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			err := conn.WriteJSON(&realtime.JsonRPC2Subscribe{
					JsonRpc: "2.0",
					Method:  "subscribe",
					Params:  realtime.JsonRPC2SubscribeParams{
						Channel: "lightning_board_snapshot_" + string(rc.ProductCode),
					},
				})
			if err != nil {
				atomic.StoreUint32(&rc.Subscribed, 0)
				return errors.Wrapf(err, "can not write subscribe")
			}
			atomic.StoreUint32(&rc.Subscribed, 1)
		} else {
			notify := new(realtime.JsonRPC2BoardSnapshotNotify)
			conn.SetReadDeadline(time.Now().Add(time.Minute))
			err := conn.ReadJSON(notify);
			if err != nil {
				atomic.StoreUint32(&rc.Subscribed, 0)
				return errors.Wrapf(err, "can not read message")
			}
			rc.BoardSnapshotCallback(rc.ProductCode, notify.Params.Message, rc.CallbackData)
		}
		return nil
	}
}

func (c *RealAPIClient) RealBoardSnapshotStart(productCode types.ProductCode, callback realtime.BoardSnapshotCallback, callbackData interface{}) (error){
	if c.realtimeChannel != nil {
		return errors.Errorf("already exists realtime api connection")
	}
	wsRequest := &client.WSRequest {
		URL: c.endpoint,
		Headers: make(map[string]string),
	}
	rc := &realtime.RealtimeChannel{
		ProductCode:           productCode,
		RealtimeType:          types.RealtimeTypeBoardSnapshot,
                BoardSnapshotCallback: callback,
                BoardCallback:         nil,
                TickerCallback:        nil,
                ExecutionsCallback:    nil,
		CallbackData:          callbackData,
		Subscribed:            0,
		UnsubscribeChan:       make(chan *realtime.JsonRPC2Subscribe),
		Merge:                 false,
		GetBoardResponseFull:  nil,
	}
	err := c.wsClient.Start(wsRequest, c.RealBoardSnapshotCallback, rc)
	if err != nil {
		return errors.Wrapf(err, "can not connect realtime api")
	}
	c.realtimeChannel = rc
	return nil
}

func (c *RealAPIClient) realBoardCallbackMerge(rc *realtime.RealtimeChannel, getBoardResponseDiff *public.GetBoardResponse) (*public.GetBoardResponse) {
	rc.GetBoardResponseFull.MidPrice = getBoardResponseDiff.MidPrice
	for _, diffAsk := range getBoardResponseDiff.Asks {
		if diffAsk.Price == 0 {
			continue
		}
		if diffAsk.Size == 0 {
			for i := 0; i < len(rc.GetBoardResponseFull.Asks); i+= 1 {
				if rc.GetBoardResponseFull.Asks[i].Price == diffAsk.Price {
					rc.GetBoardResponseFull.Asks = append(rc.GetBoardResponseFull.Asks[:i], rc.GetBoardResponseFull.Asks[i+1:]...)
					break
				}
			}
		} else  {
			var i int
			for i = 0; i < len(rc.GetBoardResponseFull.Asks); i+= 1 {
				if rc.GetBoardResponseFull.Asks[i].Price == diffAsk.Price {
					rc.GetBoardResponseFull.Asks[i].Size = diffAsk.Size
					break
				}
			}
			if i == len(rc.GetBoardResponseFull.Asks) {
				rc.GetBoardResponseFull.Asks = append(rc.GetBoardResponseFull.Asks, diffAsk)
			}
		}
	}
	for _, diffBid := range getBoardResponseDiff.Bids {
		if diffBid.Price == 0 {
			continue
		}
		if diffBid.Size == 0 {
			for i := 0; i < len(rc.GetBoardResponseFull.Bids); i+= 1 {
				if rc.GetBoardResponseFull.Bids[i].Price == diffBid.Price {
					rc.GetBoardResponseFull.Bids = append(rc.GetBoardResponseFull.Bids[:i], rc.GetBoardResponseFull.Bids[i+1:]...)
					break
				}
			}
		} else  {
			var i int
			for i = 0; i < len(rc.GetBoardResponseFull.Bids); i+= 1 {
				if rc.GetBoardResponseFull.Bids[i].Price == diffBid.Price {
					rc.GetBoardResponseFull.Bids[i].Size = diffBid.Size
					break
				}
			}
			if i == len(rc.GetBoardResponseFull.Bids) {
				rc.GetBoardResponseFull.Bids = append(rc.GetBoardResponseFull.Bids, diffBid)
			}
		}
	}
	sort.Slice(rc.GetBoardResponseFull.Asks, func (i int, j int) bool {
		return rc.GetBoardResponseFull.Asks[i].Price < rc.GetBoardResponseFull.Asks[j].Price
	})
	sort.Slice(rc.GetBoardResponseFull.Bids, func (i int, j int) bool {
		return rc.GetBoardResponseFull.Bids[i].Price > rc.GetBoardResponseFull.Bids[j].Price
	})
	return rc.GetBoardResponseFull.Clone()
}

func (c *RealAPIClient) realBoardCallback(conn *websocket.Conn, calbackData interface{}) (error) {
	rc := (calbackData).(*realtime.RealtimeChannel)
	select {
	case d := <-rc.UnsubscribeChan:
		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		err := conn.WriteJSON(d)
		if err != nil {
			atomic.StoreUint32(&rc.Subscribed, 0)
			return errors.Wrapf(err, "can not write unsubscribed")
		}
		atomic.StoreUint32(&rc.Subscribed, 0)
		return nil
	default:
		if atomic.LoadUint32(&rc.Subscribed) == 0 {
			if rc.Merge {
				conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
				err := conn.WriteJSON(&realtime.JsonRPC2Subscribe{
						JsonRpc: "2.0",
						Method: "subscribe",
						Params: realtime.JsonRPC2SubscribeParams{
							Channel: "lightning_board_snapshot_" + string(rc.ProductCode),
						},
					})
				if err != nil {
					atomic.StoreUint32(&rc.Subscribed, 0)
					return errors.Wrapf(err, "can not write subscribe")
				}
			}
			conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			err := conn.WriteJSON(&realtime.JsonRPC2Subscribe{
					JsonRpc: "2.0",
					Method: "subscribe",
					Params: realtime.JsonRPC2SubscribeParams{
						Channel: "lightning_board_" + string(rc.ProductCode),
					},
				})
			if err != nil {
				return errors.Wrapf(err, "can not write subscribe")
			}
			atomic.StoreUint32(&rc.Subscribed, 1)
		} else {
			notify := new(realtime.JsonRPC2BoardNotify)
			conn.SetReadDeadline(time.Now().Add(time.Minute))
			err := conn.ReadJSON(notify);
			if err != nil {
				atomic.StoreUint32(&rc.Subscribed, 0)
				return errors.Wrapf(err, "can not read message")
			}
			if rc.Merge {
				if strings.Contains(notify.Params.Channel, "lightning_board_snapshot") {
					rc.GetBoardResponseFull = notify.Params.Message
					rc.BoardCallback(rc.ProductCode, notify.Params.Message, rc.CallbackData)
				} else if rc.GetBoardResponseFull != nil {
					getBoardResponse := c.realBoardCallbackMerge(rc, notify.Params.Message)
					rc.BoardCallback(rc.ProductCode, getBoardResponse, rc.CallbackData)
				}
			} else {
				rc.BoardCallback(rc.ProductCode, notify.Params.Message, rc.CallbackData)
			}
		}
		return nil
	}
}

func (c *RealAPIClient) RealBoardStart(productCode types.ProductCode, callback realtime.BoardCallback, callbackData interface{}, merge bool) (error){
	if c.realtimeChannel != nil {
		return errors.Errorf("already exists realtime api connection")
	}
	wsRequest := &client.WSRequest {
		URL: c.endpoint,
		Headers: make(map[string]string),
	}
	rc := &realtime.RealtimeChannel{
		ProductCode:           productCode,
		RealtimeType:          types.RealtimeTypeBoard,
                BoardSnapshotCallback: nil,
                BoardCallback:         callback,
                TickerCallback:        nil,
                ExecutionsCallback:    nil,
		CallbackData:          callbackData,
		Subscribed:            0,
		UnsubscribeChan:       make(chan *realtime.JsonRPC2Subscribe),
		Merge:                 merge,
		GetBoardResponseFull:  nil,
	}
	err := c.wsClient.Start(wsRequest, c.realBoardCallback, rc)
	if err != nil {
		return errors.Wrapf(err, "can not connect realtime api")
	}
	c.realtimeChannel = rc
	return nil
}

func (c *RealAPIClient) realTickerCallback(conn *websocket.Conn, calbackData interface{}) (error) {
	rc := (calbackData).(*realtime.RealtimeChannel)
	select {
	case d := <-rc.UnsubscribeChan:
		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		err := conn.WriteJSON(d)
		if err != nil {
			atomic.StoreUint32(&rc.Subscribed, 0)
			return errors.Wrapf(err, "can not write unsubscribed")
		}
		atomic.StoreUint32(&rc.Subscribed, 0)
		return nil
	default:
		if atomic.LoadUint32(&rc.Subscribed) == 0 {
			conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			err := conn.WriteJSON(&realtime.JsonRPC2Subscribe{
					JsonRpc: "2.0",
					Method: "subscribe",
					Params: realtime.JsonRPC2SubscribeParams{
						Channel: "lightning_ticker_" + string(rc.ProductCode),
					},
				})
			if err != nil {
				atomic.StoreUint32(&rc.Subscribed, 0)
				return errors.Wrapf(err, "can not write subscribe")
			}
			atomic.StoreUint32(&rc.Subscribed, 1)
		} else {
			notify := new(realtime.JsonRPC2TickerNotify)
			conn.SetReadDeadline(time.Now().Add(time.Minute))
			err := conn.ReadJSON(notify);
			if err != nil {
				atomic.StoreUint32(&rc.Subscribed, 0)
				return errors.Wrapf(err, "can not read message")
			}
			rc.TickerCallback(rc.ProductCode, notify.Params.Message, rc.CallbackData)
		}
		return nil
	}
}

func (c *RealAPIClient) RealTickerStart(productCode types.ProductCode, callback realtime.TickerCallback, callbackData interface{}) (error){
	if c.realtimeChannel != nil {
		return errors.Errorf("already exists realtime api connection")
	}
	wsRequest := &client.WSRequest {
		URL: c.endpoint,
		Headers: make(map[string]string),
	}
	rc := &realtime.RealtimeChannel{
		ProductCode:           productCode,
		RealtimeType:          types.RealtimeTypeTicker,
                BoardSnapshotCallback: nil,
                BoardCallback:         nil,
                TickerCallback:        callback,
                ExecutionsCallback:    nil,
		CallbackData:          callbackData,
		Subscribed:            0,
		UnsubscribeChan:       make(chan *realtime.JsonRPC2Subscribe),
		Merge:                 false,
		GetBoardResponseFull:  nil,
	}
	err := c.wsClient.Start(wsRequest, c.realTickerCallback, rc)
	if err != nil {
		return errors.Wrapf(err, "can not connect realtime api")
	}
	c.realtimeChannel = rc
	return nil
}

func (c *RealAPIClient) realExecutionsCallback(conn *websocket.Conn, calbackData interface{}) (error) {
	rc := (calbackData).(*realtime.RealtimeChannel)
	select {
	case d := <-rc.UnsubscribeChan:
		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		err := conn.WriteJSON(d)
		if err != nil {
			atomic.StoreUint32(&rc.Subscribed, 0)
			return errors.Wrapf(err, "can not write unsubscribed")
		}
		atomic.StoreUint32(&rc.Subscribed, 0)
		return nil
	default:
		if atomic.LoadUint32(&rc.Subscribed) == 0 {
			conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
			err := conn.WriteJSON(&realtime.JsonRPC2Subscribe{
					JsonRpc: "2.0",
					Method: "subscribe",
					Params: realtime.JsonRPC2SubscribeParams{
						Channel: "lightning_executions_" + string(rc.ProductCode),
					},
				})
			if err != nil {
				atomic.StoreUint32(&rc.Subscribed, 0)
				return errors.Wrapf(err, "can not write subscribe")
			}
			atomic.StoreUint32(&rc.Subscribed, 1)
		} else {
			notify := new(realtime.JsonRPC2ExecutionsNotify)
			conn.SetReadDeadline(time.Now().Add(time.Minute))
			err := conn.ReadJSON(notify);
			if err != nil {
				atomic.StoreUint32(&rc.Subscribed, 0)
				return errors.Wrapf(err, "can not read message")
			}
			rc.ExecutionsCallback(rc.ProductCode, notify.Params.Message, rc.CallbackData)
		}
		return nil
	}
}

func (c *RealAPIClient) RealExecutionsStart(productCode types.ProductCode, callback realtime.ExecutionsCallback, callbackData interface{}) (error){
	if c.realtimeChannel != nil {
		return errors.Errorf("already exists realtime api connection")
	}
	wsRequest := &client.WSRequest {
		URL: c.endpoint,
		Headers: make(map[string]string),
	}
	rc := &realtime.RealtimeChannel{
		ProductCode:           productCode,
		RealtimeType:          types.RealtimeTypeExecutions,
                BoardSnapshotCallback: nil,
                BoardCallback:         nil,
                TickerCallback:        nil,
                ExecutionsCallback:    callback,
		CallbackData:          callbackData,
		Subscribed:            0,
		UnsubscribeChan:       make(chan *realtime.JsonRPC2Subscribe),
		Merge:                 false,
		GetBoardResponseFull:  nil,
	}
	err := c.wsClient.Start(wsRequest, c.realExecutionsCallback, rc)
	if err != nil {
		return errors.Wrapf(err, "can not connect realtime api")
	}
	c.realtimeChannel = rc
	return nil
}

func (c *RealAPIClient) RealStop() (error) {
	if c.realtimeChannel == nil {
		return errors.Errorf("not found realtime api connection")
	}
	if atomic.LoadUint32(&c.realtimeChannel.Subscribed) == 1 {
		switch  c.realtimeChannel.RealtimeType {
		case types.RealtimeTypeBoardSnapshot:
			c.realtimeChannel.UnsubscribeChan <- &realtime.JsonRPC2Subscribe{
									JsonRpc: "2.0",
									Method:  "unsubscribe",
									Params:  realtime.JsonRPC2SubscribeParams{
										Channel: "lightning_board_snapshot_" + string(c.realtimeChannel.ProductCode),
									},
								}
		case types.RealtimeTypeBoard:
			c.realtimeChannel.UnsubscribeChan <- &realtime.JsonRPC2Subscribe{
									JsonRpc: "2.0",
									Method:  "unsubscribe",
									Params:  realtime.JsonRPC2SubscribeParams{
										Channel: "lightning_board_" + string(c.realtimeChannel.ProductCode),
									},
								}
			if c.realtimeChannel.Merge {
				c.realtimeChannel.UnsubscribeChan <- &realtime.JsonRPC2Subscribe{
										JsonRpc: "2.0",
										Method:  "unsubscribe",
										Params:  realtime.JsonRPC2SubscribeParams{
											Channel: "lightning_board_snapshot_" + string(c.realtimeChannel.ProductCode),
										},
									}
			}
		case types.RealtimeTypeTicker:
			c.realtimeChannel.UnsubscribeChan <- &realtime.JsonRPC2Subscribe{
									JsonRpc: "2.0",
									Method:  "unsubscribe",
									Params:  realtime.JsonRPC2SubscribeParams{
										Channel: "lightning_ticker_" + string(c.realtimeChannel.ProductCode),
									},
								}
		case types.RealtimeTypeExecutions:
			c.realtimeChannel.UnsubscribeChan <- &realtime.JsonRPC2Subscribe{
									JsonRpc: "2.0",
									Method:  "unsubscribe",
									Params:  realtime.JsonRPC2SubscribeParams{
										Channel: "lightning_executions_" + string(c.realtimeChannel.ProductCode),
									},
								}
		}
	}
	c.wsClient.Stop()
	close(c.realtimeChannel.UnsubscribeChan)
	c.realtimeChannel = nil
	return nil
}

func NewRealAPIClient(wsClient *client.WSClient) (*RealAPIClient) {
	return &RealAPIClient{
		endpoint:                  realtimeApiEndpoint,
		wsClient:                  wsClient,
		realtimeChannel:           nil,
	}
}

