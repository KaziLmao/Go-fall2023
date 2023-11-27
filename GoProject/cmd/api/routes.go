package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/mhelmets", app.listMHelmetsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/mhelmets", app.createMHelmetHandler)
	router.HandlerFunc(http.MethodGet, "/v1/mhelmets/:id", app.showMHelmetHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/mhelmets/:id", app.updateMHelmetHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/mhelmets/:id", app.deleteMHelmetHandler)

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	return app.recoverPanic(app.rateLimit(router))
}
