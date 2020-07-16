package crawler

type DocumentNews struct {
	Title    string `json:"Title"`
	SubTitle string `json:"SubTitle"`
	Url      string `json:"Url"`
	//Document string `json:"Document"`
}

type DocRequest struct {
	Document        string   `json:"Document"`
	RelatedDocument map[string]string `json:"RelatedDocument"`
}

type DocResponse struct {
	Document        string   `json:"Document"`
	RelatedDocument map[string]string `json:"RelatedDocument"`
}

//type Document struct {
//	Name            string            `json:"Name"`
//	Url             string            `json:"Url"`
//	RelatedDocument []map[string]string `json:"RelatedDocument"`
//}
//
