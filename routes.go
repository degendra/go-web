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
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Home",
		"GET",
		"/pages/{level}/{link}",
		Page,
	},
	Route{
		"Index",
		"HEAD",
		"/",
		Index,
	},
	Route{
		"Home",
		"HEAD",
		"/pages/{level}/{link}",
		Page,
	},
}
