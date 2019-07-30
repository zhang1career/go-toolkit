package main

import (
	"github.com/gorilla/mux"
	"github.com/zhang1career/app/ruler/controller/condController"
	"github.com/zhang1career/app/ruler/controller/homeController"
	"github.com/zhang1career/app/ruler/controller/ruleController"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		homeController.IndexAction,
	},
	Route{
		"RuleIndex",
		"GET",
		"/rules",
		ruleController.IndexAction,
	},
	Route{
		"RuleView",
		"GET",
		"/rules/{ruleId}",
		ruleController.ViewAction,
	},
	Route{
		"RuleIndex",
		"GET",
		"/conds",
		condController.IndexAction,
	},
}