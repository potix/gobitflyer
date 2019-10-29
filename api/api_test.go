package api_test

import (
	"log"
	"fmt"
	"time"
	"testing"
	"github.com/potix/gobitflyer/client"
	"github.com/potix/gobitflyer/api/types"
	"github.com/potix/gobitflyer/api/public"
	"github.com/potix/gobitflyer/api/private"
	"github.com/potix/gobitflyer/api"
)

func createApiClient(t *testing.T) (*api.APIClient) {
	httpClient := client.NewHTTPClient(30, 0, 180, nil)
	authenticator, err := api.NewAuthenticator("apikey")
	if err != nil {
		t.Errorf("can not create authenticator: %v", err)
	}
	return api.NewAPIClient(httpClient, authenticator)
}

func TestPubMarkets(t *testing.T) {
	apiClient := createApiClient(t)
	httpResponse, getMarketsResponse, err :=  apiClient.PubGetMarkets()
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getMarketsResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v\n%#v", httpResponse, getMarketsResponse))
}

func TestPubBoard(t *testing.T) {
	apiClient := createApiClient(t)
	httpResponse, getBoardResponse, err :=  apiClient.PubGetBoard("BTC_JPY")
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getBoardResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v\n%#v", httpResponse, getBoardResponse))
}

func TestPubTicker(t *testing.T) {
	apiClient := createApiClient(t)
	httpResponse, getTickerResponse, err :=  apiClient.PubGetTicker("BTC_JPY")
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getTickerResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v\n%#v", httpResponse, getTickerResponse))
}

func TestPubExecutions(t *testing.T) {
	apiClient := createApiClient(t)
	httpResponse, getExecutionsResponse, err :=  apiClient.PubGetExecutions("BTC_JPY", 10, 0, 0)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getExecutionsResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v\n%#v", httpResponse, getExecutionsResponse))
}

func TestPubBoardState(t *testing.T) {
	apiClient := createApiClient(t)
	httpResponse, getBoardStateResponse, err :=  apiClient.PubGetBoardState("BTC_JPY")
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getBoardStateResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v\n%#v", httpResponse, getBoardStateResponse))
}

func TestPubHealth(t *testing.T) {
	apiClient := createApiClient(t)
	httpResponse, getHealthResponse, err :=  apiClient.PubGetHealth("BTC_JPY")
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getHealthResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v\n%#v", httpResponse, getHealthResponse))
}

func TestPubGetChats(t *testing.T) {
	apiClient := createApiClient(t)
	httpResponse, getChatsResponse, err :=  apiClient.PubGetChats(1)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getChatsResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v\n%#v", httpResponse, getChatsResponse))
}

func TestPriPermissions(t *testing.T) {
	apiClient := createApiClient(t)
	httpResponse, getPermissionsResponse, err := apiClient.PriGetPermissions()
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getPermissionsResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v\n%#v", httpResponse, getPermissionsResponse))
}

func TestPriBalance(t *testing.T) {
	apiClient := createApiClient(t)
	httpResponse, getBalanceResponse, err := apiClient.PriGetBalance()
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getBalanceResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v\n%#v", httpResponse, getBalanceResponse))
}

func TestPriCollateral(t *testing.T) {
	apiClient := createApiClient(t)
	httpResponse, getCollateralResponse, err := apiClient.PriGetCollateral()
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getCollateralResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v\n%#v", httpResponse, getCollateralResponse))
}

func TestPriCollateralAccounts(t *testing.T) {
	apiClient := createApiClient(t)
	httpResponse, getCollateralAccountsResponse, err := apiClient.PriGetCollateralAccounts()
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getCollateralAccountsResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v\n%#v", httpResponse, getCollateralAccountsResponse))
}

func TestPriSendChildOrder(t *testing.T) {
	apiClient := createApiClient(t)
	httpResponse, sendChildOrderResponse, err := apiClient.PriSendChildOrder("BTC_JPY", types.OrderTypeLimit, types.SideBuy, 550000, 0.11, 1, types.TimeInForceIOC)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if sendChildOrderResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v\n%#v", httpResponse, sendChildOrderResponse))
}

func TestPriCancelChildOrderWithOrderAcceptanceId(t *testing.T) {
	apiClient := createApiClient(t)
	httpResponse, sendChildOrderResponse, err := apiClient.PriSendChildOrder("BTC_JPY", types.OrderTypeLimit, types.SideBuy, 550000, 0.11, 1, types.TimeInForceFOK)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if sendChildOrderResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v\n%#v", httpResponse, sendChildOrderResponse))
	httpResponse, err = apiClient.PriCancelChildOrder("BTC_JPY", types.IdTypeChildOrderAcceptanceId, sendChildOrderResponse.ChildOrderAcceptanceId)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	t.Log(fmt.Sprintf("%#v", httpResponse))
}

