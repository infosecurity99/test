package postgres

import (
	"context"
	"test/api/models"
	"test/config"
	"test/pkg/logger"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestProductRepo_Create(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg, logger.New(""))
	if err != nil {
		t.Errorf("ERRROR  WHILE CONNECTING TO DB ERROR %v", err)
	}

	createProduct := models.CreateProduct{
		Name:          "apple",
		Price:         100,
		OriginalPrice: 2000,
		Quantity:      34,
		CategoryID:    "0b59dd69-b7b3-43c7-95c1-19a1fd9e0677",
		BranchID:      "aa541fcc-bf74-11ee-ae0b-166244b65504",
	}

	createProductId, err := pgStore.Product().Create(context.Background(), createProduct)
	if err != nil {
		t.Errorf("error while creating product error %v", err)
	}

	product, err := pgStore.Product().GetByID(context.Background(), models.PrimaryKey{ID: createProductId})
	if err != nil {
		t.Errorf("error while getting product error %v", err)
	}

	assert.Equal(t, product.Name, createProduct.Name)
	assert.Equal(t, product.Price, createProduct.Price)
	assert.Equal(t, product.OriginalPrice, createProduct.OriginalPrice)
	assert.Equal(t, product.Quantity, createProduct.Quantity)
	assert.Equal(t, product.BranchID, createProduct.BranchID)
	assert.Equal(t, product.CategoryID, createProduct.CategoryID)
}

func TestProductRepo_GetById(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg, logger.New(""))
	if err != nil {
		t.Fatalf("Error while connecting to database: %v", err)
	}

	createProduct := models.CreateProduct{
		Name:          "apple",
		Price:         100,
		OriginalPrice: 2000,
		Quantity:      34,
		CategoryID:    "0b59dd69-b7b3-43c7-95c1-19a1fd9e0677",
		BranchID:      "aa541fcc-bf74-11ee-ae0b-166244b65504",
	}

	productId, err := pgStore.Product().Create(context.Background(), createProduct)
	if err != nil {
		t.Fatalf("Error while creating product: %v", err)
	}

	t.Run("success", func(t *testing.T) {
		product, err := pgStore.Product().GetByID(context.Background(), models.PrimaryKey{ID: productId})
		if err != nil {
			t.Fatalf("Error while getting product by ID: %v", err)
		}

		if product.ID != productId {
			t.Errorf("Expected product ID: %q, got: %q", productId, product.ID)
		}

		if product.Name == "" {
			t.Error("Expected a non-empty product name, got nothing")
		}

		if product.BranchID == "" {
			t.Error("Expected a non-empty branch ID, got nothing")
		}

		if product.CategoryID == "" {
			t.Error("Expected a non-empty category ID, got nothing")
		}
		if product.Quantity < 0 {
			t.Errorf("expected > 0, but got %d", product.Quantity)
		}
		if product.Price < 0 {
			t.Errorf("expected > 0, but got %d", product.Price)
		}

		if product.OriginalPrice < 0 {
			t.Errorf("expected > 0, but got %d", product.OriginalPrice)
		}

	})
}

func TestProductRepo_GetList(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg, logger.New(""))
	if err != nil {
		t.Errorf("error while  connectiong  to db error %v", err)
	}
	productRepo, err := pgStore.Product().GetList(context.Background(), models.GetListRequest{
		Page:  1,
		Limit: 100,
	})
	if err != nil {
		t.Errorf("erroring while getting productresp error %v", err)
	}

	if len(productRepo.Products) != 5 {
		t.Errorf("expected  5, but got :%d", len(productRepo.Products))
	}

	assert.Equal(t, len(productRepo.Products), 5)
}

func TestProductRepo_Update(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg, logger.New(""))
	if err != nil {
		t.Errorf("erroring while  db connect problamem %v", err)
	}

	createProduct := models.CreateProduct{
		Name:          "apple",
		Price:         100,
		OriginalPrice: 2000,
		Quantity:      34,
		CategoryID:    "0b59dd69-b7b3-43c7-95c1-19a1fd9e0677",
		BranchID:      "aa541fcc-bf74-11ee-ae0b-166244b65504",
	}

	productid, err := pgStore.Product().Create(context.Background(), createProduct)
	updateProduct := models.UpdateProduct{
		ID:            productid,
		Name:          "apple",
		Price:         100,
		OriginalPrice: 2000,
		Quantity:      34,
		CategoryID:    "0b59dd69-b7b3-43c7-95c1-19a1fd9e0677",
	}

	updateProductId, err := pgStore.Product().Update(context.Background(), updateProduct)
	if err != nil {
		t.Errorf("erroring while updatated product id %v", err)
	}

	product, err := pgStore.Product().GetByID(context.Background(), models.PrimaryKey{ID: updateProductId})
	if err != nil {
		t.Errorf("erroring while  updateproduct errror %v", err)
	}

	assert.Equal(t, updateProductId, product.ID)
	assert.Equal(t, product.Name, updateProduct.Name)
	assert.Equal(t, product.Price, updateProduct.Price)
	assert.Equal(t, product.Price, updateProduct.Price)
	assert.Equal(t, product.Quantity, updateProduct.Quantity)
	assert.Equal(t, product.OriginalPrice, updateProduct.OriginalPrice)
}

func TestProductRepo_Delete(t *testing.T) {
	cfg := config.Load()

	pgStore, err := New(context.Background(), cfg, logger.New(""))
	if err != nil {
		t.Fatalf("Error while connecting to database: %v", err)
	}

	createProduct := models.CreateProduct{
		Name:          "apple",
		Price:         100,
		OriginalPrice: 2000,
		Quantity:      34,
		CategoryID:    "0b59dd69-b7b3-43c7-95c1-19a1fd9e0677",
		BranchID:      "aa541fcc-bf74-11ee-ae0b-166244b65504",
	}

	productId, err := pgStore.Product().Create(context.Background(), createProduct)
	if err != nil {
		t.Fatalf("Error while creating product: %v", err)
	}

	if err = pgStore.Product().Delete(context.Background(), models.PrimaryKey{ID: productId}); err != nil {
		t.Errorf("error delating product %v", err)
	}

}
