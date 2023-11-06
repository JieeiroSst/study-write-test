package repositories

import "tet/internal/models"

type ProductRepositoryInterface interface {
	Add(product models.Product) error
}