func TestPriCancelAllChildOrders(t *testing.T) {
	apiClient := createApiClient(t)
	httpResponse, sendChildOrderResponse, err := apiClient.PriSendChildOrder("BTC_JPY", types.OrderTypeLimit, types.SideBuy, 550000, 0.12, 1, types.TimeInForceFOK)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if sendChildOrderResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v\n%#v", httpResponse, sendChildOrderResponse))
	httpResponse, err = apiClient.PriCancelAllChildOrders("BTC_JPY")
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	t.Log(fmt.Sprintf("%#v", httpResponse))
}

func TestPriGetChildOrders(t *testing.T) {
	apiClient := createApiClient(t)
	httpResponse, sendChildOrderResponse, err := apiClient.PriSendChildOrder("BTC_JPY", types.OrderTypeLimit, types.SideBuy, 550000, 0.12, 1, types.TimeInForceGTC)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if sendChildOrderResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v\n%#v", httpResponse, sendChildOrderResponse))

	time.Sleep(2 * time.Second)

	httpResponse, getChildOrdersResponse, err := apiClient.PriGetChildOrders("BTC_JPY", 10, 0, 0, types.OrderStateActive)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getChildOrdersResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v", httpResponse))
	t.Log(fmt.Sprintf("%#v", getChildOrdersResponse))
	if len(getChildOrdersResponse) == 0 {
		t.Errorf("no order")
	}

	childOrderId := getChildOrdersResponse[0].ChildOrderId

	httpResponse, getChildOrdersResponse, err = apiClient.PriGetChildOrdersById("BTC_JPY", types.IdTypeChildOrderAcceptanceId, sendChildOrderResponse.ChildOrderAcceptanceId)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getChildOrdersResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v", httpResponse))
	if len(getChildOrdersResponse) == 0 {
		t.Errorf("no order")
	}

	httpResponse, err = apiClient.PriCancelChildOrder("BTC_JPY", types.IdTypeChildOrderId, childOrderId)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	t.Log(fmt.Sprintf("%#v", httpResponse))

	time.Sleep(2 * time.Second)

	httpResponse, getChildOrdersResponse, err = apiClient.PriGetChildOrders("BTC_JPY", 10, 0, 0, types.OrderStateActive)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getChildOrdersResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v", httpResponse))
	if len(getChildOrdersResponse) != 0 {
		t.Errorf("exist order")
	}
}


func TestPriGetExecutions(t *testing.T) {
	apiClient := createApiClient(t)
	httpResponse, getExecutionsResponse, err := apiClient.PriGetExecutions("BTC_JPY", 10, 0, 0)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getExecutionsResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v", httpResponse))
	t.Log(fmt.Sprintf("%#v", getExecutionsResponse))

	httpResponse, getExecutionsResponse, err = apiClient.PriGetExecutionsById("BTC_JPY", types.IdTypeChildOrderAcceptanceId, "NOT EXISTS")
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getExecutionsResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v", httpResponse))
	t.Log(fmt.Sprintf("%#v", getExecutionsResponse))
}

func TestPriGetBalanceHistory(t *testing.T) {
	apiClient := createApiClient(t)
	httpResponse, getBalanceHistoryResponse, err := apiClient.PriGetBalanceHistory("JPY", 10, 0, 0)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getBalanceHistoryResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v", httpResponse))
	t.Log(fmt.Sprintf("%#v", getBalanceHistoryResponse))
}

func TestPriGetPositions(t *testing.T) {
	apiClient := createApiClient(t)
	httpResponse, getPositionsResponse, err := apiClient.PriGetPositions()
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getPositionsResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v", httpResponse))
	t.Log(fmt.Sprintf("%#v", getPositionsResponse))
}

func TestPriGetCollateralHistory(t *testing.T) {
	apiClient := createApiClient(t)
	httpResponse, getCollateralHistoryResponse, err := apiClient.PriGetCollateralHistory(10, 0, 0)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getCollateralHistoryResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v", httpResponse))
	t.Log(fmt.Sprintf("%#v", getCollateralHistoryResponse))
}

func TestPriGetTradingCommission(t *testing.T) {
	apiClient := createApiClient(t)
	httpResponse, getTradingCommissionResponse, err := apiClient.PriGetTradingCommission("BTC_JPY")
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getTradingCommissionResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v", httpResponse))
	t.Log(fmt.Sprintf("%#v", getTradingCommissionResponse))
}



