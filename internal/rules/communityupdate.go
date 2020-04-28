package rules

import (
	"encoding/json"
	"io"

	"github.com/TVolly/goapi-addresses/internal/models"
	"github.com/go-playground/validator"
)

type CommunityUpdateRule struct {
	Name string `json:"name" validate:"required"`
}

func (v *CommunityUpdateRule) Load(body io.Reader) error {
	if body == nil {
		return ErrNoBody
	}

	if err := json.NewDecoder(body).Decode(v); err != nil {
		return err
	}

	return v.validate()
}

func (v *CommunityUpdateRule) Fill(m *models.Community) {
	m.Name = v.Name
}

func (v *CommunityUpdateRule) validate() error {
	return validator.New().Struct(v)
}
