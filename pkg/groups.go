package groups

import (
	"context"
	"github.com/mentos1386/golang-rest-example/pkg/api"
)

type groupServer struct {
	groups map[int64]*api.Group
}

func (u *groupServer) groupsGet(ctx context.Context) ([]api.Group, error) {
	var groups []api.Group
	for _, group := range u.groups {
		groups = append(groups, *group)
	}
	return groups, nil
}
