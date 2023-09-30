package main

import (
	"TGoBot/bot"
	command_dto "TGoBot/dto/command"
	"TGoBot/examples/collegeBot/handlers"
	"TGoBot/pkg/logger"
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	ctxWithCancel, cancelFunction := context.WithCancel(ctx)
	tbot := bot.NewBot(bot.BotOptions{
		Token:       "6584315900:AAHVPF9Xx_ydCfyVPZn1h_S3OPLRtWTpZ9c",
		Ctx:         ctxWithCancel,
		LoggerLevel: logger.LevelDebug,
	})

	cmd := &handlers.TestCommandHandler{}
	tbot.AddCommand(&command_dto.BotCommand{
		Command:     "/test",
		Description: "пример использование бота",
	}, cmd)
	// tbot.AddHandler(cmd)

	go tbot.RunUpdate()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		slog.Error("app - Run - signal: " + s.String())
	case err := <-tbot.Notify():
		slog.Error("tbot eror " + err.Error())
	}
	// s.Stop()
	tbot.DeleteCommand()
	cancelFunction()
}
