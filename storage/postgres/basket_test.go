package postgres

import (
	"context"
	"test/api/models"
	"test/config"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestBaketRepo_Create(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("errro  whiling %v", err)
	}
	createBasket := models.CreateBasket{
		CustomerID: "c5eebf53-a536-4745-b816-2264af15d61f",
		TotalSum:   1111,
	}

	basketid, err := pgStore.Basket().Create(context.Background(), createBasket)
	if err != nil {
		t.Errorf("error while creating basket error: %v", err)
	}

	basket, err := pgStore.Basket().GetByID(context.Background(), models.PrimaryKey{ID: basketid})
	if err != nil {
		t.Errorf("error while getting basket error: %v", err)
	}

	assert.Equal(t, basket.CustomerID, createBasket.CustomerID)
	assert.Equal(t, basket.TotalSum, createBasket.TotalSum)
}

func TestBasketRepo_GetById(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("errro  whiling %v", err)
	}
	createBasket := models.CreateBasket{
		CustomerID: "c5eebf53-a536-4745-b816-2264af15d61f",
		TotalSum:   1111,
	}

	basketid, err := pgStore.Basket().Create(context.Background(), createBasket)
	if err != nil {
		t.Errorf("error while creating basket error: %v", err)
	}

	t.Run("success", func(t *testing.T) {
		basket, err := pgStore.Basket().GetByID(context.Background(), models.PrimaryKey{ID: basketid})
		if err != nil {
			t.Errorf("error while getting basket error: %v", err)
		}

		if basketid != basket.ID {
			t.Errorf("expected: %q, but got %q", basketid, basket.ID)
		}

		if basket.TotalSum < 0 {
			t.Errorf("expected > 0, but got %d", basket.TotalSum)
		}

	})
}

func TestBaksetRepo_GetList(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("errro  whiling %v", err)
	}

	basketRepo, err := pgStore.Basket().GetList(context.Background(), models.GetListRequest{
		Page:  1,
		Limit: 100,
	})
	if err != nil {
		t.Errorf("error while getting basketRepo error: %v", err)
	}

	if len(basketRepo.Baskets) != 1 {
		t.Errorf("expected 1, but got: %d", len(basketRepo.Baskets))
	}

	assert.Equal(t, len(basketRepo.Baskets), 1)

}

func TestBaksetRepo_Update(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("errro  whiling %v", err)
	}
	createBasket := models.CreateBasket{
		CustomerID: "c5eebf53-a536-4745-b816-2264af15d61f",
		TotalSum:   1111,
	}
	basketid, err := pgStore.Basket().Create(context.Background(), createBasket)
	if err != nil {
		t.Errorf("error while creating basket error: %v", err)
	}

	updateBasket := models.UpdateBasket{
		ID:         basketid,
		CustomerID: "c5eebf53-a536-4745-b816-2264af15d61f",
		TotalSum:   12222,
	}
	updatebasketid, err := pgStore.Basket().Update(context.Background(), updateBasket)
	if err != nil {
		t.Errorf("error while update basket error: %v", err)
	}
	basket, err := pgStore.Basket().GetByID(context.Background(), models.PrimaryKey{
		ID: updatebasketid,
	})
	if err != nil {
		t.Errorf("error while getting basket error: %v", err)
	}
	assert.Equal(t, basketid, basket.ID)
	assert.Equal(t, basket.CustomerID, updateBasket.CustomerID)
	assert.Equal(t, basket.TotalSum, updateBasket.TotalSum)
}

func Test_Delete(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg)
	if err != nil {
		t.Errorf("errro  whiling %v", err)
	}
	createBasket := models.CreateBasket{
		CustomerID: "c5eebf53-a536-4745-b816-2264af15d61f",
		TotalSum:   1111,
	}
	basketid, err := pgStore.Basket().Create(context.Background(), createBasket)
	if err != nil {
		t.Errorf("error while creating basket error: %v", err)
	}

	if err = pgStore.Basket().Delete(context.Background(), models.PrimaryKey{ID: basketid}); err != nil {
		t.Errorf("Error deleting basket: %v", err)
	}
}
