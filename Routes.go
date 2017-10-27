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
		"CategoryIndex",
		"GET",
		"/categories",
		CategoryIndex,
	},
	Route{
		"CategoryCreate",
		"POST",
		"/category",
		CategoryCreate,
	},
	Route{
		"CategoryShow",
		"GET",
		"/category/{name}",
		CategoryShow,
	},
	Route{
		"CategoryDelete",
		"DELETE",
		"category/{categoryId}",
		CategoryDelete,
	},
}
