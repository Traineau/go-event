package cqrs_core

type CommandHandler interface {
	Handle(message CommandMessage) error
}

type QueryHandler interface {
	Handle(message QueryMessage) error
}
