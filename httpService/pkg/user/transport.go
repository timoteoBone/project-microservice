package user

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/timoteoBone/project-microservice/grpcService/pkg/entities"
	myerr "github.com/timoteoBone/project-microservice/httpService/pkg/errors"

	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPSrv(endpoint Endpoints) http.Handler {
	rt := mux.NewRouter()

	rt.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		endpoint.CreateUs,
		decodeCreateUserReq,
		encodeCreateUserResp,
	))

	rt.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		endpoint.GetUs,
		decodeGetUserReq,
		encodeGetUserResp,
	))

	return rt
}

func decodeCreateUserReq(ctx context.Context, r *http.Request) (interface{}, error) {

	var request entities.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return nil, myerr.ErrInvalidDataForm
	}

	return request, nil
}

func encodeCreateUserResp(ctx context.Context, wr http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(wr).Encode(response)
}

func decodeGetUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var request entities.GetUserRequest
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, myerr.ErrInvalidDataForm
	}

	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, myerr.ErrInvalidDataForm
	}

	request.UserID = num
	return request, nil
}

func encodeGetUserResp(ctx context.Context, wr http.ResponseWriter, response interface{}) error {

	return json.NewEncoder(wr).Encode(response)
}
