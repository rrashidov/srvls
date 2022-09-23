package pers

import "github.com/rrashidov/srvls/internal/model"

type Persistence interface {
	SaveTenant(tenant *model.Tenant) error
	GetTenant(id string) (*model.Tenant, error)
	SaveFunction(function *model.Function) error
	SaveFunctionExecution(functionExec *model.FunctionExecution) error
}
