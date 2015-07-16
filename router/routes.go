package router

import (
	"github.com/cyarie/linksecret/application/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc handlers.WebHandler
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.WebHandler{
			H: handlers.IndexHandler,
		},
	},

	Route{
		"HashRedirect",
		"GET",
		"/{linkHash}",
		handlers.WebHandler{
			H: handlers.RedirectHash,
		},
	},

	Route{
		"GenerateLink",
		"POST",
		"/generate",
		handlers.WebHandler{
			H: handlers.GenerateLink,
		},
	},
}
