package api

const (
	content string = "application/json; charset=UTF-8"
)


type HandlerInfo struct {
   Method  string  `json:"method"`
   Path  string    `json:"path"`
   Name   string   `json:"name,omitempty"`
}


