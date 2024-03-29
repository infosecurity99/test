package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type basketProductService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewBasketProductService(storage storage.IStorage, log logger.ILogger) basketProductService {
	return basketProductService{
		storage: storage,
		log:     log,
	}
}

func (b basketProductService) Create(ctx context.Context, createProduct models.CreateBasketProduct) (models.BasketProduct, error) {
	id, err := b.storage.BasketProduct().Create(ctx, createProduct)
	if err != nil {
		b.log.Error("error in service layer while creating basket product", logger.Error(err))

		return models.BasketProduct{}, err
	}

	createdProduct, err := b.storage.BasketProduct().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		b.log.Error("error in service layer is while getting by id", logger.Error(err))

		return models.BasketProduct{}, err
	}

	return createdProduct, nil
}

func (b basketProductService) Get(ctx context.Context, key models.PrimaryKey) (models.BasketProduct, error) {
	basketProduct, err := b.storage.BasketProduct().GetByID(ctx, key)
	if err != nil {
		b.log.Error("error in service layer is while getting basket product", logger.Error(err))

		return models.BasketProduct{}, err
	}

	return basketProduct, nil
}

func (b basketProductService) GetList(ctx context.Context, request models.GetListRequest) (models.BasketProductResponse, error) {
	basketProducts, err := b.storage.BasketProduct().GetList(ctx, request)
	if err != nil {
		b.log.Error("error in service layer while getting list", logger.Error(err))

		return models.BasketProductResponse{}, err
	}

	return basketProducts, nil
}

func (b basketProductService) Update(ctx context.Context, product models.UpdateBasketProduct) (models.BasketProduct, error) {
	id, err := b.storage.BasketProduct().Update(ctx, product)
	if err != nil {
		b.log.Error("error in service layer while getting list", logger.Error(err))

		return models.BasketProduct{}, err
	}

	updatedBasketProduct, err := b.storage.BasketProduct().GetByID(ctx, models.PrimaryKey{ID: id})
	if err != nil {
		b.log.Error("error in service layer while getting by id", logger.Error(err))

		return models.BasketProduct{}, err
	}

	return updatedBasketProduct, nil
}

func (b basketProductService) Delete(ctx context.Context, key models.PrimaryKey) error {
	err := b.storage.BasketProduct().Delete(ctx, key)
	return err
}
