package products

import "poke-store/internal/domain/products"

type ProductsRepository interface {
	Save(product products.Product) (products.Entity, error)
}
