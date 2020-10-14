package crawler

import (
	"time"
)

const (
	GB      = "g1.globo.com"
	GV      = "www.saude.gov.br"
	BBC     = "www.bbc.com"
	UolNews = "noticias.uol.com.br"

	StartG1         = "https://g1.globo.com/bemestar/coronavirus/"
	//StartGV         = "https://antigo.saude.gov.br/fakenews/"
	StartFatoOuFake = "https://g1.globo.com/fato-ou-fake/coronavirus/"
	StartBBCNews    = "https://www.bbc.com/portuguese/topics/clmq8rgyyvjt"
	StartUol        = "https://noticias.uol.com.br/coronavirus"

	FilterGB  = "https://g1.globo\\.com/(bemestar.+)$"
	FilterFF  = "https://g1\\.globo\\.com/fato-ou-fake/(coronavirus.+)$"
	FilterGV  = "https://antigo.\\saude\\.gov\\.br/(fakenews.+)$"
	FilterBBC = "https://www\\.bbc\\.com/(portuguese.+)$"
)

var nl = []string{"coronavirus", "covid-19", "pandemia", "sars-cov-2",
	"cloroquina", "corona", "virus", "vírus",
	"vacina", "coronavac", "máscara", "coronavírus", "achatar a curva", "assintomático", "Autoisolamento", "caso suspeito",
	"distância social", "epidemia", "estado de calamidade", "grupo de risco", "paciente zero", "período de incubação", "quarentena",
	"taxa de transmissão", "teste rt-pcr", "transmissão comunitária ou sustentável"}

//RelatedNews is used to describe article model.
type RelatedNews struct {
	Url      string    `json:"url"`
	Time     time.Time `json:"time"`
	Date     string    `json:"date"`
	Title    string    `json:"title"`
	Subtitle string    `json:"subtitle"`
	Body     string    `validator:"required"`
	Msg      string    `json:"msg"`
}

var (
	strEmpty = string("")
)
