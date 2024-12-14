package entities

type ProductEntity struct {
	ProductID   string  `json:"PID"`
	Product     string  `json:"product"`
	Description string  `json:"description"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
}
