package main

import (
	goborg_tg "github.com/Avirtan/goborg-tg"
	command_dto "github.com/Avirtan/goborg-tg/dto/command"
	method_dto "github.com/Avirtan/goborg-tg/dto/method"
	update_dto "github.com/Avirtan/goborg-tg/dto/update"
	"github.com/Avirtan/goborg-tg/method"
	"github.com/Avirtan/goborg-tg/pkg/logger"

	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	ctxWithCancel, cancelFunction := context.WithCancel(ctx)
	tbot := goborg_tg.NewBot(goborg_tg.BotOptions{
		Token:       "6584315900:AAHsf2CpLAsnuyl5H1CyB-gOy4wv8S9TIMI",
		Ctx:         ctxWithCancel,
		LoggerLevel: logger.LevelDebug,
	})
	cmd := &TestCommandHandler{}
	tbot.AddHandlerWithMiddleware(cmd, &MiddlwareTest{}, &MiddlwareTest1{})
	tbot.AddCommand(
		&command_dto.BotCommand{
			Command:     "/test",
			Description: "test",
		},
		cmd,
	)
	go tbot.RunUpdate()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		slog.Error("app - Run - signal: " + s.String())
	case err := <-tbot.Notify():
		slog.Error("tbot eror " + err.Error())
	}
	tbot.DeleteCommand()
	cancelFunction()
}

type TestCommandHandler struct {
}

type MiddlwareTest struct{}
type MiddlwareTest1 struct{}

func (m *MiddlwareTest) Check(ctx context.Context, update *update_dto.Update) bool {
	return true
}

func (m *MiddlwareTest1) Check(ctx context.Context, update *update_dto.Update) bool {
	return false
}

func (c *TestCommandHandler) Action(ctx context.Context, update *update_dto.Update) {
	if update.Message != nil {
		method.SendMessage(
			ctx,
			method_dto.SendMessage[int64]{
				ChatID: update.Message.From.Id,
				Text:   "успешно",
			},
		)
	}
}
