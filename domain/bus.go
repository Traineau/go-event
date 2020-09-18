package domain

import (
	cqrs_core "goevent/cqrs-core"
	domain_order "goevent/domain/order"
)

var CommandBus *cqrs_core.CommandBus
var QueryBus *cqrs_core.QueryBus

func InitBusses(){
	CommandBus = cqrs_core.NewCommandBus()
	QueryBus = cqrs_core.NewQueryBus()

	_ = CommandBus.RegisterHandler(domain_order.NewCreateOrderCommandHandler(), &domain_order.CreateOrderCommand{})
}
