package model

type Diagnosis struct {
	Id          string   `json:"id,omitempty"`
	Meta        *Meta    `json:"meta,omitempty"`
	Status      string   `json:"status,omitempty"`
	Code        *Code    `json:"code,omitempty"`
	Appointment *ApptRef `json:"appointment,omitempty"`
}

type Meta struct {
	LastUpdated string `json:"lastUpdated,omitempty"`
}

type Code struct {
	Coding *[]Coding `json:"coding,omitempty"`
}

type Coding struct {
	System string `json:"system,omitempty"`
	Code   string `json:"code,omitempty"`
	Name   string `json:"name,omitempty"`
}

type ApptRef struct {
	Reference string `json:"reference,omitempty"`
}
