package api

import (
	"context"
	"database/sql"

	"github.com/mentos1386/golang-rest-example/pkg/openapi"
)

type Group struct {
	ID      int64
	Name    string
	UserIds []int64
}

func (u *ApiService) getUsersForGroupId(id openapi.ID) ([]openapi.ID, error) {
	rows, err := u.db.Query("SELECT id FROM users WHERE group_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var userIds []openapi.ID
	for rows.Next() {
		var userId int64
		err = rows.Scan(&userId)
		if err != nil {
			return nil, err
		}
		userIds = append(userIds, openapi.ID(userId))
	}
	return userIds, nil
}

func (u *ApiService) GroupsGet(ctx context.Context) ([]openapi.Group, error) {
	rows, err := u.db.Query("SELECT * FROM groups")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var groups []openapi.Group

	for rows.Next() {
		var group Group
		err = rows.Scan(&group.ID, &group.Name)
		if err != nil {
			return nil, err
		}
		userIds, err := u.getUsersForGroupId(openapi.ID(group.ID))
		if err != nil {
			return nil, err
		}

		groups = append(groups, openapi.Group{
			ID:      openapi.ID(group.ID),
			Name:    group.Name,
			UserIds: userIds,
		})
	}

	return groups, nil
}

func (u *ApiService) GroupsPost(ctx context.Context, group *openapi.GroupUpdate) (*openapi.Group, error) {
	row := u.db.QueryRow("INSERT INTO groups (name) VALUES ($1) RETURNING id", group.Name)

	var id int64
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	// Just created group has no users
	var userIds []openapi.ID

	return &openapi.Group{
		ID:      openapi.ID(id),
		Name:    group.Name,
		UserIds: userIds,
	}, nil
}

func (u *ApiService) GroupsIDGet(ctx context.Context, params openapi.GroupsIDGetParams) (openapi.GroupsIDGetRes, error) {
	row := u.db.QueryRow("SELECT * FROM groups WHERE id = $1", params.ID)
	var group Group
	err := row.Scan(&group.ID, &group.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return &openapi.Error{Message: "Group not found", Code: 404}, nil
		}

		return nil, err
	}

	userIds, err := u.getUsersForGroupId(openapi.ID(group.ID))
	if err != nil {
		return nil, err
	}

	return &openapi.Group{
		ID:      openapi.ID(group.ID),
		Name:    group.Name,
		UserIds: userIds,
	}, nil
}

func (u *ApiService) GroupsIDPut(ctx context.Context, group *openapi.GroupUpdate, params openapi.GroupsIDPutParams) (openapi.GroupsIDPutRes, error) {
	res, err := u.db.Exec("UPDATE groups SET name = $1 WHERE id = $2", group.Name, params.ID)
	if err != nil {
		return nil, err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return &openapi.Error{Message: "Group not found", Code: 404}, nil
	}

	userIds, err := u.getUsersForGroupId(params.ID)
	if err != nil {
		return nil, err
	}

	return &openapi.Group{
		ID:      params.ID,
		Name:    group.Name,
		UserIds: userIds,
	}, nil
}

func (u *ApiService) GroupsIDDelete(ctx context.Context, params openapi.GroupsIDDeleteParams) (openapi.GroupsIDDeleteRes, error) {
	res, err := u.db.Exec("DELETE FROM groups WHERE id = $1", params.ID)
	if err != nil {
		return nil, err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return &openapi.Error{Message: "Group not found", Code: 404}, nil
	}

	return &openapi.Ok{Message: "OK"}, nil
}