func TestPriSendParentOrder(t *testing.T) {
	apiClient := createApiClient(t)
	sendParentOrderParameter1 := &private.SendParentOrderParameter {
		ProductCode: "BTC_JPY",
		ConditionType: types.ConditionTypeLimit,
		Side: types.SideBuy,
		Size: 0.11,
		Price: 550000,
	}
	sendParentOrderParameter2 := &private.SendParentOrderParameter {
		ProductCode: "BTC_JPY",
		ConditionType: types.ConditionTypeLimit,
		Side: types.SideSell,
		Size: 0.11,
		Price: 2000000,
	}
	httpResponse, sendParentOrderResponse, err := apiClient.PriSendParentOrder(types.OrderMethodIFD, 1, types.TimeInForceIOC, sendParentOrderParameter1, sendParentOrderParameter2)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if sendParentOrderResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v\n%#v", httpResponse, sendParentOrderResponse))
}

func TestPriCancelParentOrder(t *testing.T) {
	apiClient := createApiClient(t)
	sendParentOrderParameter1 := &private.SendParentOrderParameter {
		ProductCode: "BTC_JPY",
		ConditionType: types.ConditionTypeLimit,
		Side: types.SideBuy,
		Size: 0.11,
		Price: 550000,
	}
	sendParentOrderParameter2 := &private.SendParentOrderParameter {
		ProductCode: "BTC_JPY",
		ConditionType: types.ConditionTypeLimit,
		Side: types.SideSell,
		Size: 0.11,
		Price: 2000000,
	}
	httpResponse, sendParentOrderResponse, err := apiClient.PriSendParentOrder(types.OrderMethodIFD, 1, types.TimeInForceGTC, sendParentOrderParameter1, sendParentOrderParameter2)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if sendParentOrderResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v\n%#v", httpResponse, sendParentOrderResponse))

	httpResponse, err = apiClient.PriCancelParentOrder("BTC_JPY", types.IdTypeParentOrderAcceptanceId, sendParentOrderResponse.ParentOrderAcceptanceId)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	t.Log(fmt.Sprintf("%#v", httpResponse))

}

func TestPriGetParentOrders(t *testing.T) {
	apiClient := createApiClient(t)
	sendParentOrderParameter1 := &private.SendParentOrderParameter {
		ProductCode: "BTC_JPY",
		ConditionType: types.ConditionTypeLimit,
		Side: types.SideBuy,
		Size: 0.11,
		Price: 550000,
	}
	sendParentOrderParameter2 := &private.SendParentOrderParameter {
		ProductCode: "BTC_JPY",
		ConditionType: types.ConditionTypeLimit,
		Side: types.SideSell,
		Size: 0.11,
		Price: 2000000,
	}
	httpResponse, sendParentOrderResponse, err := apiClient.PriSendParentOrder(types.OrderMethodIFD, 1, types.TimeInForceGTC, sendParentOrderParameter1, sendParentOrderParameter2)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if sendParentOrderResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v\n%#v", httpResponse, sendParentOrderResponse))

	time.Sleep(2 * time.Second)

	httpResponse, getParentOrdersResponse, err := apiClient.PriGetParentOrders("BTC_JPY", 10, 0, 0, types.OrderStateActive)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getParentOrdersResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v", httpResponse))
	t.Log(fmt.Sprintf("%#v", getParentOrdersResponse))
	if len(getParentOrdersResponse) == 0 {
		t.Errorf("no order")
	}

	parentOrderId := getParentOrdersResponse[0].ParentOrderId

	httpResponse, getChildOrdersResponse, err := apiClient.PriGetChildOrdersById("BTC_JPY", types.IdTypeParentOrderId, parentOrderId)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getChildOrdersResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v", httpResponse))
	if len(getChildOrdersResponse) == 0 {
		t.Errorf("no order")
	}

	time.Sleep(2 * time.Second)

	httpResponse, err = apiClient.PriCancelParentOrder("BTC_JPY", types.IdTypeParentOrderId, parentOrderId)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	t.Log(fmt.Sprintf("%#v", httpResponse))
}


