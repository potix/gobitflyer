package main

import (
        "log"
	"os"
        "time"
        "github.com/potix/gobitflyer/client"
        "github.com/potix/gobitflyer/api/public"
        "github.com/potix/gobitflyer/api"
)

func realtimeTickerCallback(getTickerResponse *public.GetTickerResponse, callbackData interface{}) {
        log.Printf("%#v", getTickerResponse)
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
        log.Printf("%#v --- n%#v", httpResponse, getMarketsResponse)

        httpResponse, getBoardResponse, err :=  apiClient.PubGetBoard("BTC_JPY")
        if err != nil {
                log.Printf("error: %v", err)
		os.Exit(1)
        }
        log.Printf("%#v --- %#v", httpResponse, getBoardResponse)

        httpResponse, getTickerResponse, err :=  apiClient.PubGetTicker("BTC_JPY")
        if err != nil {
                log.Printf("error: %v", err)
		os.Exit(1)
        }
        log.Printf("%#v --- %#v", httpResponse, getTickerResponse)

	// realtime api
        err = apiClient.RealTickerStart("BTC_JPY", realtimeTickerCallback, nil)
        if err != nil {
                log.Printf("error: %v", err)
		os.Exit(1)
        }

        time.Sleep(20 * time.Second)

        err = apiClient.RealTickerStop("BTC_JPY")
        if err != nil {
                log.Printf("error: %v", err)
		os.Exit(1)
        }
}

