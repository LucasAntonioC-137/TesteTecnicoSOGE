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

// Criação de sugestoes no banco de dados
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
