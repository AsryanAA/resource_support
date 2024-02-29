package models

type Position struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DivisionId  string `json:"division_id"`
}
