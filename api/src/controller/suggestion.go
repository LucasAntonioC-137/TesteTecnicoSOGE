package controller

import (
	"encoding/json"
	"errors"
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

// GetSuggestionsGroupedByStatus retorna sugestões agrupadas ou filtradas por status
// @Summary Listar sugestões agrupadas ou filtradas por status
// @Description Se nenhum status for informado, retorna sugestões agrupadas por status. Se um status for passado, retorna apenas as sugestões com aquele status.
// @Tags suggestions
// @Produce json
// @Param status query string false "Filtrar por status (open, under review, implemented)"
// @Success 200 {object} []models.Suggestion
// @Failure 500 {object} models.ErrorResponse
// @Router /suggestions/grouped-by-status [get]
func GetSuggestionsGroupedByStatus(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")

	db, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	rep := repository.NewSuggestionInterface(db)

	// Se o status foi informado, filtra por ele
	if status != "" {
		suggestions, err := rep.FilterSuggestions(status, "")
		if err != nil {
			answers.Erro(w, http.StatusInternalServerError, err)
			return
		}
		answers.JSON(w, http.StatusOK, map[string][]models.Suggestion{
			status: suggestions,
		})
		return
	}

	// Caso contrário, retorna agrupado por status
	grouped, err := rep.GetSuggestionsGroupedByStatus()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, grouped)
}

// GetSuggestionsGroupedBySector retorna sugestões agrupadas ou filtradas por setor
// @Summary Listar sugestões agrupadas ou filtradas por setor
// @Description Se nenhum setor for informado, retorna sugestões agrupadas por setor. Se um setor for passado, retorna apenas as sugestões com aquele setor.
// @Tags suggestions
// @Produce json
// @Param sector query string false "Filtrar por setor (ex: IT, HR, Logistics)"
// @Success 200 {object} []models.Suggestion
// @Failure 500 {object} models.ErrorResponse
// @Router /suggestions/grouped-by-sector [get]
func GetSuggestionsGroupedBySector(w http.ResponseWriter, r *http.Request) {
	sector := r.URL.Query().Get("sector")

	db, err := db.LoadDataBase()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	rep := repository.NewSuggestionInterface(db)

	if sector != "" {
		suggestions, err := rep.FilterSuggestions("", sector)
		if err != nil {
			answers.Erro(w, http.StatusInternalServerError, err)
			return
		}
		answers.JSON(w, http.StatusOK, map[string][]models.Suggestion{
			sector: suggestions,
		})
		return
	}

	grouped, err := rep.GetSuggestionsGroupedBySector()
	if err != nil {
		answers.Erro(w, http.StatusInternalServerError, err)
		return
	}

	answers.JSON(w, http.StatusOK, grouped)
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