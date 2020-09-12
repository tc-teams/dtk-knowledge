package external

type BotRequest struct {
	Description string   `json:"description"`
	Related     []string `json:"related,omitiempty"`
}
