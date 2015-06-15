package main

import "net/http"

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
		RetrieveTargets,
	},
	Route{
		"CreateTarget",
		"POST",
		"/targets",
		CreateTarget,
	},
	Route{
		"ShowTarget",
		"GET",
		"/targets/{id}",
		ShowTarget,
	},
	Route{
		"UpdateTarget",
		"PUT",
		"/ips/{id}",
		UpdateTarget,
	},
	Route{
		"DeleteTarget",
		"DELETE",
		"/ips/{id}",
		DeleteTarget,
	},
}
