package router

import (
	"github.com/cyarie/linksecret/application/environment"
	"github.com/cyarie/linksecret/application/handlers"
	"github.com/gorilla/mux"
)

func Router(env *environment.Env) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler handlers.WebHandler

		handler = handlers.WebHandler{env, route.HandlerFunc.H}
		router.Methods(route.Method).Path(route.Pattern).Name(route.Name).Handler(&handler)
	}

	return router
}