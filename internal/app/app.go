package app

import (
	"github.com/rrashidov/srvls/internal/model"
	"github.com/rrashidov/srvls/internal/pers"
)

type App interface {
	CreateTenant(tenant model.Tenant) error
}

type Application struct {
	p pers.Persistence
}

func (app Application) CreateTenant(tenant *model.Tenant) error {
	return app.p.SaveTenant(tenant)
}
