package app

import (
	"testing"

	"github.com/rrashidov/srvls/internal/model"
	"github.com/rrashidov/srvls/internal/pers"
	"github.com/rrashidov/srvls/internal/pers/inmem"
)

const (
	test_tenant_id string = "test-tenant-id"
)

func TestCreateTenant(t *testing.T) {

	in_mem_pers := create_persistence()

	app := create_test_app(in_mem_pers)

	tenant := create_test_tenant()

	err := app.CreateTenant(tenant)

	if err != nil {
		t.Error("Problem creating tenant", err)
	}

	tenant_exists := in_mem_pers.TenantExists(test_tenant_id)

	if !tenant_exists {
		t.Error("Created tenant not found in persistence")
	}
}

func create_test_app(p pers.Persistence) *Application {
	return &Application{
		p: p,
	}
}

func create_test_tenant() *model.Tenant {
	return &model.Tenant{
		ID:   test_tenant_id,
		Name: "test-name",
	}
}

func create_persistence() *inmem.InMemoryPersistence {
	return &inmem.InMemoryPersistence{}
}
