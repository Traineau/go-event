package models

import "time"

type Order struct {
	ID 			uint			`json:"id"`
	Reference 	string			`json:"reference"`
	Client 		string			`json:"client"`
	TotalPrice	uint			`json:"-"`
	Date		time.Time		`json:"time"`
	Lines		[]*OrderLine	`json:"-"`
}
