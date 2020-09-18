package models

import (
	"database/sql"
	"fmt"
	"time"
)

// Repository struct for db connection
type Repository struct {
	Conn *sql.DB
}

func (repository *Repository) GetOrders() ([]*Order, error) {
	rows, err := repository.Conn.Query("SELECT fo.id, fo.reference, SUM(fol.price) AS total, fo.client, fo.date FROM FoodOrder fo " +
	"\nJOIN FoodOrderLine fol on fo.id = fol.id_food_order " +
	"\nGROUP BY fo.id")

	if err != nil {
		return nil, fmt.Errorf("could not prepare query: %v", err)
	}
	var orders []*Order
	var id, total uint
	var reference, client string
	var date time.Time
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &reference, &total, &client, &date)
		if err == sql.ErrNoRows {
			return nil, nil
		}
		if err != nil {
			return nil, fmt.Errorf("could not get articles : %v", err)
		}
		order := &Order{
			ID:        		id,
			Reference:     	reference,
			TotalPrice: 	total,
			Date: 			date,
			Client: 		client,
		}
		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		fmt.Print(err)
	}

	return orders, nil
}

func (repository *Repository) NewOrder(order *Order) error{
	stmt, err := repository.Conn.Prepare("INSERT INTO FoodOrder(reference, client, date) VALUES(?,?,?)")
	if err != nil {
		return err
	}

	order.Date = time.Now()

	res, errExec := stmt.Exec(order.Reference, order.Client, order.Date)
	if errExec != nil {
		return fmt.Errorf("could not exec stmt: %v", errExec)
	}

	lastInsertedID, errInsert := res.LastInsertId()

	if errInsert != nil {
		return fmt.Errorf("could not retrieve last inserted ID: %v", errInsert)
	}

	order.ID = uint(lastInsertedID)

	return nil
}

func (repository *Repository) AddOrderLine(line *OrderLine) error {
	return nil
}