func TestPriGetParentOrder(t *testing.T) {
	apiClient := createApiClient(t)
	sendParentOrderParameter1 := &private.SendParentOrderParameter {
		ProductCode: "BTC_JPY",
		ConditionType: types.ConditionTypeLimit,
		Side: types.SideBuy,
		Size: 0.11,
		Price: 550000,
	}
	sendParentOrderParameter2 := &private.SendParentOrderParameter {
		ProductCode: "BTC_JPY",
		ConditionType: types.ConditionTypeLimit,
		Side: types.SideSell,
		Size: 0.11,
		Price: 2000000,
	}
	httpResponse, sendParentOrderResponse, err := apiClient.PriSendParentOrder(types.OrderMethodIFD, 1, types.TimeInForceGTC, sendParentOrderParameter1, sendParentOrderParameter2)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if sendParentOrderResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v\n%#v", httpResponse, sendParentOrderResponse))

	time.Sleep(2 * time.Second)

	httpResponse, getParentOrderResponse, err := apiClient.PriGetParentOrder(types.IdTypeParentOrderAcceptanceId, sendParentOrderResponse.ParentOrderAcceptanceId)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getParentOrderResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v", httpResponse))
	t.Log(fmt.Sprintf("%#v", getParentOrderResponse))

	parentOrderId := getParentOrderResponse.ParentOrderId

	httpResponse, getParentOrderResponse, err = apiClient.PriGetParentOrder(types.IdTypeParentOrderId, parentOrderId)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	if getParentOrderResponse == nil {
		t.Errorf("no response")
	}
	t.Log(fmt.Sprintf("%#v", httpResponse))
	t.Log(fmt.Sprintf("%#v", getParentOrderResponse))

	time.Sleep(2 *time.Second)

	httpResponse, err = apiClient.PriCancelParentOrder("BTC_JPY", types.IdTypeParentOrderId, parentOrderId)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if httpResponse == nil {
		t.Errorf("no http response")
	}
	t.Log(fmt.Sprintf("%#v", httpResponse))
}


type testCallbackData struct {
	t *testing.T
	m string
}

func tickerCallback(productCode types.ProductCode, getTickerResponse *public.GetTickerResponse, callbackData interface{}) {
	tcbd := (callbackData).(*testCallbackData)
	if tcbd.m != "test" {
		 tcbd.t.Errorf("mismatch message")
	}
	log.Printf("%#v", getTickerResponse)
}

func TestRealSubscribeTicker(t *testing.T) {
	apiClient := createApiClient(t)
	tcbd := &testCallbackData{
		t: t,
		m: "test",
	}
	err := apiClient.RealTickerStart("BTC_JPY", tickerCallback, tcbd)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	time.Sleep(20 * time.Second)

	err = apiClient.RealTickerStop("BTC_JPY")
	if err != nil {
		t.Errorf("error: %v", err)
	}

}

func boardSnapshotCallback(productCode types.ProductCode, getBoardResponse *public.GetBoardResponse, callbackData interface{}) {
	tcbd := (callbackData).(*testCallbackData)
	if tcbd.m != "test" {
		 tcbd.t.Errorf("mismatch message")
	}
	log.Printf("%#v", getBoardResponse)
}

func TestRealSubscribeBoardSnapshot(t *testing.T) {
	apiClient := createApiClient(t)
	tcbd := &testCallbackData{
		t: t,
		m: "test",
	}
	err := apiClient.RealBoardSnapshotStart("BTC_JPY", boardSnapshotCallback, tcbd)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	time.Sleep(20 * time.Second)

	err = apiClient.RealBoardSnapshotStop("BTC_JPY")
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func boardCallback(productCode types.ProductCode, getBoardResponse *public.GetBoardResponse, callbackData interface{}) {
	tcbd := (callbackData).(*testCallbackData)
	if tcbd.m != "test" {
		 tcbd.t.Errorf("mismatch message")
	}
	log.Printf("%#v", getBoardResponse)
}

func TestRealSubscribeBoard(t *testing.T) {
	apiClient := createApiClient(t)
	tcbd := &testCallbackData{
		t: t,
		m: "test",
	}
	err := apiClient.RealBoardStart("BTC_JPY", boardCallback, tcbd)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	time.Sleep(20 * time.Second)

	err = apiClient.RealBoardStop("BTC_JPY")
	if err != nil {
		t.Errorf("error: %v", err)
	}
}

func executionsCallback(productCode types.ProductCode, getExecutionsResponse public.GetExecutionsResponse, callbackData interface{}) {
	tcbd := (callbackData).(*testCallbackData)
	if tcbd.m != "test" {
		 tcbd.t.Errorf("mismatch message")
	}
	log.Printf("%#v", getExecutionsResponse)
}

func TestRealSubscribeExecutions(t *testing.T) {
	apiClient := createApiClient(t)
	tcbd := &testCallbackData{
		t: t,
		m: "test",
	}
	err := apiClient.RealExecutionsStart("BTC_JPY", executionsCallback, tcbd)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	time.Sleep(20 * time.Second)

	err = apiClient.RealExecutionsStop("BTC_JPY")
	if err != nil {
		t.Errorf("error: %v", err)
	}
}
