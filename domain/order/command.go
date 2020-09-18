package domain_order

import (
	"errors"
	cqrs_core "goevent/cqrs-core"
	"goevent/database"
	"goevent/helpers"
	"goevent/models"
)

type CreateOrderCommand struct {
	Client string
}

type AddOrderLineCommand struct {
	Price uint
	Meal string
	IDOrder uint
}

type CreateOrderCommandHandler struct {}

func (ch CreateOrderCommandHandler) Handle(command cqrs_core.CommandMessage) error {
	db := database.DbConn
	repository := models.Repository{Conn: db}

	switch cmd := command.Payload().(type) {
	case *CreateOrderCommand:
		order := &models.Order{
			TotalPrice: 230,
			Client: cmd.Client,
			Reference: helpers.RandomString10(),
			Lines: []*models.OrderLine{},
		}
		repository.NewOrder(order)
	case *AddOrderLineCommand:
		orderLine := &models.OrderLine{
			Meal:    cmd.Meal,
			Price:   cmd.Price,
			IDOrder: cmd.IDOrder,
		}
		repository.AddOrderLine(orderLine)
	default:
		return errors.New("bad command type")
	}

	return nil
}

func NewCreateOrderCommandHandler() *CreateOrderCommandHandler {
	return &CreateOrderCommandHandler{}
}


