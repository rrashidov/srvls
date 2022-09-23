package inmem

import (
	"errors"

	"github.com/rrashidov/srvls/internal/model"
)

var (
	ErrTenantNotFound error = errors.New("TenantNotFoundError")
)

type InMemoryPersistence struct {
	tenants []*model.Tenant
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
	return nil
}

func (p *InMemoryPersistence) TenantExists(id string) bool {
	for _, tenant := range p.tenants {
		if tenant.ID == id {
			return true
		}
	}
	return false
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
