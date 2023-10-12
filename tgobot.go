package TGoBot

import (
	"context"
	"errors"
	"log/slog"

	command_dto "github.com/Avirtan/TGoBot/dto/command"
	update_dto "github.com/Avirtan/TGoBot/dto/update"
	"github.com/Avirtan/TGoBot/handler"
	"github.com/Avirtan/TGoBot/method"
	"github.com/Avirtan/TGoBot/pkg/logger"
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

type TGoBot struct {
	token    string
	offset   uint64
	notify   chan error
	ctx      context.Context
	handlers []BotHandler
	commands map[*command_dto.BotCommand]BotHandler
}

func NewBot(option BotOptions) *TGoBot {
	logger.New(option.LoggerLevel.String())
	method.New(option.Token)
	return &TGoBot{
		token:    option.Token,
		ctx:      option.Ctx,
		notify:   make(chan error, 1),
		commands: make(map[*command_dto.BotCommand]BotHandler),
	}
}

func (t *TGoBot) AddHandler(handler handler.IHandler) {
	t.handlers = append(t.handlers, BotHandler{
		Handler: handler,
	})
}

func (t *TGoBot) AddHandlerWithMiddleware(handler handler.IHandler, middleware ...handler.Middleware) {
	t.handlers = append(t.handlers, BotHandler{
		Handler:    handler,
		Middleware: middleware,
	})
}

func (t *TGoBot) AddCommand(botCommand *command_dto.BotCommand, handler handler.IHandler) {
	if botCommand.Command[0] != '/' {
		t.notify <- errors.New("command must start with /")
		return
	}
	t.commands[botCommand] = BotHandler{
		Handler: handler,
	}
}

func (t *TGoBot) AddCommandWithMiddleware(botCommand *command_dto.BotCommand, handler handler.IHandler, middleware ...handler.Middleware) {
	if botCommand.Command[0] != '/' {
		t.notify <- errors.New("command must start with /")
		return
	}
	t.commands[botCommand] = BotHandler{
		Handler:    handler,
		Middleware: middleware,
	}
}

func (t *TGoBot) GetCommand() {
	err := method.GetMyCommands(t.ctx)
	if err != nil {
		slog.Error("GetCommand", "error", err.Error())
	}
}

func (t *TGoBot) DeleteCommand() {
	method.DeleteMyCommands(t.ctx)
}

func (t *TGoBot) RunUpdate() {
	commands := make([]*command_dto.BotCommand, 0, 10)
	for command := range t.commands {
		commands = append(commands, command)
	}
	if len(commands) > 0 {
		method.SetMyCommands(t.ctx, commands)
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
					if value.Message != nil {
						for key, handler := range t.commands {
							if key.Command == value.Message.Text {
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
					}
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

func (t *TGoBot) Notify() <-chan error {
	return t.notify
}

func (t *TGoBot) GetTypeMessage() {
}
