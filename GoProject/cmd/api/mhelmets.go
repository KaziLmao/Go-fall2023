package main

import (
	"GoProject/internal/data"
	"GoProject/internal/validator"
	"errors"
	"fmt"
	"net/http"
)

func (app *application) createMHelmetHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name          string  `json:"name"`
		Year          int32   `json:"year"`
		Material      string  `json:"material"`
		Ventilation   bool    `json:"ventilation"`
		Protection    string  `json:"protection"`
		Weight        float64 `json:"weight"`
		SunProtection bool    `json:"sun_protection"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	helmet := &data.Helmet{
		Name:          input.Name,
		Year:          input.Year,
		Material:      input.Material,
		Ventilation:   input.Ventilation,
		Protection:    input.Protection,
		Weight:        input.Weight,
		SunProtection: input.SunProtection,
	}

	v := validator.New()

	if data.ValidateHelmet(v, helmet); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Helmets.Insert(helmet)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", helmet.ID))

	err = app.writeJSON(w, http.StatusCreated, envelope{"helmet": helmet}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) showMHelmetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	helmet, err := app.models.Helmets.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"helmet": helmet}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateMHelmetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	helmet, err := app.models.Helmets.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Name          string  `json:"name"`
		Year          int32   `json:"year"`
		Material      string  `json:"material"`
		Ventilation   bool    `json:"ventilation"`
		Protection    string  `json:"protection"`
		Weight        float64 `json:"weight"`
		SunProtection bool    `json:"sun_protection"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	helmet.Name = input.Name
	helmet.Year = input.Year
	helmet.Material = input.Material
	helmet.Ventilation = input.Ventilation
	helmet.Protection = input.Protection
	helmet.Weight = input.Weight
	helmet.SunProtection = input.SunProtection

	v := validator.New()
	if data.ValidateHelmet(v, helmet); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Helmets.Update(helmet)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"helmet": helmet}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteMHelmetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Helmets.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "motorcycle helmet successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
