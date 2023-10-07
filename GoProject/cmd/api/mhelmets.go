package main

import (
	"GoProject/internal/data"
	"fmt"
	"net/http"
	"time"
)

func (app *application) createMHelmetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new motorcycle helmet")
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

	err = app.writeJSON(w, http.StatusOK, envelope{"Moto helmet": helmet}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
