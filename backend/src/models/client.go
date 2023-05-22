package models

type Client struct {
	ID              int64               `db:"id"`
	Name            string              `db:"name"`
	Provider        *Provider           `db:"provider"`
	ProviderID      int64               `db:"providerID"`
}

type ClientRequest struct {
	ID              int64               `json:"id"`
	Name            string              `json:"name"`
	Provider        *ProviderRequest    `json:"provider"`
	ProviderID      int64               `json:"provider_id"`
}

type ClientResponse struct {
	ID              int64               `json:"id"`
	Name            string              `json:"name"`
	Provider        *ProviderResponse   `json:"provider"`
	ProviderID      int64               `json:"provider_id"`
}

