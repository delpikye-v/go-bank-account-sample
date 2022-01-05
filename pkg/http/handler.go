package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	endpoint "github.com/go-kit/kit/bankaccount/pkg/endpoint"
	http1 "github.com/go-kit/kit/transport/http"
)

// makeCreateAccHandler creates the handler logic
func makeCreateAccHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/create-acc", http1.NewServer(endpoints.CreateAccEndpoint, decodeCreateAccRequest, encodeCreateAccResponse, options...))
}

// decodeCreateAccRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeCreateAccRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.CreateAccRequest{}
	err := json.NewDecoder(r.Body).Decode(&req.Acc)
	return req, err
}

// encodeCreateAccResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeCreateAccResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeGetAccHandler creates the handler logic
func makeGetAccHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/get-acc", http1.NewServer(endpoints.GetAccEndpoint, decodeGetAccRequest, encodeGetAccResponse, options...))
}

// decodeGetAccRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeGetAccRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.GetAccRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeGetAccResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeGetAccResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeDepositAmountHandler creates the handler logic
func makeDepositAmountHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/deposit-amount", http1.NewServer(endpoints.DepositAmountEndpoint, decodeDepositAmountRequest, encodeDepositAmountResponse, options...))
}

// decodeDepositAmountRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeDepositAmountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.DepositAmountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeDepositAmountResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeDepositAmountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}

// makeWithDrawAmountHandler creates the handler logic
func makeWithDrawAmountHandler(m *http.ServeMux, endpoints endpoint.Endpoints, options []http1.ServerOption) {
	m.Handle("/with-draw-amount", http1.NewServer(endpoints.WithDrawAmountEndpoint, decodeWithDrawAmountRequest, encodeWithDrawAmountResponse, options...))
}

// decodeWithDrawAmountRequest is a transport/http.DecodeRequestFunc that decodes a
// JSON-encoded request from the HTTP request body.
func decodeWithDrawAmountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := endpoint.WithDrawAmountRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// encodeWithDrawAmountResponse is a transport/http.EncodeResponseFunc that encodes
// the response as JSON to the response writer
func encodeWithDrawAmountResponse(ctx context.Context, w http.ResponseWriter, response interface{}) (err error) {
	if f, ok := response.(endpoint.Failure); ok && f.Failed() != nil {
		ErrorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(w).Encode(response)
	return
}
func ErrorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.WriteHeader(err2code(err))
	json.NewEncoder(w).Encode(errorWrapper{Error: err.Error()})
}
func ErrorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return err
	}
	return errors.New(w.Error)
}

// This is used to set the http status, see an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/pkg/addtransport/http.go#L133
func err2code(err error) int {
	return http.StatusInternalServerError
}

type errorWrapper struct {
	Error string `json:"error"`
}
