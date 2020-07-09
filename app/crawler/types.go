package crawler

import url "net/url"

type Covid struct {
	Name string     `json:"name"`
	Url  url.Values `json:"url"`
}
