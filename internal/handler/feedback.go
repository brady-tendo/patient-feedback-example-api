package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tendosystems/patient-feedback/internal/model"
)

const FeedbackPath = "/feedback"

// GetFeedback godoc
// @Summary      Get patient feedback
// @Accept       json
// @Produce      json
// @Success      200  {object}  model.Feedback
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /feedback [get]
func GetFeedback(c echo.Context) error {
	m := model.NewModel()
	if m.Feedback == nil {
		return c.String(http.StatusNotFound, fmt.Sprintf("no feedback saved"))
	}
	resp, err := json.Marshal(m.Feedback)
	if err != nil {
		return c.String(http.StatusInternalServerError, "internal error reading saved feedback")
	}
	return c.String(http.StatusOK, string(resp))
}

// SaveFeedback godoc
// @Summary      Save patient feedback. Overwrites any previous feedback
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.Feedback
// @Failure      500  {object}  string
// @Router       /feedback [post]
func SaveFeedback(c echo.Context) error {
	m := model.NewModel()
	var feedback = model.Feedback{}
	if err := c.Bind(&feedback); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	m.Feedback = &feedback

	respString, err := json.Marshal(feedback.Body)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "Feedback saved successfully: "+string(respString))
}
