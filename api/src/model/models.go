package models

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Suggestion struct {
	ID               uint    `json:"id_suggestion"`
	CollaboratorName string `json:"collaborator_name"`
	Sector           string `json:"sector"`
	Description      string `json:"description"`
	Status           string `json:"status"`
	CreatedAt		 time.Time `json:"created_at"`
}

type CreateSuggestionInput struct {
	CollaboratorName string `json:"collaborator_name" example:"João da Silva"`
	Sector           string `json:"sector" example:"TI"`
	Description      string `json:"description" example:"Sugestão para melhorar o fluxo de trabalho"`
}

type UpdateStatusInput struct {
	Status string `json:"status" example:"open"`
}

type ErrorResponse struct {
    Erro string `json:"erro"`
}

func isValidName(name string) bool {
	// Permite letras (maiúsculas/minúsculas), espaços, acentos e hífen, mas regeita números e caracteres especiais
	regex := regexp.MustCompile(`^[A-Za-zÀ-ÿ\s\-]+$`)
	return regex.MatchString(name)
}

// validate faz uma validação do nome, descrição e setor fornecidos pelo colaborador no momento de criar a sugestão
func (suggestion *Suggestion) validate() error {

	if suggestion.CollaboratorName == "" {
		return errors.New("o campo do nome do colaborador não pode estar vazio")	
	} else if !isValidName(suggestion.CollaboratorName) || len(suggestion.CollaboratorName) < 3 {
		return errors.New("o nome do colaborador deve conter apenas letras e espaços, tendo um tamanho mínimo de 3 caracteres")
	} else if suggestion.Description == "" {
		return errors.New("o campo de descrição não pode estar vazio")
	} else if suggestion.Sector == "" {
		return errors.New("o campo do setor não pode estar vazio")
	}

	return nil
}

// format faz a remoção de espaços no começo e no fim do texto
func (suggestion *Suggestion) format() error {
	suggestion.CollaboratorName = strings.TrimSpace(suggestion.CollaboratorName)
	suggestion.Sector = strings.TrimSpace(suggestion.Sector)
	suggestion.Description = strings.TrimSpace(suggestion.Description)

	return nil
}

// ValidateSatus faz uma verificação para ver se o status enviado é um dos três pré-definidos
func ValidateStatus(status string) error {
	validStatuses := []string{"open", "under review", "implemented"}

	for _, s := range validStatuses {
		if status == s {
			return nil
		}
	}
	return errors.New("status inválido: deve ser 'open', 'under review' ou 'implemented'")
}


func (suggestion *Suggestion) Prepare() error {
	if err := suggestion.validate(); err != nil {
		return err
	}

	if err := suggestion.format(); err != nil {
		return err
	}

	fmt.Println(suggestion.Status)

	return nil
}
