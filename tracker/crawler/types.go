package crawler

import (
	"time"
)

const (
	GB         = "g1.globo.com"
	Folha      = "www1.folha.uol.com.br"
	Uol        = "noticias.uol.com.br"
	StartFolha = "https://www1.folha.uol.com.br/cotidiano/coronavirus/"
	StartG1    = "https://g1.globo.com/bemestar/coronavirus/"
	StartUol   = "https://noticias.uol.com.br/coronavirus/"

	FilterGB = "https://g1.globo\\.com/(bemestar.+)$"
)

//RelatedNews is used to describe article model.
type RelatedNews struct {
	Url      string    `json:"url"`
	Time     time.Time `json:"time"`
	Date     time.Time `json:"date"`
	Title    string    `json:"title"`
	Subtitle string    `json:"subtitle"`
	Body     string    `validator:"required"`
	Msg      string    `json:"msg"`
}

type Teste struct {
	Teste string `validator:"required"`
}
