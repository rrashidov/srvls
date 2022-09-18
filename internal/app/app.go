package app

import (
	"github.com/rrashidov/srvls/internal/model"
	"github.com/rrashidov/srvls/internal/pers"
)

type App interface {
	CreateTenant(id, name string) error
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
