package pers

import "github.com/rrashidov/srvls/internal/model"

type Persistence interface {
	SaveTenant(tenant *model.Tenant) error
}
