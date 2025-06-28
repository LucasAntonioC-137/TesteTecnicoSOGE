package repository

import (
	"database/sql"
	"fmt"
	models "go-api/src/model"
)

type Suggestions struct {
	database *sql.DB
}

func NewSuggestionInterface(database *sql.DB) *Suggestions {
	return &Suggestions{database}
}

// CreateSuggestion se comunica com o banco de dados para criar e incluir a sugestão na tabela
func (repository *Suggestions) CreateSuggestion(suggestion models.Suggestion) (uint, error) {

	statement, err := repository.database.Prepare(`
	INSERT INTO suggestions (collaborator_name, sector, description) 
	VALUES ($1, $2, $3) RETURNING id`)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var id int
	err = statement.QueryRow(suggestion.CollaboratorName, suggestion.Sector, suggestion.Description).Scan(&id)
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}

// GetSuggestions faz uma consulta ao banco de dados e retorna todas as sugestôes registradas
func (repository *Suggestions) GetSuggestions() ([]models.Suggestion, error) {

	lines, err := repository.database.Query(`SELECT id, collaborator_name, sector, description, status, created_at FROM suggestions`)
	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var suggestions []models.Suggestion

	for lines.Next() {
		var suggestion models.Suggestion

		if err = lines.Scan(
			&suggestion.ID,
			&suggestion.CollaboratorName,
			&suggestion.Sector,
			&suggestion.Description,
			&suggestion.Status,
			&suggestion.CreatedAt,
		); err != nil {
			return nil, err
		}

		suggestions = append(suggestions, suggestion)
	}

	return suggestions, nil
}

func (repository *Suggestions) GetSuggestionsByStatus(status string) ([]models.Suggestion, error) {

	lines, err := repository.database.Query(`
		SELECT id, collaborator_name, sector, description, status, created_at
		FROM suggestions
		WHERE status = $1`, status)
	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var suggestions []models.Suggestion

	for lines.Next() {
		var suggestion models.Suggestion

		if err = lines.Scan(
			&suggestion.ID,
			&suggestion.CollaboratorName,
			&suggestion.Sector,
			&suggestion.Description,
			&suggestion.Status,
			&suggestion.CreatedAt,
		); err != nil {
			return nil, err
		}

		suggestions = append(suggestions, suggestion)
	}

	return suggestions, nil
}

func (repository *Suggestions) FilterSuggestions(status, sector string) ([]models.Suggestion, error) {
	query := `SELECT id, collaborator_name, sector, description, status, created_at FROM suggestions WHERE 1=1`
	var params []interface{}
	paramCount := 1

	// Verifica se foi passado algum parâmetro de filtro
	if status != "" {
		query += fmt.Sprintf(" AND status = $%d", paramCount)
		params = append(params, status)
		paramCount++
	}
	if sector != "" {
		query += fmt.Sprintf(" AND sector = $%d", paramCount)
		params = append(params, sector)
		paramCount++
	}

	fmt.Println(params...)

	lines, err := repository.database.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var suggestions []models.Suggestion
	for lines.Next() {
		var suggestion models.Suggestion
		err := lines.Scan(
			&suggestion.ID,
			&suggestion.CollaboratorName,
			&suggestion.Sector,
			&suggestion.Description,
			&suggestion.Status,
			&suggestion.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		suggestions = append(suggestions, suggestion)
	}

	return suggestions, nil
}

func (repository *Suggestions) UpdateSuggestionStatus(id int, status string) error {
	statement, err := repository.database.Prepare("UPDATE suggestions SET status = $1 WHERE id = $2")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err := statement.Exec(status, id); err != nil {
		return err
	}
	return nil
}

func (repository *Suggestions) GetSuggestionStatusByID(id uint) (string, error) {
	var status string
	err := repository.database.QueryRow("SELECT status FROM suggestions WHERE id = $1", id).Scan(&status)
	return status, err
}


