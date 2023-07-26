package product

import (
	"base-gin-golang/domain/entity"

	"github.com/jinzhu/copier"
)

type UpdateProductInput struct {
	ID          int64
	ProductCode string `json:"productCode" binding:"required"`
	ProductName string `json:"productName" binding:"required"`
	Price       int    `json:"price" binding:"required"`
}

func (pu *productUseCase) Update(input *UpdateProductInput) (*entity.Product, error) {
	data := &entity.Product{}
	err := copier.Copy(data, input)
	if err != nil {
		return nil, err
	}
	newProduct, err := pu.productRepository.Update(input.ID, data)
	if err != nil {
		return nil, err
	}
	return newProduct, nil
}
