package endpoint

import (
	"context"

	model "github.com/go-kit/kit/bankaccount/pkg/model"
	service "github.com/go-kit/kit/bankaccount/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateAccRequest collects the request parameters for the CreateAcc method.
type CreateAccRequest struct {
	Acc model.Account `json:"acc"`
}

// CreateAccResponse collects the response parameters for the CreateAcc method.
type CreateAccResponse struct {
	S0 string `json:"message"`
	E1 error  `json:"e"`
}

// MakeCreateAccEndpoint returns an endpoint that invokes CreateAcc on the service.
func MakeCreateAccEndpoint(s service.BankaccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateAccRequest)
		s0, e1 := s.CreateAcc(ctx, req.Acc)
		return CreateAccResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r CreateAccResponse) Failed() error {
	return r.E1
}

// GetAccRequest collects the request parameters for the GetAcc method.
type GetAccRequest struct {
	Uuid string `json:"uuid"`
}

// GetAccResponse collects the response parameters for the GetAcc method.
type GetAccResponse struct {
	M0 model.Account `json:"data"`
	E1 error         `json:"e"`
}

// MakeGetAccEndpoint returns an endpoint that invokes GetAcc on the service.
func MakeGetAccEndpoint(s service.BankaccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAccRequest)
		m0, e1 := s.GetAcc(ctx, req.Uuid)
		return GetAccResponse{
			E1: e1,
			M0: m0,
		}, nil
	}
}

// Failed implements Failer.
func (r GetAccResponse) Failed() error {
	return r.E1
}

// DepositAmountRequest collects the request parameters for the DepositAmount method.
type DepositAmountRequest struct {
	Number  int64   `json:"number"`
	Ammount float64 `json:"ammount"`
}

// DepositAmountResponse collects the response parameters for the DepositAmount method.
type DepositAmountResponse struct {
	S0 string `json:"message"`
	E1 error  `json:"e"`
}

// MakeDepositAmountEndpoint returns an endpoint that invokes DepositAmount on the service.
func MakeDepositAmountEndpoint(s service.BankaccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DepositAmountRequest)
		s0, e1 := s.DepositAmount(ctx, req.Number, req.Ammount)
		return DepositAmountResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r DepositAmountResponse) Failed() error {
	return r.E1
}

// WithDrawAmountRequest collects the request parameters for the WithDrawAmount method.
type WithDrawAmountRequest struct {
	Number  int64   `json:"number"`
	Ammount float64 `json:"ammount"`
}

// WithDrawAmountResponse collects the response parameters for the WithDrawAmount method.
type WithDrawAmountResponse struct {
	S0 string `json:"message"`
	E1 error  `json:"e"`
}

// MakeWithDrawAmountEndpoint returns an endpoint that invokes WithDrawAmount on the service.
func MakeWithDrawAmountEndpoint(s service.BankaccountService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(WithDrawAmountRequest)
		s0, e1 := s.WithDrawAmount(ctx, req.Number, req.Ammount)
		return WithDrawAmountResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r WithDrawAmountResponse) Failed() error {
	return r.E1
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// CreateAcc implements Service. Primarily useful in a client.
func (e Endpoints) CreateAcc(ctx context.Context, acc model.Account) (s0 string, e1 error) {
	request := CreateAccRequest{Acc: acc}
	response, err := e.CreateAccEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateAccResponse).S0, response.(CreateAccResponse).E1
}

// GetAcc implements Service. Primarily useful in a client.
func (e Endpoints) GetAcc(ctx context.Context, uuid string) (m0 model.Account, e1 error) {
	request := GetAccRequest{Uuid: uuid}
	response, err := e.GetAccEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAccResponse).M0, response.(GetAccResponse).E1
}

// DepositAmount implements Service. Primarily useful in a client.
func (e Endpoints) DepositAmount(ctx context.Context, number int64, ammount float64) (s0 string, e1 error) {
	request := DepositAmountRequest{
		Ammount: ammount,
		Number:  number,
	}
	response, err := e.DepositAmountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(DepositAmountResponse).S0, response.(DepositAmountResponse).E1
}

// WithDrawAmount implements Service. Primarily useful in a client.
func (e Endpoints) WithDrawAmount(ctx context.Context, number int64, ammount float64) (s0 string, e1 error) {
	request := WithDrawAmountRequest{
		Ammount: ammount,
		Number:  number,
	}
	response, err := e.WithDrawAmountEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(WithDrawAmountResponse).S0, response.(WithDrawAmountResponse).E1
}
