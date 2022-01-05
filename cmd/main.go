package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/bankaccount/pkg/endpoint"
	htppclient "github.com/go-kit/kit/bankaccount/pkg/http"
	"github.com/go-kit/kit/bankaccount/pkg/service"
)

func main() {
	service.InitEnv()

	var (
		// httpAddr = flag.String("http", ":6969", "http listen address")
		httpAddr = flag.String("http", ":"+service.ENV_CONFIG.Port, "http listen address")
	)

	flag.Parse()
	ctx := context.Background()
	// our napodate service
	srv := service.NewBasicBankaccountService()

	// mapping endpoints
	endpoints := endpoint.Endpoints{
		GetAccEndpoint:         endpoint.MakeGetAccEndpoint(srv),
		CreateAccEndpoint:      endpoint.MakeCreateAccEndpoint(srv),
		DepositAmountEndpoint:  endpoint.MakeDepositAmountEndpoint(srv),
		WithDrawAmountEndpoint: endpoint.MakeWithDrawAmountEndpoint(srv),
	}

	errChan := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// HTTP transport
	go func() {
		// log.Println("napodate is listening on port:", *httpAddr)
		handler := htppclient.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	go func() {
		service.CreateSession()
	}()

	log.Fatalln(<-errChan)
}
