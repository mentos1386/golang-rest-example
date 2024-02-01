// Code generated by ogen, DO NOT EDIT.

package openapi

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}
	args := [1]string{}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			origElem := elem
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'g': // Prefix: "groups"
				origElem := elem
				if l := len("groups"); len(elem) >= l && elem[0:l] == "groups" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleGroupsGetRequest([0]string{}, elemIsEscaped, w, r)
					case "POST":
						s.handleGroupsPostRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					origElem := elem
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "DELETE":
							s.handleGroupsIDDeleteRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "GET":
							s.handleGroupsIDGetRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "PUT":
							s.handleGroupsIDPutRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET,PUT")
						}

						return
					}

					elem = origElem
				}

				elem = origElem
			case 'h': // Prefix: "healthz"
				origElem := elem
				if l := len("healthz"); len(elem) >= l && elem[0:l] == "healthz" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleHealthzGetRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}

				elem = origElem
			case 'u': // Prefix: "users"
				origElem := elem
				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "GET":
						s.handleUsersGetRequest([0]string{}, elemIsEscaped, w, r)
					case "POST":
						s.handleUsersPostRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET,POST")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					origElem := elem
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch r.Method {
						case "DELETE":
							s.handleUsersIDDeleteRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "GET":
							s.handleUsersIDGetRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						case "PUT":
							s.handleUsersIDPutRequest([1]string{
								args[0],
							}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "DELETE,GET,PUT")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/group"
						origElem := elem
						if l := len("/group"); len(elem) >= l && elem[0:l] == "/group" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "PUT":
								s.handleUsersIDGroupPutRequest([1]string{
									args[0],
								}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "PUT")
							}

							return
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [1]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			origElem := elem
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'g': // Prefix: "groups"
				origElem := elem
				if l := len("groups"); len(elem) >= l && elem[0:l] == "groups" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "GroupsGet"
						r.summary = ""
						r.operationID = ""
						r.pathPattern = "/groups"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = "GroupsPost"
						r.summary = ""
						r.operationID = ""
						r.pathPattern = "/groups"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					origElem := elem
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Leaf parameter
					args[0] = elem
					elem = ""

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							// Leaf: GroupsIDDelete
							r.name = "GroupsIDDelete"
							r.summary = ""
							r.operationID = ""
							r.pathPattern = "/groups/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							// Leaf: GroupsIDGet
							r.name = "GroupsIDGet"
							r.summary = ""
							r.operationID = ""
							r.pathPattern = "/groups/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "PUT":
							// Leaf: GroupsIDPut
							r.name = "GroupsIDPut"
							r.summary = ""
							r.operationID = ""
							r.pathPattern = "/groups/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}

					elem = origElem
				}

				elem = origElem
			case 'h': // Prefix: "healthz"
				origElem := elem
				if l := len("healthz"); len(elem) >= l && elem[0:l] == "healthz" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						// Leaf: HealthzGet
						r.name = "HealthzGet"
						r.summary = ""
						r.operationID = ""
						r.pathPattern = "/healthz"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}

				elem = origElem
			case 'u': // Prefix: "users"
				origElem := elem
				if l := len("users"); len(elem) >= l && elem[0:l] == "users" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "GET":
						r.name = "UsersGet"
						r.summary = ""
						r.operationID = ""
						r.pathPattern = "/users"
						r.args = args
						r.count = 0
						return r, true
					case "POST":
						r.name = "UsersPost"
						r.summary = ""
						r.operationID = ""
						r.pathPattern = "/users"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					origElem := elem
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					// Param: "id"
					// Match until "/"
					idx := strings.IndexByte(elem, '/')
					if idx < 0 {
						idx = len(elem)
					}
					args[0] = elem[:idx]
					elem = elem[idx:]

					if len(elem) == 0 {
						switch method {
						case "DELETE":
							r.name = "UsersIDDelete"
							r.summary = ""
							r.operationID = ""
							r.pathPattern = "/users/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "GET":
							r.name = "UsersIDGet"
							r.summary = ""
							r.operationID = ""
							r.pathPattern = "/users/{id}"
							r.args = args
							r.count = 1
							return r, true
						case "PUT":
							r.name = "UsersIDPut"
							r.summary = ""
							r.operationID = ""
							r.pathPattern = "/users/{id}"
							r.args = args
							r.count = 1
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/group"
						origElem := elem
						if l := len("/group"); len(elem) >= l && elem[0:l] == "/group" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "PUT":
								// Leaf: UsersIDGroupPut
								r.name = "UsersIDGroupPut"
								r.summary = ""
								r.operationID = ""
								r.pathPattern = "/users/{id}/group"
								r.args = args
								r.count = 1
								return r, true
							default:
								return
							}
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	return r, false
}