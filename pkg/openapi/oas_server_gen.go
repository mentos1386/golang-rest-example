// Code generated by ogen, DO NOT EDIT.

package openapi

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// GroupsGet implements GET /groups operation.
	//
	// Get all groups.
	//
	// GET /groups
	GroupsGet(ctx context.Context) ([]Group, error)
	// GroupsIDDelete implements DELETE /groups/{id} operation.
	//
	// Delete group by id.
	//
	// DELETE /groups/{id}
	GroupsIDDelete(ctx context.Context, params GroupsIDDeleteParams) (GroupsIDDeleteRes, error)
	// GroupsIDGet implements GET /groups/{id} operation.
	//
	// Get group by id.
	//
	// GET /groups/{id}
	GroupsIDGet(ctx context.Context, params GroupsIDGetParams) (GroupsIDGetRes, error)
	// GroupsIDPut implements PUT /groups/{id} operation.
	//
	// Update group by id.
	//
	// PUT /groups/{id}
	GroupsIDPut(ctx context.Context, req *GroupUpdate, params GroupsIDPutParams) (GroupsIDPutRes, error)
	// GroupsPost implements POST /groups operation.
	//
	// Create group.
	//
	// POST /groups
	GroupsPost(ctx context.Context, req *GroupUpdate) (*Group, error)
	// HealthzGet implements GET /healthz operation.
	//
	// GET /healthz
	HealthzGet(ctx context.Context) (*Ok, error)
	// UsersGet implements GET /users operation.
	//
	// Get all users.
	//
	// GET /users
	UsersGet(ctx context.Context) ([]User, error)
	// UsersIDDelete implements DELETE /users/{id} operation.
	//
	// Delete user by id.
	//
	// DELETE /users/{id}
	UsersIDDelete(ctx context.Context, params UsersIDDeleteParams) (UsersIDDeleteRes, error)
	// UsersIDGet implements GET /users/{id} operation.
	//
	// Get user by id.
	//
	// GET /users/{id}
	UsersIDGet(ctx context.Context, params UsersIDGetParams) (UsersIDGetRes, error)
	// UsersIDPut implements PUT /users/{id} operation.
	//
	// Update user by id.
	//
	// PUT /users/{id}
	UsersIDPut(ctx context.Context, req *UserUpdate, params UsersIDPutParams) (UsersIDPutRes, error)
	// UsersPost implements POST /users operation.
	//
	// Create user.
	//
	// POST /users
	UsersPost(ctx context.Context, req *UserUpdate) (*User, error)
	// NewError creates *ErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
