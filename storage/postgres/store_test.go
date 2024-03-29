package postgres

import (
	"context"
	"test/config"
	"test/pkg/logger"
	"testing"
	"time"
)

func TestStoreAddProfitAndWithdrawal(t *testing.T) {
	cfg := config.Load()
	pgStore, err := New(context.Background(), cfg, logger.New(""))
	if err != nil {
		t.Fatalf("error while connection to db: %v", err)
	}

	ctx := context.Background()
	profit := float32(400.0)
	branchID := "28288bf9-3ed1-4f4f-92f4-7ab3d5f2959a"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Add profit to the store
	if err := pgStore.Store().AddProfit(ctx, profit, branchID); err != nil {
		t.Fatalf("AddProfit returned an unexpected error: %v", err)
	}

	// Withdraw the profit from the store
	if err := pgStore.Store().WithdrawalDeliveredSum(ctx, profit, branchID); err != nil {
		t.Fatalf("WithdrawalDeliveredSum returned an unexpected error: %v", err)
	}

}

func TestStoreGetStoreBudget(t *testing.T) {

	cfg := config.Load()
	pgStore, err := New(context.Background(), cfg, logger.New(""))
	if err != nil {
		t.Fatalf("error while connecting to db: %v", err)
	}

	ctx := context.Background()
	branchID := "7534edd3-06bf-4b91-a096-2015b21eda02"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	budget, err := pgStore.Store().GetStoreBudget(ctx, branchID)

	if err != nil {
		t.Fatalf("GetStoreBudget returned an unexpected error: %v", err)
	}

	expectedBudget := float32(3000.0)
	if budget != expectedBudget {
		t.Errorf("GetStoreBudget returned unexpected budget. Expected: %f, Got: %f", expectedBudget, budget)
	}
}

func TestWithdrawalDeliveredSum(t *testing.T) {

	cfg := config.Load()
	pgStore, err := New(context.Background(), cfg, logger.New(""))
	if err != nil {
		t.Fatalf("error while connecting to db: %v", err)
	}

	ctx := context.Background()
	branchID := "3f396f70-60ea-4cb1-8eb7-30fe0fc6664a"
	totalSum := float32(500)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	err = pgStore.Store().WithdrawalDeliveredSum(ctx, totalSum, branchID)

	if err != nil {
		t.Fatalf("WithdrawalDeliveredSum returned an unexpected error: %v", err)
	}

}
