package routes

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
)

// Estrutura das rotas
type Route struct {
	URI string
	Method string
	Function func(http.ResponseWriter, *http.Request)
}

// LoadRoutes carrega e disponibiliza todas as rotas já criadas
func LoadRoutes() *mux.Router{

	router := mux.NewRouter()

	for _, route := range suggestionRoutes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	// Rota para o Swagger (documentação)
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	
	return router
}
