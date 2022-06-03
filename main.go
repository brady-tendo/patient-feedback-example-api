package main

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/tendosystems/patient-feedback/docs"
	"github.com/tendosystems/patient-feedback/internal/handler"
	"github.com/tendosystems/patient-feedback/internal/model"
)

// Swagger doc stuff:
//
// @title Patient Body Test API
// @version 1.0
// @description A simple API that exposes fake patient information
//
// @contact.name Brady Lockhart
// @contact.email brady@tendo.com
//
// @host localhost:8000

const defaultPort = "8000"
const defaultFile = "patient-feedback-raw-data.json"

var bundleData []byte

func StartServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	e := echo.New()

	// Input file
	filename := os.Getenv("FILENAME")
	if filename == "" {
		filename = defaultFile
	}
	data, err := os.ReadFile(filename)
	if err != nil {
		panic("could not read input file: " + filename)
	}
	bundleData = data

	// Populate model
	m := model.NewModel()
	err = model.JSONToModel(data, m)
	if err != nil {
		panic("could not load model from input file: " + err.Error())
	}

	// Raw Data Endpoint
	e.GET("/", InputBundle)

	// Patient Endpoints
	e.GET(handler.PatientPath, handler.ListPatients)
	e.GET(handler.PatientPath+"/:id", handler.GetPatient)

	// Doctor Endpoints
	e.GET(handler.DoctorPath, handler.ListDoctors)
	e.GET(handler.DoctorPath+"/:id", handler.GetDoctor)

	// Diagnosis Endpoints
	e.GET(handler.DiagnosisPath, handler.ListDiagnoses)
	e.GET(handler.DiagnosisPath+"/:id", handler.GetDiagnosis)

	// Appointment Endpoints
	e.GET(handler.AppointmentPath, handler.ListAppointments)
	e.GET(handler.AppointmentPath+"/:id", handler.GetAppointment)

	// Body Endpoints
	e.GET(handler.FeedbackPath, handler.GetFeedback)
	e.PUT(handler.FeedbackPath, handler.SaveFeedback)

	// Swagger Docs
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Run server
	e.Logger.Fatal(e.Start(":" + port))
}

// InputBundle godoc
// @Summary      Return the raw resource bundle used to seed API responses
// @Accept       json
// @Produce      json
// @Success      200  {object}  string
// @Failure      500  {object}  string
// @Router       / [get]
func InputBundle(c echo.Context) error {
	return c.String(http.StatusOK, string(bundleData))
}

func main() {
	StartServer()
}
