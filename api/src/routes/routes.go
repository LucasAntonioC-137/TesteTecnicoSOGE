package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
	//"go-api/api/src/controller"
)
type Route struct {
	URI string
	Method string
	Function func(http.ResponseWriter, *http.Request)
	AuthISRequired bool
}

func LoadRoutes() *mux.Router{

	router := mux.NewRouter()

	for _, route := range suggestionRoutes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	// Rota para Swagger UI (documentação)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	
	return router
}
