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
		"/ips",
		RetrieveTargets,
	},
	Route{
		"CreateTarget",
		"POST",
		"/ips",
		CreateTarget,
	},
	Route{
		"ShowTarget",
		"GET",
		"/ips/{id}",
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
