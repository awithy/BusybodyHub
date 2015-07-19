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
	Auth        bool
}

type Routes []Route

var routes = Routes{
	Route{
		"LoginAccount",
		"POST",
		"/api/account/login",
		controllers.LoginAccount,
		false,
	},
	Route{
		"Logout",
		"POST",
		"/api/account/logout",
		controllers.Logout,
		true,
	},
	Route{
		"PublicCatchAll",
		"GET",
		`/{a:.+}`,
		Index,
		false,
	},
	Route{
		"Index",
		"GET",
		`/`,
		Index,
		false,
	},
}
