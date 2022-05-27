package model

type Doctor struct {
	Id   string  `json:"id,omitempty"`
	Name *[]Name `json:"name,omitempty"`
}
