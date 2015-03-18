package router

import (
	"net/http"

	"github.com/ziyadparekh/runtime/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

const (
	METHOD_GET    = "GET"
	METHOD_PUT    = "PUT"
	METHOD_POST   = "POST"
	METHOD_DELETE = "DELETE"
)

const (
	ROUTE_INDEX  = "/"
	ROUTE_PULL   = "/hooks/pullrequest"
	ROUTE_BRANCH = "/hooks/branch"
)

const (
	NAME_INDEX  = "Index"
	NAME_PULL   = "Pull Request"
	NAME_BRANCH = "Branch"
)

var routes = Routes{
	Route{
		NAME_INDEX,
		METHOD_GET,
		ROUTE_INDEX,
		handlers.Index,
	},
	Route{
		NAME_PULL,
		METHOD_POST,
		ROUTE_PULL,
		handlers.PullRequestHook,
	},
	Route{
		NAME_BRANCH,
		METHOD_POST,
		ROUTE_BRANCH,
		handlers.BranchHook,
	},
}
