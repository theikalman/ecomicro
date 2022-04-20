package user

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/examples/shipping/cargo"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
	"github.com/gorilla/mux"
)

var (
	errUsernameRequired = errors.New("username required")
	errPasswordRequired = errors.New("password required")

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

	signupHandler := kithttp.NewServer(
		makeSignupEndpoint(s),
		decodeSignupRequestForm,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/user/v1/version", getVersionHandler).Methods("GET")
	r.Handle("/user/v1/signup", signupHandler).Methods("POST")

	return r
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case cargo.ErrUnknown:
		w.WriteHeader(http.StatusNotFound)
	case errUsernameRequired, errPasswordRequired, errInvalidRequestBody:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func decodeEmptyRequestBody(_ context.Context, _ *http.Request) (interface{}, error) {
	return emptyBody, nil
}

type signupFormRequest struct {
	Username string
	Password string
}

func (s signupFormRequest) Validate() error {
	if s.Username == "" {
		return errUsernameRequired
	}

	if s.Password == "" {
		return errPasswordRequired
	}

	return nil
}

func decodeSignupRequestForm(_ context.Context, r *http.Request) (interface{}, error) {
	signupForm := signupFormRequest{
		Username: r.FormValue("username"),
		Password: r.FormValue("password"),
	}

	if err := signupForm.Validate(); err != nil {
		return nil, err
	}

	return signupForm, nil
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
