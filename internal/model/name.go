package model

type Name struct {
	Text   string    `json:"text,omitempty"`
	Family string    `json:"family,omitempty"`
	Given  []*string `json:"given,omitempty"`
}
