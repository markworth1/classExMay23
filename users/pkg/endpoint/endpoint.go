package endpoint

import (
	"context"

	service "github.com/gSchool/classExMay23/users/pkg/service"
	endpoint "github.com/go-kit/kit/endpoint"
)

// CreateRequest collects the request parameters for the Create method.
type CreateRequest struct {
	User service.User `json:"user"`
}

// CreateResponse collects the response parameters for the Create method.
type CreateResponse struct {
	E0 error `json:"e0"`
}

// MakeCreateEndpoint returns an endpoint that invokes Create on the service.
func MakeCreateEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		e0 := s.Create(ctx, req.User)
		return CreateResponse{E0: e0}, nil
	}
}

// Failed implements Failer.
func (r CreateResponse) Failed() error {
	return r.E0
}

// Failure is an interface that should be implemented by response types.
// Response encoders can check if responses are Failer, and if so they've
// failed, and if so encode them using a separate write path based on the error.
type Failure interface {
	Failed() error
}

// Create implements Service. Primarily useful in a client.
func (e Endpoints) Create(ctx context.Context, user service.User) (e0 error) {
	request := CreateRequest{User: user}
	response, err := e.CreateEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(CreateResponse).E0
}

// GetAllRequest collects the request parameters for the GetAll method.
type GetAllRequest struct{}

// GetAllResponse collects the response parameters for the GetAll method.
type GetAllResponse struct {
	S0 []string `json:"s0"`
	E1 error    `json:"e1"`
}

// MakeGetAllEndpoint returns an endpoint that invokes GetAll on the service.
func MakeGetAllEndpoint(s service.UsersService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		s0, e1 := s.GetAll(ctx)
		return GetAllResponse{
			E1: e1,
			S0: s0,
		}, nil
	}
}

// Failed implements Failer.
func (r GetAllResponse) Failed() error {
	return r.E1
}

// GetAll implements Service. Primarily useful in a client.
func (e Endpoints) GetAll(ctx context.Context) (s0 []string, e1 error) {
	request := GetAllRequest{}
	response, err := e.GetAllEndpoint(ctx, request)
	if err != nil {
		return
	}
	return response.(GetAllResponse).S0, response.(GetAllResponse).E1
}
