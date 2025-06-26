package main

import (
	"go-api/src/config"
	"go-api/src/repository/db"
	"go-api/src/routes"
	"log"
	"net/http"
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


	// Teste de ping
	log.Println("Servidor rodando na porta " + config.API_port + " ...")
	err = http.ListenAndServe(":" + config.API_port, router)
	if err != nil {
		log.Fatal("Erro ao subir o servidor:", err)
	}

}