// @title API Go para o projeto de sugestões
// @version 1.0
// @description API de sugestões que faz a criação de sugestôes, retorno de sugestôes e alteração dos status.
// @host localhost:5000
// @BasePath /

// @contact.name Lucas
// @contact.email lucas@example.com

package main

import (
	"fmt"
	_ "go-api/docs" // importa os arquivos gerados
	"go-api/src/config"
	"go-api/src/repository/db"
	"go-api/src/routes"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

func main() {

	// Carregando constantes
	config.LoadVar()
	if config.API_port == "" {
		log.Fatal("Porta da API não configurada")
	}

	// Inicialização do bando de dados
	dbConnection, err := db.LoadDataBase()
	if err != nil {
		panic(err)
	}
	defer dbConnection.Close()

	// Carregando rotas
	router := routes.LoadRoutes()


	// Iniciando a API
	fmt.Printf("Api is Running on port %s\n", config.API_port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.API_port),
		handlers.CORS(
			handlers.AllowedOrigins([]string{"http://localhost:5173", "http://localhost"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		)(router)))

}