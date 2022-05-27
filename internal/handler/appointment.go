package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tendosystems/patient-feedback/internal/model"
)

const AppointmentPath = "/appointments"

// ListAppointments godoc
// @Summary      List all appointments
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.Appointment
// @Failure      500  {object}  string
// @Router       /appointments [get]
func ListAppointments(c echo.Context) error {
	m := model.NewModel()
	var appointments []*model.Appointment
	for _, p := range m.Appointments {
		appointments = append(appointments, p)
	}

	body, err := json.Marshal(appointments)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, string(body))
}

// GetAppointment godoc
// @Summary      Get a Appointment by ID
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Appointment
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /appointments/{id} [get]
func GetAppointment(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.String(http.StatusBadRequest, "appointment id is required")
	}
	m := model.NewModel()
	for _, p := range m.Appointments {
		if id == p.Id {
			appointment, err := json.Marshal(p)
			if err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			return c.String(http.StatusOK, string(appointment))
		}
	}
	return c.String(http.StatusNotFound, fmt.Sprintf("no appointment with id %v found", id))
}
