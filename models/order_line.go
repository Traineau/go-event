package models

type OrderLine struct {
	ID 			uint64		`json:"id"`
	Meal		string		`json:"meal"`
	Price 		uint		`json:"price"`
	IDOrder		uint		`json:"id_command"`
}
