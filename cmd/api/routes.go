package main

import (
	"net/http"

	"github.com/bmizerany/pat"
)

func (app *application) routes() http.Handler {

	mux := pat.New()

	mux.Get("/api/v1/:tenant_id/functions", http.HandlerFunc(app.listFunctions))
	mux.Get("/api/v1/:tenant_id/functions/:function_id", http.HandlerFunc(app.getFunction))
	mux.Post("/api/v1/:tenant_id/functions", http.HandlerFunc(app.createFunction))
	mux.Put("/api/v1/:tenant_id/functions/:function_id", http.HandlerFunc(app.updateFunction))
	mux.Del("/api/v1/:tenant_id/functions/:function_id", http.HandlerFunc(app.deleteFunction))
	mux.Get("/api/v1/:tenant_id/functions/:function_id/executions", http.HandlerFunc(app.listFunctionExecutions))
	mux.Get("/api/v1/:tenant_id/functions/:function_id/executions/:execution_id", http.HandlerFunc(app.getFunctionExecution))
	mux.Post("/api/v1/:tenant_id/functions/:function_id/executions", http.HandlerFunc(app.triggerFunctionExecution))

	return mux
}
