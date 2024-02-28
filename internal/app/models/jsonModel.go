package models

//go:generate easyjson -all jsonModel.go

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL string `json:"shortURL"`
}
