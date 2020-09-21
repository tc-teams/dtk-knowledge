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
	PlnProcess  map[string]float64 `json:"nlp-process"`
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
	Similarity float64    `json:"similarity,omitempty"`
	Link       string    `json:"link"`
}
