package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tendosystems/patient-feedback/internal/model"
)

const PatientPath = "/patients"

// ListPatients godoc
// @Summary      List all patients
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.Patient
// @Failure      500  {object}  string
// @Router       /patients [get]
func ListPatients(c echo.Context) error {
	m := model.NewModel()
	var patients []*model.Patient
	for _, p := range m.Patients {
		patients = append(patients, p)
	}

	body, err := json.Marshal(patients)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, string(body))
}

// GetPatient godoc
// @Summary      Get a Patient by ID
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Patient
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /patients/{id} [get]
func GetPatient(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "patient id is required")
	}
	m := model.NewModel()
	for _, p := range m.Patients {
		if id == p.Id {
			patient, err := json.Marshal(p)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			return c.String(http.StatusOK, string(patient))
		}
	}
	return c.String(http.StatusNotFound, fmt.Sprintf("no patient with id %v found", id))
}
