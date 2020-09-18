package cqrs_core

import "fmt"

type CommandMessage interface {
	Payload() interface{}
	CommandType() string
}

type CommandBus struct {
	handlers map[string]CommandHandler
}

func NewCommandBus() *CommandBus{
	cBus := &CommandBus{
		handlers: make(map[string]CommandHandler),
	}

	return cBus
}

func (b *CommandBus) Dispatch(command CommandMessage) error {
	if handler, ok := b.handlers[command.CommandType()]; ok {
		return handler.Handle(command)
	}

	return fmt.Errorf("the command bus does not have a handler for command of type: %s", command)
}

func (b *CommandBus) RegisterHandler(handler CommandHandler, command interface{}) error {
	typeName := typeOf(command)
	if _, ok := b.handlers[typeName]; ok {
		return fmt.Errorf("duplicate command handler registration with command bus for command of type: %s", typeName)
	}

	b.handlers[typeName] = handler

	return nil
}

type CommandDescriptor struct {
	command interface{}
}

func NewCommandMessage(command interface{}) *CommandDescriptor {
	return &CommandDescriptor{
		command: command,
	}
}

func (c *CommandDescriptor) CommandType() string {
	return typeOf(c.command)
}

// Command returns the actual command payload of the message.
func (c *CommandDescriptor) Payload() interface{} {
	return c.command
}


