package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-api/src/answers"
	"go-api/src/model"
	"go-api/src/repository"
	"go-api/src/repository/db"
	"io"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

// GetSuggestions Lista todas as sugestões cadastradas no banco de dados
// @Summary Lista as sugestões
// @Description Lista todas as sugestões cadastradas no banco de dados
// @Tags suggestions
// @Accept json
// @Produce json
// @Success 200 {array} models.Suggestion
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /suggestions [get]
func GetSuggestions(w http.ResponseWriter, r *http.Request) {

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

// CreateSuggestion Cria uma nova sugestão
// @Summary Criar sugestão
// @Description Cria uma nova sugestão com os dados informados
// @Tags suggestions
// @Accept json
// @Produce json
// @Param suggestion body models.CreateSuggestionInput true "Dados da nova sugestão"
// @Success 201 {integer} int
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /register [post]
func CreateSuggestion(w http.ResponseWriter, r *http.Request) {
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

	// Faz a validação dos campos e define o status para "open" quando ele está vazio
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

// GetSuggestionsWithFilters retorna todas as sugestões ou sugestões filtradas por status e setor
// @Summary Buscar sugestões com ou sem filtros
// @Description Retorna todas as sugestões se você optar por não usar nenhum filtro, mas também é possível usar um filtro de cada vez ou até mesmo os dois em conjunto
// @Tags suggestions
// @Accept json
// @Produce json
// @Param status query string false "Status da sugestão (ex: open, implemented, under review)"
// @Param sector query string false "Setor relacionado à sugestão"
// @Success 200 {array} models.Suggestion
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /suggestions [get]
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

// UpdateSuggestionStatus Atualiza o status de uma sugestão existente
// @Summary Atualizar status da sugestão
// @Description Altera o status de uma sugestão existente pelo ID
// @Tags suggestions
// @Accept json
// @Produce json
// @Param id path int true "ID da sugestão"
// @Param status body models.UpdateStatusInput true "Novo status da sugestão"
// @Success 200 {integer} int
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /suggestions/{id}/status [put]
func UpdateSuggestionStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idParam := vars["id"]

	id, err := strconv.Atoi(idParam)
	if err != nil {
		answers.Erro(w, http.StatusBadRequest, errors.New("ID inválido"))
		return
	}

	var input struct {
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println("Status extraido:" + input.Status)

	if err := models.ValidateStatus(input.Status); err != nil {
		answers.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.NewSuggestionInterface(db)

	// Busca status o atual no banco
	currentStatus, err := repo.GetSuggestionStatusByID(uint(id))
	fmt.Println("Status atual:" + currentStatus)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	// Se o status já for o mesmo, retorna mensagem de "inalterado"
	if currentStatus == input.Status {
		answers.JSON(w, http.StatusOK, map[string]string{"mensagem": "O status já está definido como " + input.Status})
		return
	}

	err = repo.UpdateSuggestionStatus(id, input.Status)
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}
	
	answers.JSON(w, http.StatusOK, map[string]string{"mensagem": "Status atualizado com sucesso"})
}