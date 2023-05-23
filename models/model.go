package model

type Netflix struct {
	Id      int    `json:"id,omitempty"`
	Movie   string `json:"movie,omitempty"`
	Watched bool   `json:"watched,omitempty"`
}
