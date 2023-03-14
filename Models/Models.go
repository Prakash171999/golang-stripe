package Models

type Products struct {
	Id          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func (b *Products) TableName() string {
	return "products"
}
