package postgres

import (
	"context"
	"test/api/models"
	"test/config"
	"test/pkg/logger"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestBranchRepo_Create(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg, logger.New(""))
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	createbranchs := models.CreateBranch{
		Name:        "scsa",
		Address:     "cscsac",
		PhoneNumber: "+3432432",
	}

	branchId, err := pgStore.Branch().Create(context.Background(), createbranchs)
	if err != nil {
		t.Errorf("error while creating branch error: %v", err)
	}

	branch, err := pgStore.Branch().GetByID(context.Background(), models.PrimaryKey{ID: branchId})
	if err != nil {
		t.Errorf("error while getting branch error: %v", err)
	}

	assert.Equal(t, branch.Name, createbranchs.Name)
	assert.Equal(t, branch.Address, createbranchs.Address)
	assert.Equal(t, branch.PhoneNumber, createbranchs.PhoneNumber)

}
func TestBranchRepo_GetByID(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg, logger.New(""))
	if err != nil {
		t.Fatalf("error while connecting to the database: %v", err)
	}
	defer pgStore.Close()

	createBranch := models.CreateBranch{
		Name:        "ssadsa",
		Address:     "cscsac",
		PhoneNumber: "+345325432",
	}

	branchID, err := pgStore.Branch().Create(context.Background(), createBranch)
	if err != nil {
		t.Fatalf("error while creating branch: %v", err)
	}

	t.Run("success", func(t *testing.T) {
		retrievedBranch, err := pgStore.Branch().GetByID(context.Background(), models.PrimaryKey{ID: branchID})
		if err != nil {
			t.Fatalf("error while getting branch by ID: %v", err)
		}

		if retrievedBranch.ID != branchID {
			t.Errorf("expected branch ID: %q, but got %q", branchID, retrievedBranch.ID)
		}

		if retrievedBranch.Name == "" {
			t.Error("expected non-empty name, but got empty")
		}
		if retrievedBranch.PhoneNumber == "" {
			t.Error("expected non-empty phone number, but got empty")
		} else if len(retrievedBranch.PhoneNumber) != 10 {
			t.Errorf("expected phone length: 10, but got %d, branch ID: %s", len(retrievedBranch.PhoneNumber), retrievedBranch.ID)
		}

	})
}

func TestBranch_GetList(t *testing.T) {

	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg, logger.New(""))
	if err != nil {
		t.Errorf("error while connecting to the database: %v", err)
	}

	request := models.GetListRequest{
		Page:   1,
		Limit:  10,
		Search: "",
	}
	response, err := pgStore.Branch().GetList(context.Background(), request)
	if err != nil {
		t.Errorf("error while getting branch list: %v", err)
	}

	if response.Count < 0 {
		t.Errorf("unexpected count: %d", response.Count)
	}
	if len(response.Branches) < 0 {
		t.Errorf("unexpected number of branches: %d", len(response.Branches))
	}

}
func TestBranchRepo_Update(t *testing.T) {

	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg, logger.New(""))
	if err != nil {
		t.Fatalf("error while connecting to the database: %v", err)
	}

	createBranch := models.CreateBranch{
		Name:        "Test xsacxasc",
		Address:     "Test Awddscsacsaress",
		PhoneNumber: "+32432432",
	}
	branchID, err := pgStore.Branch().Create(context.Background(), createBranch)
	if err != nil {
		t.Fatalf("error while creating test branch: %v", err)
	}

	updateBranch := models.UpdateBranch{
		ID:          branchID,
		Name:        "2421421 Name",
		Address:     "dffcds Address",
		PhoneNumber: "+9219321",
	}

	updatedID, err := pgStore.Branch().Update(context.Background(), updateBranch)

	assert.Equal(t, updateBranch.ID, updatedID)

	updatedBranch, err := pgStore.Branch().GetByID(context.Background(), models.PrimaryKey{ID: branchID})
	if err != nil {
		t.Fatalf("error while fetching updated branch: %v", err)
	}

	assert.Equal(t, updateBranch.Name, updatedBranch.Name)
	assert.Equal(t, updateBranch.Address, updatedBranch.Address)
	assert.Equal(t, updateBranch.PhoneNumber, updatedBranch.PhoneNumber)
}

func TestBranchRepo_Delete(t *testing.T) {

	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg, logger.New(""))
	if err != nil {
		t.Fatalf("error while connecting to the database: %v", err)
	}

	createBranch := models.CreateBranch{
		Name:        "Test Branch",
		Address:     "Test Address",
		PhoneNumber: "1234567890",
	}
	branchID, err := pgStore.Branch().Create(context.Background(), createBranch)
	if err != nil {
		t.Fatalf("error while creating test branch: %v", err)
	}

	err = pgStore.Branch().Delete(context.Background(), models.PrimaryKey{ID: branchID})

	_, err = pgStore.Branch().GetByID(context.Background(), models.PrimaryKey{ID: branchID})

}
