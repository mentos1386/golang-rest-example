package api

import (
	"github.com/mentos1386/golang-rest-example/pkg/openapi"
)

type ApiService struct {
	groups []*openapi.Group

	openapi.UnimplementedHandler
}
