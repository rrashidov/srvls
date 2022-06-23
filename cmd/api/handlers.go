package main

import "net/http"

func (app *application) listFunctions(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List functions"))
}

func (app *application) getFunction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get function"))
}

func (app *application) createFunction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create function"))
}

func (app *application) updateFunction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update function"))
}

func (app *application) deleteFunction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete function"))
}

func (app *application) listFunctionExecutions(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("List function executions"))
}

func (app *application) getFunctionExecution(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get function execution"))
}

func (app *application) triggerFunctionExecution(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Trigger function execution"))
}
