package main

import (
	"coins/ranking-service/client"
	"coins/ranking-service/router"
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
	var rankListClient client.RankListClientI
	rankListClient = client.NewRankListClient(c, "https://min-api.cryptocompare.com/data/top/mktcapfull?tsym=USD")
	router, err := router.NewRouter(&rankListClient)
	if err != nil {
		log.Fatalf("error initializing router: %v", err)
	}

	fmt.Println("Service is ready...")
	log.Fatal(http.ListenAndServe(":1904", router))

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
}
