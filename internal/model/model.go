package model

import "time"

type Tenant struct {
	ID   string
	Name string
}

type Function struct {
	ID             string
	Name           string
	Tenant         Tenant
	ContainerImage string
}

type FunctionExecution struct {
	ID       string
	Function Function
	Started  time.Time
	Ended    time.Time
}
