package app

import (
	"testing"

	"github.com/rrashidov/srvls/internal/pers"
	"github.com/rrashidov/srvls/internal/pers/inmem"
)

const (
	TestTenantID   string = "test-tenant-id"
	TestTenantName string = "test-tenant-name"
)

func TestCreateTenant(t *testing.T) {

	inMemPers := createPersistence()

	app := createTestApp(inMemPers)

	err := app.CreateTenant(TestTenantID, TestTenantName)

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

func createPersistence() *inmem.InMemoryPersistence {
	return &inmem.InMemoryPersistence{}
}
