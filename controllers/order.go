package controllers

import (
	"goevent/cqrs-core"
	"goevent/database"
	"goevent/domain"
	"goevent/domain/order"
	"goevent/helpers"
	"goevent/models"
	"log"
	"net/http"
)

func GetOrders(w http.ResponseWriter, r *http.Request) {
	db := database.DbConn
	repository := models.Repository{Conn: db}

	commands, err := repository.GetOrders()
	if err != nil {
		log.Printf("could not get command list: %v", err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, commands)
}

func CreateOrder(w http.ResponseWriter, r *http.Request){
	order := models.Order{}

	if err := helpers.ReadJSON(w, r, &order); err != nil{
		log.Printf("coucou c'est l'erreur : %v", err)
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "can not parse JSON body")
		return
	}

	command := cqrs_core.NewCommandMessage(&domain_order.CreateOrderCommand{Client: order.Client})

	err := domain.CommandBus.Dispatch(command)

	if err != nil {
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusOK, command)
}

func AddOrderLine(w http.ResponseWriter, r *http.Request){
	command := cqrs_core.NewCommandMessage(&domain_order.AddOrderLineCommand{
		Price:   18,
		Meal:    "Steak",
		IDOrder: 2,
	})

	_ = domain.CommandBus.Dispatch(command)

	w.WriteHeader(http.StatusOK)
}