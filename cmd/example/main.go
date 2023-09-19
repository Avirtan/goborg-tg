package main

import (
	"TGoBot/cmd/example/handlers"
	"TGoBot/internal/bot"
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE IF NOT EXISTS student (
	user_id    INTEGER PRIMARY KEY,
    first_name VARCHAR(250)  DEFAULT '',
    last_name  VARCHAR(250)  DEFAULT '',
	username   VARCHAR(250) DEFAULT '',
	grp VARCHAR(250) DEFAULT '',
	isSub boolean DEFAULT false
);
`

func main() {
	db, err := sqlx.Connect("postgres", "postgres://user:pass@localhost:5432/mangust?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	s := gocron.NewScheduler(time.Local)
	db.MustExec(schema)

	ctx := context.Background()
	ctxWithCancel, cancelFunction := context.WithCancel(ctx)
	tbot := bot.NewBot(bot.BotOptions{
		Token: "6584315900:AAHVPF9Xx_ydCfyVPZn1h_S3OPLRtWTpZ9c",
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
	msg := &handlers.CollegeHandler{}
	tbot.AddHandler(msg)

	start := &handlers.StartHandler{
		Db: db,
	}
	tbot.AddHandler(start)
	// cmd := &handlers.HelpHandler{}
	// tbot.AddCommand(&dto.BotCommand{
	// 	Command:     "/help",
	// 	Description: "пример использование бота",
	// }, cmd)
	// tbot.AddCommand(&dto.BotCommand{
	// 	Command:     "/test1",
	// 	Description: "тестовая1 комманда",
	// }, cmd)

	// tbot.GetCommand()
	go tbot.RunUpdate()
	s.StartAsync()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		slog.Error("app - Run - signal: " + s.String())
	case err := <-tbot.Notify():
		slog.Error("tbot eror " + err.Error())
	}
	s.Stop()
	// tbot.DeleteCommand()
	cancelFunction()
}
