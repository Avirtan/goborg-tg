package bot

import (
	"TGoBot/internal/dto"
	"TGoBot/internal/handler"
	"TGoBot/internal/method"
	"context"
	"errors"
)

type BotOptions struct {
	Token string
	Ctx   context.Context
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
	t.methodHandler.SetMyCommands(t.ctx, botCommand)
}

func (t *TGoBot) GetCommand() {
	t.methodHandler.GetMyCommands()
}

func (t *TGoBot) DeleteCommand() {
	t.methodHandler.DeleteMyCommands()
}

func (t *TGoBot) RunUpdate() {
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

			for _, value := range response.Update {
				if value.Message != nil {
					for key, handler := range t.commands {
						if key.Command == value.Message.Text {
							go handler.Action(&value, t.methodHandler)
						}
					}
				}
				for _, handler := range t.handlers {
					go handler.Action(&value, t.methodHandler)
				}
				t.offset = value.UpdateID + 1
			}
		}
	}
}

func (t *TGoBot) Notify() <-chan error {
	return t.notify
}

func (t *TGoBot) GetTypeMessage() {
}
