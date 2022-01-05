package http

import (
	"context"
	"net/http"

	endpoint "github.com/go-kit/kit/bankaccount/pkg/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// NewHTTPServer is a good little server
func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(accessControl) // @see https://stackoverflow.com/a/51456342

	r.Methods("GET").Path("/api/accounts/get").Handler(httptransport.NewServer(
		endpoints.GetAccEndpoint,
		decodeGetAccRequest,
		encodeGetAccResponse,
	))

	r.Methods("POST").Path("/api/accounts/create").Handler(httptransport.NewServer(
		endpoints.CreateAccEndpoint,
		decodeCreateAccRequest,
		encodeCreateAccResponse,
	))

	r.Methods("POST").Path("/api/accounts/deposit").Handler(httptransport.NewServer(
		endpoints.DepositAmountEndpoint,
		decodeDepositAmountRequest,
		encodeDepositAmountResponse,
	))

	r.Methods("POST").Path("/api/accounts/withdraw").Handler(httptransport.NewServer(
		endpoints.WithDrawAmountEndpoint,
		decodeWithDrawAmountRequest,
		encodeWithDrawAmountResponse,
	))
	return r
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		w.Header().Add("Content-Type", "application/json")
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
