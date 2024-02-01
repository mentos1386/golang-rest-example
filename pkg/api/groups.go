package api

import (
	"context"
	"github.com/mentos1386/golang-rest-example/pkg/openapi"
)

func (u *ApiService) GroupsGet(ctx context.Context) ([]openapi.Group, error) {
	var groups []openapi.Group
	for _, group := range u.groups {
		groups = append(groups, *group)
	}
	groups = append(groups, openapi.Group{ID: openapi.ID(1), Name: "Admins"})
	return groups, nil
}
