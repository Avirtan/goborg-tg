package handlers

import (
	"context"

	method_dto "github.com/Avirtan/TGoBot/dto/method"
	update_dto "github.com/Avirtan/TGoBot/dto/update"
	"github.com/Avirtan/TGoBot/method"
)

type StartHandler struct {
}

func (c *StartHandler) Action(ctx context.Context, update *update_dto.Update) {
	if update.Message != nil && update.Message.Text == "/start" {
		method.SendMessage(
			ctx,
			method_dto.SendMessage[int64]{
				ChatID: update.Message.From.Id,
				Text:   "Введите группу в формате */*/*",
			},
		)
		// student := model.Student{
		// 	UserId:    update.Message.From.Id,
		// 	FirstName: update.Message.From.FirstName,
		// 	LastName:  update.Message.From.LastName,
		// 	Username:  update.Message.From.Username,
		// 	Grout:     "",
		// }

		// _, err := c.Db.NamedExec(`INSERT INTO student (user_id, first_name, last_name, username)
		// VALUES (:user_id, :first_name, :last_name, :username)`, student)
		// if err != nil {
		// 	fmt.Println(err.Error())
		// }
	}
}
