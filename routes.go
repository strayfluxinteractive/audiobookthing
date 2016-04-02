package main

import "net/http"

//Route represents a basic route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

//Routes holds our routes
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"AudiobookIndex",
		"GET",
		"/audiobooks",
		AudiobookIndex,
	},
}
