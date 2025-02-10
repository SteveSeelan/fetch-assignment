package models

import (
	"errors"
	"regexp"
	"strings"
)

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

// Validate the item object
func (i *Item) Validate() error {
	var validationErrors []string

	if matched, err := regexp.MatchString(`^[\w\s\-]+$`, i.ShortDescription); !matched || err != nil || i.ShortDescription == "" {
		validationErrors = append(validationErrors, "ShortDescription must match pattern: ^[\\w\\s\\-]+$")
	}
	if matched, err := regexp.MatchString(`^\d+\.\d{2}$`, i.Price); !matched || err != nil || i.Price == "" {
		validationErrors = append(validationErrors, "Price must match pattern: ^\\d+\\.\\d{2}$")
	}

	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, "\n"))
	}
	return nil
}
