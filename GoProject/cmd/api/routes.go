package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/mhelmets", app.createMHelmetHandler)
	router.HandlerFunc(http.MethodGet, "/v1/mhelmets/:id", app.showMHelmetHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/mhelmets/:id", app.updateMHelmetHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/mhelmets/:id", app.deleteMHelmetHandler)
	return router
}
