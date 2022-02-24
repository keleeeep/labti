/*
 * @Author: Adrian Faisal
 * @Date: 04/09/21 1.12 PM
 */

package product

import (
	"context"
	"fmt"
	"github.com/apldex/workshop-labti/internal/pkg/model"
	"github.com/apldex/workshop-labti/internal/pkg/resource/db"
)

type Usecase interface {
	CreateProduct(ctx context.Context, product *model.Product) error
	GetProduct(ctx context.Context, id int) (*model.Product, error)
}

type usecase struct {
	dbResource db.Persistent
}

func NewUsecase(dbResource db.Persistent) Usecase {
	return &usecase{dbResource: dbResource}
}

func (uc *usecase) CreateProduct(ctx context.Context, product *model.Product) error {
	if product.Stock == 0 {
		return fmt.Errorf("stock product cannot be 0")
	}

	if len(product.Name) < 3 {
		return fmt.Errorf("minimum product name is 3 characters")
	}

	err := uc.dbResource.CreateProduct(ctx, product)
	if err != nil {
		return fmt.Errorf("create product failed: %v", err)
	}

	return nil
}

func (uc *usecase) GetProduct(ctx context.Context, id int) (*model.Product, error) {
	product, err := uc.dbResource.GetProduct(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get product: %v", err)
	}

	return product, nil
}
