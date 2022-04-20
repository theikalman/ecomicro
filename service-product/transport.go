package product

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
	"github.com/gorilla/mux"
)

var (
	errInvalidProductID   = errors.New("invalid product id")
	errInvalidRequestBody = errors.New("invalid request body")
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

	getProductByIDHandler := kithttp.NewServer(
		makeGetProductByIDEndpoint(s),
		decodeGetProductByIDRequestQuery,
		encodeResponse,
		opts...,
	)

	getProductsHandler := kithttp.NewServer(
		makeGetProductsEndpoint(s),
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
	r.Handle("/product/v1/product/{productID}", getProductByIDHandler).Methods("GET")
	r.Handle("/product/v1/product", getProductsHandler).Methods("GET") // TODO Pagination
	r.Handle("/product/v1/product", createProductHandler).Methods("POST")

	return r
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case errInvalidProductID, errInvalidRequestBody:
		w.WriteHeader(http.StatusBadRequest)
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

type getProductByIDQuery struct {
	productID uint
}

func decodeGetProductByIDRequestQuery(_ context.Context, r *http.Request) (interface{}, error) {
	queries := mux.Vars(r)
	productIDQuery, ok := queries["productID"]
	if !ok {
		return nil, errInvalidProductID
	}

	productID, err := strconv.ParseUint(productIDQuery, 10, 64)
	if err != nil {
		log.Printf("unable to parse to uint: %s", err) // TODO Make logger to be injectable
		return nil, errInvalidProductID
	}

	return getProductByIDQuery{
		productID: uint(productID),
	}, nil
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
