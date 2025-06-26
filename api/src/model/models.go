package models

import (
	"errors"
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

func (suggestion *Suggestion) validate() error {

	if suggestion.CollaboratorName == "" {
		return errors.New("o campo do nome do colaborador não pode estar vazio")	
	} else if len(suggestion.CollaboratorName) < 3 {
		return errors.New("o nome do colaborador deve ter no mínimo 3 caracteres")
	} else if suggestion.Description == "" {
		return errors.New("o campo de descrição não pode estar vazio")
	} else if suggestion.Sector == "" {
		return errors.New("o campo do setor não pode estar vazio")
	}

	return nil
}

func (suggestion *Suggestion) format() error {
	suggestion.CollaboratorName = strings.TrimSpace(suggestion.CollaboratorName)
	suggestion.Sector = strings.TrimSpace(suggestion.Sector)
	suggestion.Description = strings.TrimSpace(suggestion.Description)

	return nil
}

func (suggestion *Suggestion) Prepare() error {
	if err := suggestion.validate(); err != nil {
		return err
	}

	if err := suggestion.format(); err != nil {
		return err
	}

	if suggestion.Status == "" {
		suggestion.Status = "open"
	}

	return nil
}
