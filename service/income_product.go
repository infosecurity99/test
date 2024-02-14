package service

import (
	"context"
	"test/api/models"
	"test/pkg/logger"
	"test/storage"
)

type incomeProductService struct {
	storage storage.IStorage
	log     logger.ILogger
}

func NewIncomeProductService(storage storage.IStorage, log logger.ILogger) incomeProductService {
	return incomeProductService{
		storage: storage,
		log:     log,
	}
}

func (i incomeProductService) CreateMultiple(ctx context.Context, request models.CreateIncomeProducts) error {
	if err := i.storage.IncomeProduct().CreateMultiple(ctx, request); err != nil {
		i.log.Error("error while creating multiple income products", logger.Error(err))

		return err
	}

	return nil
}

func (i incomeProductService) GetList(ctx context.Context, request models.GetListRequest) (models.IncomeProductsResponse, error) {
	incomeProducts, err := i.storage.IncomeProduct().GetList(ctx, request)
	if err != nil {
		i.log.Error("error in service layer while getting list", logger.Error(err))

		return models.IncomeProductsResponse{}, err
	}
	return incomeProducts, nil
}

func (i incomeProductService) UpdateMultiple(ctx context.Context, response models.UpdateIncomeProducts) error {
	if err := i.storage.IncomeProduct().UpdateMultiple(ctx, response); err != nil {
		i.log.Error("error in service layer while updating", logger.Error(err))

		return err
	}

	return nil
}

func (i incomeProductService) DeleteMultiple(ctx context.Context, response models.DeleteIncomeProducts) error {
	err := i.storage.IncomeProduct().DeleteMultiple(ctx, response)
	return err
}
