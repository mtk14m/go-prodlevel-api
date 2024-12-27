package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdateOn    string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// List of products
type Products []*Product

// Get all products
func GetProducts() Products {
	return productList
}


func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Espresso",
		Description: "Strong and bold coffee shot",
		Price:       2.50,
		SKU:         "COF-ESP-01",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
		DeletedOn:   "",
	},
	&Product{
		ID:          2,
		Name:        "Latte",
		Description: "Smooth and creamy coffee with milk",
		Price:       3.50,
		SKU:         "COF-LAT-01",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
		DeletedOn:   "",
	},
	&Product{
		ID:          3,
		Name:        "Cappuccino",
		Description: "Rich coffee with steamed milk foam",
		Price:       3.75,
		SKU:         "COF-CAP-01",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
		DeletedOn:   "",
	},
	&Product{
		ID:          4,
		Name:        "Americano",
		Description: "Espresso with hot water",
		Price:       2.75,
		SKU:         "COF-AME-01",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
		DeletedOn:   "",
	},
	&Product{
		ID:          5,
		Name:        "Mocha",
		Description: "Coffee with chocolate and steamed milk",
		Price:       4.00,
		SKU:         "COF-MOC-01",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
		DeletedOn:   "",
	},
}
