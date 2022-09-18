package app

import (
	"testing"

	"github.com/rrashidov/srvls/internal/model"
	"github.com/rrashidov/srvls/internal/pers"
	"github.com/rrashidov/srvls/internal/pers/inmem"
)

const (
	TestTenantID string = "test-tenant-id"
)

func TestCreateTenant(t *testing.T) {

	inMemPers := createPersistence()

	app := createTestApp(inMemPers)

	tenant := createTestTenant()

	err := app.CreateTenant(tenant)

	if err != nil {
		t.Error("Problem creating tenant", err)
	}

	tenantExists := inMemPers.TenantExists(TestTenantID)

	if !tenantExists {
		t.Errorf("Created tenant %s not found in persistence", TestTenantID)
	}
}

func createTestApp(p pers.Persistence) *Application {
	return &Application{
		p: p,
	}
}

func createTestTenant() *model.Tenant {
	return &model.Tenant{
		ID:   TestTenantID,
		Name: "test-name",
	}
}

func createPersistence() *inmem.InMemoryPersistence {
	return &inmem.InMemoryPersistence{}
}
