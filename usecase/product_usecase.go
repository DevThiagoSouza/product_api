package usecase

import (
	"game-shop-api/model"
	"game-shop-api/repository"
)

type ProductUsecase struct {
	repository repository.ProductRespository
}

func NewProductUsecase(repo repository.ProductRespository) ProductUsecase {
	return ProductUsecase{
		repository: repo,
	}
}

func (p *ProductUsecase) GetProducts() ([]model.Procuct, error) {
	return p.repository.GetProducts(), nil
}

func (p *ProductUsecase) CreateProduct(product model.Procuct) (model.Procuct, error) {
	productId, err := p.repository.CreateProduct(product)
	if err != nil {
		return model.Procuct{}, err
	}

	product.ID = productId

	return product, nil
}

func (p *ProductUsecase) GetProductsById(id int) (*model.Procuct, error) {
	product, err := p.repository.GetProductsById(id)
	if err != nil {
		return nil, nil
	}

	return product, nil
}

// func (p *ProductUsecase) DeleteProducts(id int) (model.Procuct, error) {
// 	product, err := p.repository.DeleteProducts(id)
// 	if err != nil {
// 		return model.Procuct{}, nil
// 	}
// 	return product, nil
// }
