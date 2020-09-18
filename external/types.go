package external

import (
	"time"
)

type PlnRequest struct {
	Description string   `json:"description"`
	News        []string `json:"news"`
}

type PlnResponse struct {
	Description string            `json:"description"`
	PlnProcess  map[string]string `json:"pln-process"`
}

type BotRequest struct {
	Description string `json:"description"`
}

type BotResponse struct {
	Description string       `json:"description"`
	Text        []TextResult `json:"text,omitempty"`
}
type TextResult struct {
	Date       time.Time `json:"date"`
	Title      string    `json:"title"`
	Similarity string    `json:"similarity,omitempty"`
	Link       string    `json:"link"`
}
