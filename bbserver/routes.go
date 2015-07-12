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
		"LoginAccount",
		"POST",
		"/api/account/login",
		LoginAccount,
	},
	Route{
		"PublicCatchAll6", //TODO: Get rid of this crap by learning some damn RegEx.  This was done without Internet :(
		"GET",
		`/{[a-zA-Z0-9=\-\/]+}/{[a-zA-Z0-9=\-\/]+}/{[a-zA-Z0-9=\-\/]+}/{[a-zA-Z0-9=\-\/]+}/{[a-zA-Z0-9=\-\/]+}/{[a-zA-Z0-9=\-\/]+}`,
		Index,
	},
	Route{
		"PublicCatchAll5",
		"GET",
		`/{[a-zA-Z0-9=\-\/]+}/{[a-zA-Z0-9=\-\/]+}/{[a-zA-Z0-9=\-\/]+}/{[a-zA-Z0-9=\-\/]+}/{[a-zA-Z0-9=\-\/]+}`,
		Index,
	},
	Route{
		"PublicCatchAll4",
		"GET",
		`/{[a-zA-Z0-9=\-\/]+}/{[a-zA-Z0-9=\-\/]+}/{[a-zA-Z0-9=\-\/]+}/{[a-zA-Z0-9=\-\/]+}`,
		Index,
	},
	Route{
		"PublicCatchAll3",
		"GET",
		`/{[a-zA-Z0-9=\-\/]+}/{[a-zA-Z0-9=\-\/]+}/{[a-zA-Z0-9=\-\/]+}`,
		Index,
	},
	Route{
		"PublicCatchAll2",
		"GET",
		`/{[a-zA-Z0-9=\-\/]+}/{[a-zA-Z0-9=\-\/]+}`,
		Index,
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
