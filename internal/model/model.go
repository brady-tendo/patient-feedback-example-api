package model

type Model struct {
	Patients     []*Patient
	Doctors      []*Doctor
	Appointments []*Appointment
	Diagnoses    []*Diagnosis
	Feedback     *Feedback
}

var globalModel *Model

// NewModel returns the existing singleton model, or creates a new one
// This is not thread safe!
func NewModel() *Model {
	if globalModel == nil {
		globalModel = &Model{}
	}
	return globalModel
}
