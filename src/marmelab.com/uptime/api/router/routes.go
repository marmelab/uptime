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
		"GetTargets",
		"GET",
		"/targets",
		handlers.GetTargets,
	},
	Route{
		"PostTarget",
		"POST",
		"/targets",
		handlers.PostTarget,
	},
	Route{
		"GetTarget",
		"GET",
		"/targets/{id}",
		handlers.GetTarget,
	},
	Route{
		"PutTarget",
		"PUT",
		"/targets/{id}",
		handlers.PutTarget,
	},
	Route{
		"DeleteTarget",
		"DELETE",
		"/targets/{id}",
		handlers.DeleteTarget,
	},
	Route{
		"GetResults",
		"GET",
		"/results",
		handlers.GetResults,
	},
	Route{
		"PostResult",
		"POST",
		"/results",
		handlers.PostResult,
	},
	Route{
		"GetResult",
		"GET",
		"/results/{id}",
		handlers.GetResult,
	},
	Route{
		"PutResult",
		"PUT",
		"/results/{id}",
		handlers.PutResult,
	},
	Route{
		"DeleteResult",
		"DELETE",
		"/results",
		handlers.DeleteResult,
	},
}
