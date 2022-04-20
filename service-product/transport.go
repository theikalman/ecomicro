package product

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/examples/shipping/cargo"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
	"github.com/gorilla/mux"
)

// MakeHandler returns a handler for the booking service.
func MakeHandler(s Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	getVersionHandler := kithttp.NewServer(
		makeGetVersionEndpoint(s),
		decodeEmptyRequestBody,
		encodeResponse,
		opts...,
	)

	createProductHandler := kithttp.NewServer(
		makeCreateProductEndpoint(s),
		decodeCreateProductRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/product/v1/version", getVersionHandler).Methods("GET")
	r.Handle("/product/v1/product", createProductHandler).Methods("POST")

	return r
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case cargo.ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func decodeCreateProductRequest(_ context.Context, request *http.Request) (interface{}, error) {
	var productRequest struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}

	if err := json.NewDecoder(request.Body).Decode(&productRequest); err != nil {
		return nil, err
	}

	product := Product{
		Name:        productRequest.Name,
		Description: productRequest.Description,
	}

	return product, nil
}

func decodeEmptyRequestBody(_ context.Context, _ *http.Request) (interface{}, error) {
	return emptyBody, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}