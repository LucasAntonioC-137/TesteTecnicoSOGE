package controller

import (
	"encoding/json"
	"errors"
	"go-api/src/answers"
	models "go-api/src/model"

	//"go-api/api/src/model"
	"go-api/src/repository"
	"go-api/src/repository/db"
	"io"
	"net/http"

	//"github.com/gorilla/mux"
)

type suggestionController struct {
	//Usecase
}

func NewSuggestionController() suggestionController {
	return suggestionController{}
}

// Lista todas as sugestões disponíveis no banco de dados
func GetSuggestions(w http.ResponseWriter, r*http.Request) {

	bd, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	defer bd.Close()

	rep := repository.NewSuggestionInterface(bd)

	suggestions, err := rep.GetSuggestions()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, suggestions)
}

// Cria uma sugestão no banco de dados
func CreateSuggestion(w http.ResponseWriter, r*http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		answers.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var suggestion models.Suggestion

	if err = json.Unmarshal(requestBody, &suggestion); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	// Faz a validação dos campos e define o status para aberto quando ele está nulo
	if err = suggestion.Prepare(); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	bd, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer bd.Close()

	rep := repository.NewSuggestionInterface(bd)
	
	suggestion.ID, err = rep.CreateSuggestion(suggestion)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}
	
	answers.JSON(w, http.StatusCreated, suggestion.ID)
}

// Filtra todas as sugestões disponíveis no banco de dados por seu status
func GetSuggestionsByStatus(w http.ResponseWriter, r*http.Request) {
	
	//vars := mux.Vars(r)
	//status := vars["status"]
	//if status == "" {
	//	answers.Erro(w, http.StatusBadRequest, errors.New("definir um status é obrigatório"))
	//	return
	//}
	status := r.URL.Query().Get("status")

	if status == "" {
		answers.Erro(w, http.StatusBadRequest, errors.New("o parâmetro 'status' é obrigatório"))
		return
	}

	bd, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusUnauthorized, err)
		return
	}

	defer bd.Close()


	rep := repository.NewSuggestionInterface(bd)

	suggestions, err := rep.GetSuggestionsByStatus(status)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, suggestions)
}

// Função que retorna todas as sugestões ou também retorna as mesmas filtradas por alguma caracteristica
func GetSuggestionsWithFilters(w http.ResponseWriter, r *http.Request) {
	// Pega os filtros da URL, se existirem
	status := r.URL.Query().Get("status")
	sector := r.URL.Query().Get("sector")

	db, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	// Chama o repositório, passando os filtros (podem estar vazios)
	rep := repository.NewSuggestionInterface(db)
	suggestions, err := rep.FilterSuggestions(status, sector)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, suggestions)
}


