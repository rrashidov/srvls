package app

import (
	"github.com/rrashidov/srvls/internal/model"
	"github.com/rrashidov/srvls/internal/pers"
)

type App interface {
	CreateTenant(id, name string) error
	CreateFunction(tenantId, id, name, containerImage string) error
}

type Application struct {
	p pers.Persistence
}

func (app Application) CreateTenant(id, name string) error {
	t := &model.Tenant{
		ID:   id,
		Name: name,
	}
	return app.p.SaveTenant(t)
}

func (app Application) CreateFunction(tenantId, id, name, containerImager string) error {
	tenant, err := app.p.GetTenant(tenantId)

	if err != nil {
		return err
	}

	function := &model.Function{
		ID:             id,
		Tenant:         *tenant,
		Name:           name,
		ContainerImage: containerImager,
	}

	return app.p.SaveFunction(function)
}
