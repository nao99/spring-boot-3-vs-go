package model

type Product struct {
	ID          uint64 `json:"id"`
	Category    string `json:"category"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
