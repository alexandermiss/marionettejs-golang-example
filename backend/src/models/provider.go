package models

type Provider struct {
	ID      int64   `db:"id"`
	Name    string  `db:"name"`
}

type ProviderRequest struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
}

type ProviderResponse struct {
	ID      int64   `json:"id"`
	Name    string  `json:"name"`
}

