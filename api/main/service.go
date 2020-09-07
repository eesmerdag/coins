package main

import (
	"coins/api/client"
	"coins/api/router"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := &http.Client{
		Timeout: time.Second * 10,
	}

	var priceSvcClient client.PriceSvcClientI
	priceSvcClient = client.NewPriceListClient(c, "http://127.0.0.1:1903/usd-prices")

	var rankSvcClient client.RankSvcClientI
	rankSvcClient = client.NewRankSvcClient(c, "http://127.0.0.1:1904/coin-ranks")
	router, err := router.NewRouter(&priceSvcClient, &rankSvcClient)
	if err != nil {
		log.Fatalf("error initializing router: %v", err)
	}

	fmt.Println("Service is ready...")
	log.Fatal(http.ListenAndServe(":1902", router))

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
}
