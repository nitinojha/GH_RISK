package models

import "errors"

// Risk represents the structure of a risk
type Risk struct {
	ID          string `json:"id"`
	State       string `json:"state"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// Validate checks if the Risk fields are valid
func (r *Risk) Validate() error {
	validStates := map[string]bool{
		"open":          true,
		"closed":        true,
		"accepted":      true,
		"investigating": true,
	}

	if r.State == "" || !validStates[r.State] {
		return errors.New("invalid or missing state")
	}
	if r.Title == "" {
		return errors.New("title is required")
	}
	if r.Description == "" {
		return errors.New("description is required")
	}
	return nil
}
