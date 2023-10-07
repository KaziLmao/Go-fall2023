package main

import (
	"GoProject/internal/data"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createMHelmetHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name          string  `json:"name"`
		Year          int32   `json:"year"`
		Material      string  `json:"material"`
		Ventilation   bool    `json:"ventilation"`
		Protection    string  `json:"protection"`
		Design        string  `json:"design"`
		Weight        float64 `json:"weight"`
		SunProtection bool    `json:"sun_protection"`
		Lining        string  `json:"lining"`
		Fastening     string  `json:"fastening"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMHelmetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	helmet := data.Helmet{
		ID:            id,
		CreatedAt:     time.Now(),
		Name:          "SpeedMaster X1",
		Year:          2022,
		Material:      "Carbon Fiber",
		Ventilation:   true,
		Protection:    "Snell",
		Design:        "Aerodynamic Racing Design",
		Weight:        1.2,
		SunProtection: false,
		Lining:        "Moisture-wicking and Antibacterial Fabric",
		Fastening:     "Double D-ring Chin Strap",
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"helmet": helmet}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
