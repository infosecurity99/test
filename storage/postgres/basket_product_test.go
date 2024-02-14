package postgres

import (
	"context"
	"test/api/models"
	"test/config"
	"test/pkg/logger"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestBasketProductRepo_Create(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg, logger.New(""))
	if err != nil {
		t.Errorf("error while connecting to db: %v", err)
		return
	}
	defer pgStore.Close()

	createBasketProduct := models.CreateBasketProduct{
		BasketID:  "9b2c0728-ab31-4ec0-aa5d-864084491cdb",
		ProductID: "cc894270-9c85-4ad4-8e87-dcc540a483b3",
		Quantity:  12,
	}

	basketProductID, err := pgStore.BasketProduct().Create(context.Background(), createBasketProduct)
	if err != nil {
		t.Errorf("error while creating basket product: %v", err)
		return
	}

	basketProduct, err := pgStore.BasketProduct().GetByID(context.Background(), models.PrimaryKey{ID: basketProductID})
	if err != nil {
		t.Errorf("error while getting basket product by ID: %v", err)
		return
	}

	assert.Equal(t, basketProduct.BasketID, createBasketProduct.BasketID)
	assert.Equal(t, basketProduct.ProductID, createBasketProduct.ProductID)
	assert.Equal(t, basketProduct.Quantity, createBasketProduct.Quantity)
}

func TestBaketProduct_GetById(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg, logger.New(""))
	if err != nil {
		t.Errorf("error whiling  to db connect %v", err)
	}

	createBasketProduct := models.CreateBasketProduct{
		BasketID:  "9b2c0728-ab31-4ec0-aa5d-864084491cdb",
		ProductID: "cc894270-9c85-4ad4-8e87-dcc540a483b3",
		Quantity:  12,
	}

	basketProductID, err := pgStore.BasketProduct().Create(context.Background(), createBasketProduct)
	if err != nil {
		t.Errorf("error while creating basket product: %v", err)
		return
	}

	t.Run("success", func(t *testing.T) {
		basketProduct, err := pgStore.BasketProduct().GetByID(context.Background(), models.PrimaryKey{ID: basketProductID})
		if err != nil {
			t.Errorf("error while getting basket product by ID: %v", err)
			return
		}

		if basketProduct.ID != basketProductID {
			t.Errorf("expected: %q, but got %q", basketProductID, basketProduct.ID)
		}

		if basketProduct.Quantity < 0 {
			t.Errorf("expected > 0, but got %d", basketProduct.Quantity)
		}

		if basketProduct.BasketID == "" {
			t.Error("expected some basket id, but got nothing")
		}
		if basketProduct.ProductID == "" {
			t.Error("expected some product id, but got nothing")
		}
	})
}

func TestBasketProduct_GetList(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg, logger.New(""))
	if err != nil {
		t.Errorf("error while connecting to db error %v", err)
	}

	basketproductResp, err := pgStore.BasketProduct().GetList(context.Background(), models.GetListRequest{
		Page:  1,
		Limit: 100,
	})
	if err != nil {
		t.Errorf("error while getting basketproduct error %v", err)
	}

	if len(basketproductResp.BasketProducts) != 2 {
		t.Errorf("expected 2, but got: %d", len(basketproductResp.BasketProducts))

	}
	assert.Equal(t, len(basketproductResp.BasketProducts), 2)

}

func TestBasketProduct_Update(t *testing.T) {

	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg, logger.New(""))
	if err != nil {
		t.Errorf("error while connecting to db: %v", err)
		return
	}
	defer pgStore.Close()

	createBasketProduct := models.CreateBasketProduct{
		BasketID:  "9b2c0728-ab31-4ec0-aa5d-864084491cdb",
		ProductID: "cc894270-9c85-4ad4-8e87-dcc540a483b3",
		Quantity:  12,
	}

	basketProductID, err := pgStore.BasketProduct().Create(context.Background(), createBasketProduct)
	if err != nil {
		t.Errorf("error while creating basket product: %v", err)
		return
	}

	updateBasketProduct := models.UpdateBasketProduct{
		ID:        basketProductID,
		ProductID: "cc894270-9c85-4ad4-8e87-dcc540a483b3",
		Quantity:  20,
	}

	updatedID, err := pgStore.BasketProduct().Update(context.Background(), updateBasketProduct)
	if err != nil {
		t.Errorf("error while updating basket product: %v", err)
		return
	}

	updatedProduct, err := pgStore.BasketProduct().GetByID(context.Background(), models.PrimaryKey{ID: updatedID})
	if err != nil {
		t.Errorf("error while getting basket product by ID: %v", err)
		return
	}

	assert.Equal(t, updateBasketProduct.ProductID, updatedProduct.ProductID)
	assert.Equal(t, updateBasketProduct.Quantity, updatedProduct.Quantity)
}

func TestBasketProduct_Delete(t *testing.T) {

	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg, logger.New(""))
	if err != nil {
		t.Errorf("error while connecting to db: %v", err)
		return
	}
	defer pgStore.Close()

	createBasketProduct := models.CreateBasketProduct{
		BasketID:  "9b2c0728-ab31-4ec0-aa5d-864084491cdb",
		ProductID: "cc894270-9c85-4ad4-8e87-dcc540a483b3",
		Quantity:  12,
	}

	basketProductID, err := pgStore.BasketProduct().Create(context.Background(), createBasketProduct)
	if err != nil {
		t.Errorf("error while creating basket product: %v", err)
		return
	}

	if err = pgStore.Basket().Delete(context.Background(), models.PrimaryKey{ID: basketProductID}); err != nil {
		t.Errorf("Error deleting basketproduct: %v", err)
	}
}
