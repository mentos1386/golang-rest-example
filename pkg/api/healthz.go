package api

import (
	"context"
	"github.com/mentos1386/golang-rest-example/pkg/openapi"
)

func (u *ApiService) HealthzGet(ctx context.Context) (*openapi.Ok, error) {
	err := u.db.Ping()
	if err != nil {
		return nil, err
	}

	return &openapi.Ok{Message: "OK"}, nil
}
