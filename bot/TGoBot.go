package bot

import (
	"TGoBot/dto"
	"TGoBot/handler"
	"TGoBot/method"
	"TGoBot/pkg/logger"
	"context"
	"errors"
	"log/slog"
)

type BotOptions struct {
	Token       string
	Ctx         context.Context
	LoggerLevel logger.LevelLogger
}

type TGoBot struct {
	token         string
	offset        uint64
	notify        chan error
	methodHandler *method.MethodHandler
	ctx           context.Context
	handlers      []handler.IHandler
	commands      map[*dto.BotCommand]handler.IHandler
}

func NewBot(option BotOptions) *TGoBot {
	logger.New(option.LoggerLevel.String())
	return &TGoBot{
		token: option.Token,
		methodHandler: &method.MethodHandler{
			Token: option.Token,
		},
		ctx:      option.Ctx,
		notify:   make(chan error, 1),
		commands: make(map[*dto.BotCommand]handler.IHandler),
	}
}

func (t *TGoBot) GetMethodHandler() *method.MethodHandler {
	return t.methodHandler
}

func (t *TGoBot) AddHandler(handler handler.IHandler) {
	t.handlers = append(t.handlers, handler)
}

func (t *TGoBot) AddCommand(botCommand *dto.BotCommand, handler handler.IHandler) {
	if botCommand.Command[0] != '/' {
		t.notify <- errors.New("command must start is /")
		return
	}
	t.commands[botCommand] = handler
}

func (t *TGoBot) GetCommand() {
	err := t.methodHandler.GetMyCommands()
	if err != nil {
		slog.Error("GetCommand", "error", err.Error())
	}
}

func (t *TGoBot) DeleteCommand() {
	t.methodHandler.DeleteMyCommands()
}

func (t *TGoBot) RunUpdate() {
	commands := make([]*dto.BotCommand, 0, 10)
	for command := range t.commands {
		commands = append(commands, command)
	}
	if len(commands) > 0 {
		t.methodHandler.SetMyCommands(t.ctx, commands)
		t.GetCommand()
	}
	for {
		select {
		case <-t.ctx.Done():
			return
		default:
			response, err := t.methodHandler.GetUpdates(t.ctx, t.offset, 100, 50)
			if err != nil {
				t.notify <- err
				return
			}
			go func(context.Context, dto.UpdateResponse) {
				for _, value := range response.Update {
					if value.Message != nil {
						for key, handler := range t.commands {
							if key.Command == value.Message.Text {
								go handler.Action(t.ctx, &value, t.methodHandler)
							}
						}
					}
					for _, handler := range t.handlers {
						go handler.Action(t.ctx, &value, t.methodHandler)
					}
				}
			}(t.ctx, *response)
			t.offset = response.Update[len(response.Update)-1].UpdateID + 1
		}
	}
}

func (t *TGoBot) Notify() <-chan error {
	return t.notify
}

func (t *TGoBot) GetTypeMessage() {
}
