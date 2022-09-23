package app

import (
	"testing"

	"github.com/rrashidov/srvls/internal/pers"
	"github.com/rrashidov/srvls/internal/pers/inmem"
)

const (
	TestTenantID   string = "test-tenant-id"
	TestTenantName string = "test-tenant-name"

	NonExistinTenantID string = "non-existing=tenant"

	TestFunctionID             string = "test-function-id"
	TestFunctionName           string = "test-function-name"
	TestFunctionContainerImage string = "test-function-container-image"
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

func TestCreateFunction(t *testing.T) {
	inMemPers := createPersistence()

	app := createTestApp(inMemPers)

	app.CreateTenant(TestTenantID, TestTenantName)

	err := app.CreateFunction(TestTenantID, TestFunctionID, TestFunctionName, TestFunctionContainerImage)

	if err != nil {
		t.Errorf("Problem creating function: %s", err)
	}

	functionExists := inMemPers.FunctionExists(TestTenantID, TestFunctionID)

	if !functionExists {
		t.Errorf("Function not saved in persistence layer")
	}
}

func TestCreateFunctionInNonExistingTenant(t *testing.T) {
	inMemPers := createPersistence()

	app := createTestApp(inMemPers)

	err := app.CreateFunction(NonExistinTenantID, TestFunctionID, TestFunctionName, TestFunctionContainerImage)

	if err == nil {
		t.Errorf("Creating function in non-existing tenant should fail: %s", NonExistinTenantID)
	}
}

func TestTriggerFunctionExecutionHappyPath(t *testing.T) {
	inMemPers := createPersistence()

	app := createTestApp(inMemPers)

	app.CreateTenant(TestTenantID, TestTenantName)

	app.CreateFunction(TestTenantID, TestFunctionID, TestFunctionName, TestFunctionContainerImage)

	funcExecId, err := app.TriggerFunction(TestTenantID, TestFunctionID)

	if err != nil {
		t.Errorf("Problem triggering function execution: %v", err.Error())
	}

	if len(funcExecId) == 0 {
		t.Errorf("Triggering function execution did not return execution id")
	}

	funcExecExists := inMemPers.FuncExecExists(funcExecId)

	if !funcExecExists {
		t.Errorf("Triggering function execution should persist functionExecution model entity")
	}
}

func TestTriggerFunctionInNonExistingTenant(t *testing.T) {
	inMemPers := createPersistence()

	app := createTestApp(inMemPers)

	_, err := app.TriggerFunction(NonExistinTenantID, TestFunctionID)

	if err == nil {
		t.Errorf("Triggering function in non existing tenant should fail")
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
