package router

import (
	"net/http"
	"../handlers"
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
		"RetrieveTargets",
		"GET",
		"/targets",
		handlers.RetrieveTargets,
	},
	Route{
		"CreateTarget",
		"POST",
		"/targets",
		handlers.CreateTarget,
	},
	Route{
		"ShowTarget",
		"GET",
		"/targets/{id}",
		handlers.ShowTarget,
	},
	Route{
		"UpdateTarget",
		"PUT",
		"/targets/{id}",
		handlers.UpdateTarget,
	},
	Route{
		"DeleteTarget",
		"DELETE",
		"/targets/{id}",
		handlers.DeleteTarget,
	},
}
