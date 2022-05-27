package model

import (
	"encoding/json"
	"fmt"
)

const (
	PatientResourceType     = "Patient"
	DoctorResourceType      = "Doctor"
	AppointmentResourceType = "Appointment"
	DiagnosisResourceType   = "Diagnosis"
)

type Bundle struct {
	Entries []*Entry `json:"entry,omitempty"`
}

type Entry struct {
	Resource json.RawMessage `json:"resource,omitempty"`
}

type Resource struct {
	ResourceType string `json:"resourceType,omitempty"`
}

func JSONToModel(data []byte, model *Model) error {
	fileData := Bundle{}
	err := json.Unmarshal(data, &fileData)
	if err != nil {
		return fmt.Errorf("couldn't unmarshal data: %v", err)
	}

	return Unbundle(&fileData, model)
}

func Unbundle(b *Bundle, model *Model) error {
	if model == nil {
		model = NewModel()
	}
	for _, entry := range b.Entries {
		resource := Resource{}
		if err := json.Unmarshal(entry.Resource, &resource); err != nil {
			return fmt.Errorf("problem unmarshalling resource: %v", err)
		}
		switch resource.ResourceType {
		case PatientResourceType:
			patient := Patient{}
			err := json.Unmarshal(entry.Resource, &patient)
			if err != nil {
				return fmt.Errorf("problem unmarshalling patient: %v", err)
			}
			model.Patients = append(model.Patients, &patient)
			break
		case DoctorResourceType:
			doctor := Doctor{}
			err := json.Unmarshal(entry.Resource, &doctor)
			if err != nil {
				return fmt.Errorf("problem unmarshalling doctor: %v", err)
			}
			model.Doctors = append(model.Doctors, &doctor)
			break
		case DiagnosisResourceType:
			diagnosis := Diagnosis{}
			err := json.Unmarshal(entry.Resource, &diagnosis)
			if err != nil {
				return fmt.Errorf("problem unmarshalling diagnosis: %v", err)
			}
			model.Diagnoses = append(model.Diagnoses, &diagnosis)
			break
		case AppointmentResourceType:
			appointment := Appointment{}
			err := json.Unmarshal(entry.Resource, &appointment)
			if err != nil {
				return fmt.Errorf("problem unmarshalling appointment: %v", err)
			}
			model.Appointments = append(model.Appointments, &appointment)
			break
		default:
			return fmt.Errorf("unknown resource type %v", resource.ResourceType)
		}
	}

	return nil
}
