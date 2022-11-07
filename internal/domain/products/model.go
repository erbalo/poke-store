package products

import "time"

type Request struct {
	Sku   string
	Name  string
	Price float32
}

type Product struct {
	ID        int64
	Sku       string
	Name      string
	Price     float32
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Entity struct {
	ID        int64
	Sku       string
	Name      string
	Price     float32
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
