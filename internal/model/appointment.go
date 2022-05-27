package model

type Appointment struct {
	Id      string      `json:"id,omitempty"`
	Status  string      `json:"status,omitempty"`
	Type    []*ApptType `json:"type,omitempty"`
	Subject *Subject    `json:"subject,omitempty"`
	Actor   *Actor      `json:"actor,omitempty"`
	Period  *Period     `json:"period,omitempty"`
}

type ApptType struct {
	Text string `json:"text,omitempty"`
}

type Subject struct {
	Reference string `json:"reference,omitempty"`
}

type Actor struct {
	Reference string `json:"reference,omitempty"`
}

type Period struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}
