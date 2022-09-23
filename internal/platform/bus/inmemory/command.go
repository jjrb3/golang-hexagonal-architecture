package inmemory

import (
	"context"

	"github.com/jjrb3/golang-hexagonal-architecture/kit/command"
)

// CommandBus is an in-memory implementation of the command.bus.
type CommandBus struct {
	handlers map[command.Type]command.Handler
}

// NewCommandBus initializes a new instance of CommandBus.
func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[command.Type]command.Handler),
	}
}

// Dispatch implements the command.Bus interface.
func (b *CommandBus) Dispatch(ctx context.Context, cmd command.Command) error {
	handler, ok := b.handlers[cmd.Type()]
	if !ok {
		return nil
	}

	return handler.Handle(ctx, cmd)
}

// Register implements the command.
func (b *CommandBus) Register(cmdType command.Type, handler command.Handler) {
	b.handlers[cmdType] = handler
}
