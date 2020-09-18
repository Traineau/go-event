package router

import (
	"github.com/gorilla/mux"
	"goevent/controllers"
	"net/http"
)

// Route struct defining all of this project routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes slice of Route
type Routes []Route

// newRouter registers public routes
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
		Name:        "Home",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: controllers.RenderHome,
	},
	Route{
		Name:        "List Orders",
		Method:      "GET",
		Pattern:     "/orders",
		HandlerFunc: controllers.GetOrders,
	},
	Route{
		Name:        "Create Order",
		Method:      "POST",
		Pattern:     "/order/new",
		HandlerFunc: controllers.CreateOrder,
	},
}
