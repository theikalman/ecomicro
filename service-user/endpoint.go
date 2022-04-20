package user

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"golang.org/x/crypto/bcrypt"
)

func makeGetVersionEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return s.Version(), nil
	}
}

func makeSignupEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		formRequest, ok := request.(signupFormRequest)
		if !ok {
			return nil, errInvalidRequestBody
		}

		hashedPass, err := bcrypt.GenerateFromPassword([]byte(formRequest.Password), 14)
		if err != nil {
			return nil, err
		}

		u := User{
			Username: formRequest.Username,
			Password: string(hashedPass),
		}

		return s.Signup(u)
	}
}
