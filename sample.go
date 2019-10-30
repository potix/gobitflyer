package main

import (
        "log"
	"os"
        "time"
        "github.com/potix/gobitflyer/client"
        "github.com/potix/gobitflyer/api/types"
        "github.com/potix/gobitflyer/api/public"
        "github.com/potix/gobitflyer/api"
)

func realtimeTickerCallback(productCode types.ProductCode, getTickerResponse *public.GetTickerResponse, callbackData interface{}) {
        log.Printf("===== ticker =====")
        log.Printf("timestamp %v", getTickerResponse.Timestamp)
        log.Printf("Best Bid Price %v Size %v", getTickerResponse.BestBid, getTickerResponse.BestBidSize)
        log.Printf("Best Ask Price %v Size %v", getTickerResponse.BestAsk, getTickerResponse.BestAskSize)
        log.Printf("Total Bid %v", getTickerResponse.TotalBidDepth)
        log.Printf("Total Ask %v", getTickerResponse.TotalAskDepth)
        log.Printf("Volume %v Volume By Product %v", getTickerResponse.Volume, getTickerResponse.VolumeByProduct)
}

func realtimeBoardCallback(productCode types.ProductCode, getBoardResponse *public.GetBoardResponse, callbackData interface{}) {
        log.Printf("===== board=====")
        log.Printf("Mid Price %v", getBoardResponse.MidPrice)
        log.Printf("--- asks ---")
	for i := 0; i < 6; i+= 1 {
		log.Printf("%#v", getBoardResponse.Asks[i])
	}
        log.Printf("--- bids ---")
	for i := 0; i < 6; i+= 1 {
		log.Printf("%#v", getBoardResponse.Bids[i])
	}
}

func main() {
        httpClient := client.NewHTTPClient(0, 0, 0, nil)
        authenticator, err := api.NewAuthenticator("sampleApiKeyFile")
        if err != nil {
                log.Printf("can not create authenticator: %v", err)
		os.Exit(1)
        }
        apiClient :=  api.NewAPIClient(httpClient, authenticator)

	// public api
        httpResponse, getMarketsResponse, err :=  apiClient.PubGetMarkets()
        if err != nil {
                log.Printf("error: %v", err)
		os.Exit(1)
        }
        log.Printf("===== markkets =====")
        log.Printf("status %v", httpResponse.Status)
	for _, market := range getMarketsResponse {
		log.Printf("%v", market.ProductCode)
	}

        httpResponse, getBoardResponse, err :=  apiClient.PubGetBoard("BTC_JPY")
        if err != nil {
                log.Printf("error: %v", err)
		os.Exit(1)
        }
        log.Printf("===== board=====")
        log.Printf("status %v", httpResponse.Status)
        log.Printf("Mid Price %v", getBoardResponse.MidPrice)
        log.Printf("--- asks ---")
	for i := 0; i < 6; i+= 1 {
		log.Printf("%#v", getBoardResponse.Asks[i])
	}
        log.Printf("--- bids ---")
	for i := 0; i < 6; i+= 1 {
		log.Printf("%#v", getBoardResponse.Bids[i])
	}

        httpResponse, getTickerResponse, err :=  apiClient.PubGetTicker("BTC_JPY")
        if err != nil {
                log.Printf("error: %v", err)
		os.Exit(1)
        }
        log.Printf("===== ticker =====")
        log.Printf("status %v", httpResponse.Status)
        log.Printf("timestamp %v", getTickerResponse.Timestamp)
        log.Printf("Best Bid Price %v Size %v", getTickerResponse.BestBid, getTickerResponse.BestBidSize)
        log.Printf("Best Ask Price %v Size %v", getTickerResponse.BestAsk, getTickerResponse.BestAskSize)
        log.Printf("Total Bid %v", getTickerResponse.TotalBidDepth)
        log.Printf("Total Ask %v", getTickerResponse.TotalAskDepth)
        log.Printf("Volume %v Volume By Product %v", getTickerResponse.Volume, getTickerResponse.VolumeByProduct)

	// realtime api
	wsClient := client.NewWSClient(0, 0, 3, 1, nil)
        realApiClient1 := api.NewRealAPIClient(wsClient)
        err = realApiClient1.RealTickerStart("BTC_JPY", realtimeTickerCallback, nil)
        if err != nil {
                log.Printf("error: %v", err)
		os.Exit(1)
        }
	wsClient = client.NewWSClient(0, 0, 3, 1, nil)
        realApiClient2 := api.NewRealAPIClient(wsClient)
        err = realApiClient2.RealBoardStart("BTC_JPY", realtimeBoardCallback, nil, true)
        if err != nil {
                log.Printf("error: %v", err)
		os.Exit(1)
        }

        time.Sleep(20 * time.Second)

        err = realApiClient1.RealStop()
        if err != nil {
                log.Printf("error: %v", err)
		os.Exit(1)
        }
        err = realApiClient2.RealStop()
        if err != nil {
                log.Printf("error: %v", err)
		os.Exit(1)
        }
}

