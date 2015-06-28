package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"HostIndex",
		"GET",
		"/api/hosts",
		HostIndex,
	},
	Route{
		"HostCreate",
		"POST",
		"/api/hosts",
		HostCreate,
	},
	Route{
		"HostDetail",
		"GET",
		"/api/hosts/{hostId}",
		HostDetail,
	},
	Route{
		"PublicCatchAll",
		"GET",
		`/{[a-zA-Z0-9=\-\/]+}`,
		Index,
	},
	Route{
		"Index",
		"GET",
		`/`,
		Index,
	},
}
