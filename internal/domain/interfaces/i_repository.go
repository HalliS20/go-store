package interfaces

import (
	"go-shop/internal/models"
)

type IProductRepository interface {
	GetProducts() []models.Product
	GetProduct(id int) models.Product
}
