package api

import (
	"context"
	"database/sql"

	"github.com/mentos1386/golang-rest-example/pkg/openapi"
)

type User struct {
	ID      int64
	Name    string
	Email   string
	GroupID int64
}

func (u *ApiService) UsersGet(ctx context.Context) ([]openapi.User, error) {
	rows, err := u.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []openapi.User

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
		users = append(users, openapi.User{
			ID:   openapi.ID(user.ID),
			Name: user.Name,
		})
	}

	return users, nil
}

func (u *ApiService) UsersPost(ctx context.Context, user *openapi.UserUpdate) (*openapi.User, error) {
	row := u.db.QueryRow("INSERT INTO users (name, email, group_id) VALUES ($1, $2, $3) RETURNING id", user.Name, user.Email, user.GroupID)

	var id int64
	err := row.Scan(&id)
	if err != nil {
		return nil, err
	}

	return &openapi.User{
		ID:      openapi.ID(id),
		Name:    user.Name,
		Email:   user.Email,
		GroupID: user.GroupID,
	}, nil
}

func (u *ApiService) UsersIDGet(ctx context.Context, params openapi.UsersIDGetParams) (openapi.UsersIDGetRes, error) {
	row := u.db.QueryRow("SELECT * FROM users WHERE id = $1", params.ID)
	var user User
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return &openapi.Error{Message: "User not found", Code: 404}, nil
		}

		return nil, err
	}
	return &openapi.User{
		ID:      openapi.ID(user.ID),
		Name:    user.Name,
		Email:   user.Email,
		GroupID: openapi.ID(user.GroupID),
	}, nil
}

func (u *ApiService) UsersIDPut(ctx context.Context, user *openapi.UserUpdate, params openapi.UsersIDPutParams) (openapi.UsersIDPutRes, error) {
	res, err := u.db.Exec("UPDATE users SET name = $1, email = $2, group_id = $3 WHERE id = $4", user.Name, user.Email, user.GroupID, params.ID)
	if err != nil {
		return nil, err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return &openapi.Error{Message: "User not found", Code: 404}, nil
	}

	return &openapi.User{
		ID:      params.ID,
		Name:    user.Name,
		Email:   user.Email,
		GroupID: openapi.ID(user.GroupID),
	}, nil
}

func (u *ApiService) UsersIDDelete(ctx context.Context, params openapi.UsersIDDeleteParams) (openapi.UsersIDDeleteRes, error) {
	res, err := u.db.Exec("DELETE FROM users WHERE id = $1", params.ID)
	if err != nil {
		return nil, err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return &openapi.Error{Message: "User not found", Code: 404}, nil
	}

	return &openapi.Ok{Message: "OK"}, nil
}
