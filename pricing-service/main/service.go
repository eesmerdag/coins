package main

import (
	"coins/pricing-service/client"
	"coins/pricing-service/router"
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
	var priceListClient client.PriceListClientI
	priceListClient = client.NewPriceListClient(c, "https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest")
	router, err := router.NewRouter(&priceListClient)
	if err != nil {
		log.Fatalf("error initializing router: %v", err)
	}

	fmt.Println("Service is ready...")
	log.Fatal(http.ListenAndServe(":1903", router))

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
}
