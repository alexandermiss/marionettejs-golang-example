package models

type Token struct {
	ID          int64           `db:"id"`
	Name        string          `db:"name"`
	Client      *Client         `db:"client"`
	ClientID    int64           `db:"clientID"`
}

type TokenRequest struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	Client      *ClientRequest  `json:"client"`
	ClientID    int64           `json:"client_id"`
}

type TokenResponse struct {
	ID          int64           `json:"id"`
	Name        string          `json:"name"`
	Client      *ClientResponse `json:"client"`
	ClientID    int64           `json:"client_id"`
}

