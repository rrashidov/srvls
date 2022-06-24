package model

import "time"

type FunctionExecutionStatus int

const (
	FunctionExecutionStatusOK FunctionExecutionStatus = iota
	FunctionExecutionStatusERROR
)

type Tenant struct {
	id          string
	name        string
	description string
}

type Function struct {
	id              string
	tenant_id       string
	name            string
	description     string
	container_image string
}

type FunctionExecution struct {
	id          string
	function_id string
	triggered   time.Time
	started     time.Time
	finished    time.Time
	status      FunctionExecutionStatus
}
