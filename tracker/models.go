package tracker

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
	Url      string              `validate:"required,max=500"`
	Date     time.Time            `validate:"required,max=10"`
	Title    string              `validate:"required,max=500"`
	Subtitle string              `validate:"required,max=500"`
	Body     string              `validate:"required,max=500"`
}
