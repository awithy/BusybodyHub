package main

import (
	"net/http"

	"github.com/awithy/busybodyhub/bbserver/controllers"
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
		"LoginAccount",
		"POST",
		"/api/account/login",
		controllers.LoginAccount,
	},
	Route{
		"PublicCatchAll",
		"GET",
		`/{a:.+}`,
		Index,
	},
	Route{
		"Index",
		"GET",
		`/`,
		Index,
	},
}
