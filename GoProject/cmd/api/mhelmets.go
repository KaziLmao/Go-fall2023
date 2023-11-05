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
		Year:          int32(input.Year),
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
		Name          *string  `json:"name"`
		Year          *int32   `json:"year"`
		Material      *string  `json:"material"`
		Ventilation   *bool    `json:"ventilation"`
		Protection    *string  `json:"protection"`
		Weight        *float64 `json:"weight"`
		SunProtection *bool    `json:"sun_protection"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Name != nil {
		helmet.Name = *input.Name
	}
	if input.Year != nil {
		helmet.Year = *input.Year
	}
	if input.Material != nil {
		helmet.Material = *input.Material
	}
	if input.Ventilation != nil {
		helmet.Ventilation = *input.Ventilation
	}
	if input.Protection != nil {
		helmet.Protection = *input.Protection
	}
	if input.Weight != nil {
		helmet.Weight = *input.Weight
	}
	if input.SunProtection != nil {
		helmet.SunProtection = *input.SunProtection
	}

	v := validator.New()
	if data.ValidateHelmet(v, helmet); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Helmets.Update(helmet)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
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

func (app *application) listMHelmetsHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name       string
		Material   string
		Protection string
		data.Filters
	}
	v := validator.New()
	qs := r.URL.Query()
	input.Name = app.readString(qs, "name", "")
	input.Material = app.readString(qs, "material", "")
	input.Protection = app.readString(qs, "protection", "")
	input.Filters.Page = app.readInt(qs, "page", 1, v)
	input.Filters.PageSize = app.readInt(qs, "page_size", 20, v)
	input.Filters.Sort = app.readString(qs, "sort", "id")
	input.Filters.SortSafelist = []string{"id", "name", "year", "material", "ventilation", "protection", "weight", "sun_protection",
		"-id", "-name", "-year", "-material", "-ventilation", "-protection", "-weight", "-sun_protection"}

	if data.ValidateFilters(v, input.Filters); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	helmets, metadata, err := app.models.Helmets.GetAll(input.Name, input.Material, input.Protection, input.Filters)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"helmets": helmets, "metadata": metadata}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
