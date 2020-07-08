package tracker

const (
	G1         = "g1.globo.com"
	Folha      = "www1.folha.uol.com.br"
	Uol        = "noticias.uol.com.br"
	StartFolha = "https://www1.folha.uol.com.br/cotidiano/coronavirus/"
	StartG1    = "https://g1.globo.com/bemestar/coronavirus/"
	StartUol   = "https://noticias.uol.com.br/coronavirus/"
)

//News is used to describe article model.
type News struct {
	Title    string `validate:"required,max=500"`
	SubTitle string `validate:"required,max=500"`
	//Date     string `validate:"required,max=500"`
	Page string `validate:"required,max=500"`
}
