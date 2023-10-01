package main

import (
	"TGoBot/bot"
	command_dto "TGoBot/dto/command"
	inline_dto "TGoBot/dto/inlineQuery"
	method_dto "TGoBot/dto/method"
	update_dto "TGoBot/dto/update"
	"TGoBot/examples/collegeBot/model"
	"TGoBot/method"
	"TGoBot/pkg/logger"
	"TGoBot/request"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/signal"
	"slices"
	"strings"
	"syscall"
)

const url = "https://info.kma29.ru/api.php/"

func main() {
	ctx := context.Background()
	ctxWithCancel, cancelFunction := context.WithCancel(ctx)
	tbot := bot.NewBot(bot.BotOptions{
		Token:       "6584315900:AAHVPF9Xx_ydCfyVPZn1h_S3OPLRtWTpZ9c",
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
	// if update.Message != nil {
	// 	msgHandler.SendMessage(
	// 		method_dto.SendMessage{
	// 			ChatID: update.Message.From.Id,
	// 			Text:   "test",
	// 			ReplyMarkup: message_dto.Keyboard{
	// 				// InlineKeyboard: dto.InlineKeyboard{
	// 				// 	InlineKeyboard: [][]dto.InlineKeyboardButton{
	// 				// 		{
	// 				// 			{
	// 				// 				Text:         "test",
	// 				// 				CallbackData: "kb1",
	// 				// 			},
	// 				// 			{
	// 				// 				Text:         "test1",
	// 				// 				CallbackData: "kb2",
	// 				// 			},
	// 				// 		},
	// 				// 		{
	// 				// 			{
	// 				// 				Text:         "test3",
	// 				// 				CallbackData: "kb3",
	// 				// 			},
	// 				// 		},
	// 				// 	}},
	// 				ForceReply: message_dto.ForceReply{
	// 					ForceReply:            true,
	// 					InputFieldPlaceholder: "test",
	// 				},
	// 			},
	// 		},
	// 	)
	// }
	if update.Message != nil {
		if update.Message.Text == "run" {
			resp, err := request.Request(request.Get, url, []byte{})
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			body, _ := io.ReadAll(resp.Body)
			resp.Body.Close()
			err = os.WriteFile("test.json", body, 0644)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("write")
		}
	}
	if update.InlineQuery != nil {
		query := update.InlineQuery
		groups := make([]string, 0)
		for _, item := range c.College.Items {
			if slices.Contains(groups, item.Group) {
				continue
			}
			groups = append(groups, item.Group)
		}
		arr := []inline_dto.InlineQueryResult{}
		for _, item := range groups {
			if strings.Contains(item, query.Query) {
				answerArticle := &inline_dto.InlineQueryResultArticle{
					Type:  "article",
					Id:    item,
					Title: item,
					InputMssageContent: &inline_dto.InputTextMessageContent{
						MessageText: item,
					},
				}
				arr = append(arr, answerArticle)
			}
		}
		answer := method_dto.AnswerInlineQuery{
			InlineQueryId: query.Id,
			Results:       arr,
		}
		fmt.Println(answer)
		method.AnswerInlineQuery(ctx, answer)
	}
	if update.ChosenInlineResult != nil {
		fmt.Println(update.ChosenInlineResult.Query)
	}
}
