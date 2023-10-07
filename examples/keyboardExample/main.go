package main

import (
	tgobot "github.com/Avirtan/TGoBot"
	command_dto "github.com/Avirtan/TGoBot/dto/command"
	message_dto "github.com/Avirtan/TGoBot/dto/message"
	method_dto "github.com/Avirtan/TGoBot/dto/method"
	update_dto "github.com/Avirtan/TGoBot/dto/update"
	"github.com/Avirtan/TGoBot/examples/collegeBot/model"
	"github.com/Avirtan/TGoBot/method"
	"github.com/Avirtan/TGoBot/pkg/logger"

	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	ctxWithCancel, cancelFunction := context.WithCancel(ctx)
	tbot := tgobot.NewBot(tgobot.BotOptions{
		Token:       "your_token",
		Ctx:         ctxWithCancel,
		LoggerLevel: logger.LevelDebug,
	})
	bytes, err := os.ReadFile("test.json")
	if err != nil {
		fmt.Println(err)
	}
	college := model.ApiCollege{}
	_ = json.Unmarshal(bytes, &college)
	cmd := &TestCommandHandler{
		College: college,
	}
	tbot.AddHandler(cmd)
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
	College model.ApiCollege
}

func (c *TestCommandHandler) Action(ctx context.Context, update *update_dto.Update) {
	if update.Message != nil {
		method.SendMessage(
			ctx,
			method_dto.SendMessage[int64]{
				ChatID: update.Message.From.Id,
				Text:   "test",
				ReplyMarkup:
				//&message_dto.ForceReply{
				// 	ForceReply:            true,
				// 	InputFieldPlaceholder: "test1",
				// },

				&message_dto.InlineKeyboardMarkup{
					InlineKeyboard: [][]message_dto.InlineKeyboardButton{
						{
							{
								Text:         "test",
								CallbackData: "kb1",
							},
							{
								Text:         "test1",
								CallbackData: "kb2",
							},
						},
						{
							{
								Text:         "test3",
								CallbackData: "kb3",
							},
						},
					}},
			},
		)
	}
	if update.CallbackQuery != nil {
		fmt.Println(update.CallbackQuery)
	}
	// if update.Message != nil {
	// 	if update.Message.Text == "run" {
	// 		resp, err := request.Request(request.Get, url, []byte{})
	// 		if err != nil {
	// 			fmt.Println(err.Error())
	// 			return
	// 		}
	// 		body, _ := io.ReadAll(resp.Body)
	// 		resp.Body.Close()
	// 		err = os.WriteFile("test.json", body, 0644)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 		fmt.Println("write")
	// 	}
	// 	if update.Message.Text == "get" {
	// 		method.GetMe(ctx)
	// 	}

	// }
	// if update.InlineQuery != nil {
	// 	query := update.InlineQuery
	// 	groups := make([]string, 0)
	// 	for _, item := range c.College.Items {
	// 		if slices.Contains(groups, item.Group) {
	// 			continue
	// 		}
	// 		groups = append(groups, item.Group)
	// 	}
	// 	arr := []inline_dto.InlineQueryResult{}
	// 	for _, item := range groups {
	// 		if strings.Contains(item, query.Query) {
	// 			answerArticle := &inline_dto.InlineQueryResultArticle{
	// 				Type:  "article",
	// 				Id:    item,
	// 				Title: item,
	// 				InputMssageContent: &inline_dto.InputTextMessageContent{
	// 					MessageText: item,
	// 				},
	// 			}
	// 			arr = append(arr, answerArticle)
	// 		}
	// 	}
	// 	answer := method_dto.AnswerInlineQuery{
	// 		InlineQueryId: query.Id,
	// 		Results:       arr,
	// 	}
	// 	fmt.Println(answer)
	// 	method.AnswerInlineQuery(ctx, answer)
	// }
	// if update.ChosenInlineResult != nil {
	// 	fmt.Println(update.ChosenInlineResult.Query)
	// }
}
