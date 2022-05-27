package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tendosystems/patient-feedback/internal/model"
)

const DiagnosisPath = "/diagnoses"

// ListDiagnoses godoc
// @Summary      List all diagnoses
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.Diagnosis
// @Failure      500  {object}  string
// @Router       /diagnoses [get]
func ListDiagnoses(c echo.Context) error {
	m := model.NewModel()
	var diagnoses []*model.Diagnosis
	for _, p := range m.Diagnoses {
		diagnoses = append(diagnoses, p)
	}

	body, err := json.Marshal(diagnoses)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, string(body))
}

// GetDiagnosis godoc
// @Summary      Get a Diagnosis by ID
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Diagnosis
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /diagnoses/{id} [get]
func GetDiagnosis(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "diagnosis id is required")
	}
	m := model.NewModel()
	for _, p := range m.Diagnoses {
		if id == p.Id {
			diagnosis, err := json.Marshal(p)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			return c.String(http.StatusOK, string(diagnosis))
		}
	}
	return c.String(http.StatusNotFound, fmt.Sprintf("no diagnosis with id %v found", id))
}
