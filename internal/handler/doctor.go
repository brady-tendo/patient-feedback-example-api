package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tendosystems/patient-feedback/internal/model"
)

const DoctorPath = "/doctors"

// ListDoctors godoc
// @Summary      List all doctors
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.Doctor
// @Failure      500  {object}  string
// @Router       /doctors [get]
func ListDoctors(c echo.Context) error {
	m := model.NewModel()
	var doctors []*model.Doctor
	for _, p := range m.Doctors {
		doctors = append(doctors, p)
	}

	body, err := json.Marshal(doctors)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, string(body))
}

// GetDoctor godoc
// @Summary      Get a Doctor by ID
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Doctor
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /doctors/{id} [get]
func GetDoctor(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "doctor id is required")
	}
	m := model.NewModel()
	for _, p := range m.Doctors {
		if id == p.Id {
			doctor, err := json.Marshal(p)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			return c.String(http.StatusOK, string(doctor))
		}
	}
	return c.String(http.StatusNotFound, fmt.Sprintf("no doctor with id %v found", id))
}
