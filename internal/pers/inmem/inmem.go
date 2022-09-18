package inmem

import "github.com/rrashidov/srvls/internal/model"

type InMemoryPersistence struct {
	tenants []*model.Tenant
}

func (p *InMemoryPersistence) SaveTenant(tenant *model.Tenant) error {
	p.tenants = append(p.tenants, tenant)
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
