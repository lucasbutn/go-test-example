package main

import (
	"log"
	"net/http"
	"test-example/balance"
)

func main() {
	rc, err := balance.NewRestClient("http://anyUrl.com/movements")
	if err != nil {
		panic(err)
	}

	retrier := balance.NewDelayedRetrier(2, 500)

	srv, err := balance.NewService(rc, retrier)
	if err != nil {
		panic(err)
	}

	ctrl, err := balance.NewController(srv)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/balances", ctrl.GetBalance)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
