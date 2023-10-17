package goborg_tg

import (
	"context"
	"errors"
	"log/slog"

	command_dto "github.com/Avirtan/goborg-tg/dto/command"
	update_dto "github.com/Avirtan/goborg-tg/dto/update"
	"github.com/Avirtan/goborg-tg/handler"
	"github.com/Avirtan/goborg-tg/method"
	"github.com/Avirtan/goborg-tg/pkg/logger"
)

type BotOptions struct {
	Token       string
	Ctx         context.Context
	LoggerLevel logger.LevelLogger
}

type BotHandler struct {
	Handler    handler.IHandler
	Middleware []handler.Middleware
}

type GoborgTG struct {
	token    string
	offset   uint64
	notify   chan error
	ctx      context.Context
	handlers []BotHandler
	commands []*command_dto.BotCommand
}

func NewBot(option BotOptions) *GoborgTG {
	logger.New(option.LoggerLevel.String())
	method.New(option.Token)
	return &GoborgTG{
		token:    option.Token,
		ctx:      option.Ctx,
		notify:   make(chan error, 1),
		commands: make([]*command_dto.BotCommand, 0, 5),
	}
}

func (t *GoborgTG) AddHandler(handler handler.IHandler) {
	t.handlers = append(t.handlers, BotHandler{
		Handler: handler,
	})
}

func (t *GoborgTG) AddHandlerWithMiddleware(handler handler.IHandler, middleware ...handler.Middleware) {
	t.handlers = append(t.handlers, BotHandler{
		Handler:    handler,
		Middleware: middleware,
	})
}

func (t *GoborgTG) AddCommand(botCommand *command_dto.BotCommand, handler handler.IHandler) {
	if botCommand.Command[0] != '/' {
		t.notify <- errors.New("command must start with /")
		return
	}
	t.handlers = append(t.handlers, BotHandler{
		Handler: handler,
	})
	t.commands = append(t.commands, botCommand)
}

func (t *GoborgTG) AddCommandWithMiddleware(botCommand *command_dto.BotCommand, handler handler.IHandler, middleware ...handler.Middleware) {
	if botCommand.Command[0] != '/' {
		t.notify <- errors.New("command must start with /")
		return
	}
	t.handlers = append(t.handlers, BotHandler{
		Handler:    handler,
		Middleware: middleware,
	})
	t.commands = append(t.commands, botCommand)
}

func (t *GoborgTG) GetCommand() {
	err := method.GetMyCommands(t.ctx)
	if err != nil {
		slog.Error("GetCommand", "error", err.Error())
	}
}

func (t *GoborgTG) DeleteCommand() {
	method.DeleteMyCommands(t.ctx)
}

func (t *GoborgTG) RunUpdate() {
	if len(t.commands) > 0 {
		method.SetMyCommands(t.ctx, t.commands)
		t.GetCommand()
	}
	for {
		select {
		case <-t.ctx.Done():
			return
		default:
			response, err := method.GetUpdates(t.ctx, t.offset, 100, 50)
			if err != nil {
				t.notify <- err
				return
			}
			go func(context.Context, update_dto.UpdateResponse) {
				for _, value := range response.Update {
					for _, handler := range t.handlers {
						go func(ctx context.Context, update update_dto.Update, handler BotHandler) {
							for _, middleware := range handler.Middleware {
								if !middleware.Check(ctx, &update) {
									return
								}
							}
							handler.Handler.Action(ctx, &update)
						}(t.ctx, value, handler)
					}
				}
			}(t.ctx, *response)
			lenUpdate := len(response.Update)
			if lenUpdate > 0 {
				t.offset = response.Update[lenUpdate-1].UpdateID + 1
			}
		}
	}
}

func (t *GoborgTG) Notify() <-chan error {
	return t.notify
}

func (t *GoborgTG) GetTypeMessage() {
}
