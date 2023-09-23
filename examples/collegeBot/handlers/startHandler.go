package handlers

import (
	"TGoBot/dto"
	"TGoBot/examples/collegeBot/model"
	"TGoBot/method"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type StartHandler struct {
	Db *sqlx.DB
}

func (c *StartHandler) Action(update *dto.Update, msgHandler *method.MethodHandler) {
	if update.Message != nil && update.Message.Text == "/start" {
		msgHandler.SendMessage(
			dto.SendMessage{
				ChatID: update.Message.From.Id,
				Text:   "Введите группу в формате */*/*",
			},
		)
		student := model.Student{
			UserId:    update.Message.From.Id,
			FirstName: update.Message.From.FirstName,
			LastName:  update.Message.From.LastName,
			Username:  update.Message.From.Username,
			Grout:     "",
		}

		_, err := c.Db.NamedExec(`INSERT INTO student (user_id, first_name, last_name, username)
		VALUES (:user_id, :first_name, :last_name, :username)`, student)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}
