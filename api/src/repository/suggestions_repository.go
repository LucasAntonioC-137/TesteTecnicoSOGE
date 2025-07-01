package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"go-api/src/model"
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

// FilterSuggestions retorna sugestões filtradas por um status ou setor. Se nenhum argumento for informado, retorna erro.
func (repository *Suggestions) FilterSuggestions(status, sector string) ([]models.Suggestion, error) {
	// Nenhum argumento foi informado, retorna erro
	if status == "" && sector == "" {
		return nil, errors.New("pelo menos um filtro (status ou setor) deve ser informado")
	}

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

	lines, err := repository.database.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var suggestions []models.Suggestion
	for lines.Next() {
		var suggestion models.Suggestion
		if err := lines.Scan(
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

// GetSuggestionsGroupedByStatus retorna todas as sugestões agrupadas por status (ordem: open, under review, implemented)
func (repository *Suggestions) GetSuggestionsGroupedByStatus() ([]models.Suggestion, error) {
	query := `
		SELECT id, collaborator_name, sector, description, status, created_at
		FROM suggestions
		ORDER BY 
			CASE 
				WHEN status = 'open' THEN 1
				WHEN status = 'under review' THEN 2
				WHEN status = 'implemented' THEN 3
				ELSE 4
			END, created_at ASC
	`

	rows, err := repository.database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var suggestions []models.Suggestion
	for rows.Next() {
		var suggestion models.Suggestion
		err := rows.Scan(
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

// GetSuggestionsGroupedBySector retorna todas as sugestões agrupadas por setor
func (repository *Suggestions) GetSuggestionsGroupedBySector() (map[string][]models.Suggestion, error) {
	query := `SELECT id, collaborator_name, sector, description, status, created_at FROM suggestions ORDER BY sector`

	rows, err := repository.database.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	grouped := make(map[string][]models.Suggestion)

	for rows.Next() {
		var s models.Suggestion
		if err := rows.Scan(&s.ID, &s.CollaboratorName, &s.Sector, &s.Description, &s.Status, &s.CreatedAt); err != nil {
			return nil, err
		}
		grouped[s.Sector] = append(grouped[s.Sector], s)
	}

	return grouped, nil
}

// UpdateSuggestionStatus a recebe atualiza o status da sugestão escolhida
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

// GetSuggestionStatusByID pega o status da sugestão com ID correspondente
func (repository *Suggestions) GetSuggestionStatusByID(id uint) (string, error) {
	var status string
	err := repository.database.QueryRow("SELECT status FROM suggestions WHERE id = $1", id).Scan(&status)
	return status, err
}


