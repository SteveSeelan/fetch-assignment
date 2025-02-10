package models

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Receipt struct {
	ID           string
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

// Validate the entire receipt object
func (r *Receipt) Validate() error {
	var validationErrors []string

	// Validate regex patterns
	if matched, err := regexp.MatchString(`^[\w\s\-&]+$`, r.Retailer); !matched || err != nil || r.Retailer == "" {
		validationErrors = append(validationErrors, "Retailer required and must match pattern: ^[\\w\\s\\-&]+$")
	}
	if matched, err := regexp.MatchString(`^\d+\.\d{2}$`, r.Total); !matched || err != nil || r.Total == "" {
		validationErrors = append(validationErrors, "Total required and must match pattern: ^\\d+\\.\\d{2}$")
	}

	// Validate date format (YYYY-MM-DD)
	if _, err := time.Parse("2006-01-02", r.PurchaseDate); err != nil || r.PurchaseDate == "" {
		validationErrors = append(validationErrors, fmt.Sprintf("PurchaseDate required and must be in YYYY-MM-DD format: %v", err))
	}

	// Validate time format (HH:MM)
	if _, err := time.Parse("15:04", r.PurchaseTime); err != nil || r.PurchaseTime == "" {
		validationErrors = append(validationErrors, fmt.Sprintf("PurchaseTime required and must be in HH:MM 24-hour format: %v", err))
	}

	// Validate each item
	for i, item := range r.Items {
		if err := item.Validate(); err != nil {
			validationErrors = append(validationErrors, fmt.Sprintf("Item %d: %v", i+1, err))
		}
	}

	// Return all errors as a single error message
	if len(validationErrors) > 0 {
		return errors.New(fmt.Sprintf("Validation errors:\n%s", strings.Join(validationErrors, "\n")))
	}
	return nil
}
