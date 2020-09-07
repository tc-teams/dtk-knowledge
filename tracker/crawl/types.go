package crawl

import "time"

const (
	G1         = "g1.globo.com"
	Folha      = "www1.folha.uol.com.br"
	Uol        = "noticias.uol.com.br"
	StartFolha = "https://www1.folha.uol.com.br/cotidiano/coronavirus/"
	StartG1    = "https://g1.globo.com/bemestar/coronavirus/"
	StartUol   = "https://noticias.uol.com.br/coronavirus/"
)


//RelatedNews is used to describe article model.
type RelatedNews struct {
	Url      string    `json:"Url""`
	Time     time.Time `json:"time"`
	Date     time.Time `json:"Date"`
	Title    string    `json:"Title"`
	Subtitle string    `json:"Subtitle"`
	Body     string    `json:"Body"`
	msg      string    `json:"msg,omitempty"`
}
