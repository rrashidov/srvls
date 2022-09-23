package inmem

import (
	"errors"

	"github.com/rrashidov/srvls/internal/model"
)

var (
	ErrTenantNotFound   error = errors.New("TenantNotFoundError")
	ErrFunctionNotFound error = errors.New("FunctionNotFoundError")
)

type InMemoryPersistence struct {
	tenants   []*model.Tenant
	functions []*model.Function
}

func (p *InMemoryPersistence) SaveTenant(tenant *model.Tenant) error {
	p.tenants = append(p.tenants, tenant)
	return nil
}

func (p *InMemoryPersistence) GetTenant(id string) (*model.Tenant, error) {
	for _, tenant := range p.tenants {
		if tenant.ID == id {
			return tenant, nil
		}
	}
	return nil, ErrTenantNotFound
}

func (p *InMemoryPersistence) SaveFunction(function *model.Function) error {
	p.functions = append(p.functions, function)
	return nil
}

func (p *InMemoryPersistence) GetFunction(tenantId, id string) (*model.Function, error) {
	for _, function := range p.functions {
		if function.Tenant.ID == tenantId && function.ID == id {
			return function, nil
		}
	}
	return nil, ErrFunctionNotFound
}

func (p *InMemoryPersistence) FunctionExists(tenantId, functionId string) bool {
	return true
}

func (p *InMemoryPersistence) SaveFunctionExecution(funcExec *model.FunctionExecution) error {
	return nil
}

func (p *InMemoryPersistence) FuncExecExists(id string) bool {
	return true
}
