package postgres

import (
	"context"
	"test/api/models"
	"test/config"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestBranchRepo_Create(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connection to db error: %v", err)
	}

	createbranchs := models.CreateBranch{
		Name:        "qqqq",
		Address:     "eeee",
		PhoneNumber: "+33333",
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

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("error while connecting to the database: %v", err)
	}

	createBranch := models.CreateBranch{
		Name:        "Sample Branch",
		Address:     "123 Sample St",
		PhoneNumber: "+1234567890",
	}

	branchID, err := pgStore.Branch().Create(context.Background(), createBranch)
	if err != nil {
		t.Errorf("error while creating branch: %v", err)
	}

	t.Run("success", func(t *testing.T) {
		retrievedBranch, err := pgStore.Branch().GetByID(context.Background(), models.PrimaryKey{ID: branchID})
		if err != nil {
			t.Errorf("error while getting branch by ID: %v", err)
		}
		if retrievedBranch.ID != branchID {
			t.Errorf("expected: %q, but got %q", retrievedBranch, retrievedBranch.ID)
		}

		if retrievedBranch.Name == "" {
			t.Error("expected some full name, but got nothing")
		}

		if retrievedBranch.PhoneNumber == "" {
			t.Error("expected some full name, but got nothing")
		} else if len(retrievedBranch.PhoneNumber) >= 14 || len(retrievedBranch.PhoneNumber) <= 12 {
			t.Errorf("expected phone length: 13, but got %d, user id is %s", len(retrievedBranch.PhoneNumber), retrievedBranch.ID)
		}
	})
}


 


