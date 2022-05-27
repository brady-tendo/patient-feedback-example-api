package model

type Patient struct {
	Id        string     `json:"id,omitempty"`
	Active    bool       `json:"active,omitempty"`
	Name      []*Name    `json:"name,omitempty"`
	Contact   *[]Contact `json:"contact,omitempty"`
	Gender    string     `json:"gender,omitempty"`
	BirthDate string     `json:"birthDate,omitempty"`
	Address   *[]Address `json:"address,omitempty"`
}

type Contact struct {
	System string `json:"system,omitempty"`
	Value  string `json:"value,omitempty"`
	Use    string `json:"use,omitempty"`
}

type Address struct {
	Use  string    `json:"use,omitempty"`
	Line []*string `json:"line,omitempty"`
}
