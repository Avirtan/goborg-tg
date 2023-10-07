package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	tgobot "github.com/Avirtan/TGoBot"
	command_dto "github.com/Avirtan/TGoBot/dto/command"
	"github.com/Avirtan/TGoBot/examples/collegeBot/handlers"
)

func main() {
	// db, err := sqlx.Connect("postgres", "postgres://user:pass@localhost:5432/mangust?sslmode=disable")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// s := gocron.NewScheduler(time.Local)
	// db.MustExec(schema)

	ctx := context.Background()
	ctxWithCancel, cancelFunction := context.WithCancel(ctx)
	tbot := tgobot.NewBot(tgobot.BotOptions{
		Token: "your_token",
		Ctx:   ctxWithCancel,
	})

	// s.Every(1).Day().At("23:48").Do(func() {
	// 	students := []model.Student{}
	// 	err = db.Select(&students, "SELECT * FROM student")
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// 	for _, value := range students {
	// 		tbot.GetMethodHandler().SendMessage(
	// 			dto.SendMessage{
	// 				ChatID: value.UserId,
	// 				Text:   "привет",
	// 			},
	// 		)
	// 	}
	// })
	// msg := &handlers.CollegeHandler{}
	// tbot.AddHandler(msg)

	// start := &handlers.StartHandler{
	// 	Db: db,
	// }
	cmd := &handlers.TestCommandHandler{}
	tbot.AddCommand(&command_dto.BotCommand{
		Command:     "/test",
		Description: "пример использование бота",
	}, cmd)
	// tbot.AddHandler(cmd)
	// tbot.AddCommand(&dto.BotCommand{
	// 	Command:     "/test1",
	// 	Description: "тестовая1 комманда",
	// }, cmd)

	go tbot.RunUpdate()
	// s.StartAsync()
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
