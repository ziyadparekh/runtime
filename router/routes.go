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
	ROUTE_INDEX   = "/"
	ROUTE_PAYLOAD = "/payload"
)

const (
	NAME_INDEX   = "Index"
	NAME_PAYLOAD = "Payload"
)

var routes = Routes{
	Route{
		NAME_INDEX,
		METHOD_GET,
		ROUTE_INDEX,
		handlers.Index,
	},
	Route{
		NAME_PAYLOAD,
		METHOD_POST,
		ROUTE_PAYLOAD,
		handlers.Payload,
	},
}
