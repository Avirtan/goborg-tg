package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/Avirtan/goborg-tg/method"

	method_dto "github.com/Avirtan/goborg-tg/dto/method"
	update_dto "github.com/Avirtan/goborg-tg/dto/update"
	"github.com/Avirtan/goborg-tg/examples/collegeBot/model"
	"github.com/Avirtan/goborg-tg/request"
)

const url = "https://info.kma29.ru/api.php/"

type CollegeHandler struct {
}

func (c *CollegeHandler) Action(ctx context.Context, update *update_dto.Update) {
	resp, err := request.Request(request.Get, url, []byte{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	response := model.ApiCollege{}
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println(err.Error())
	}
	resp.Body.Close()
	group := strings.ToUpper(update.Message.Text)
	for _, value := range response.Items {
		if strings.ToUpper(value.Group) == group {
			cabinet := strings.ReplaceAll(value.Cabinet, "&nbsp;", "")
			str := fmt.Sprintf("%v - %v\n%v\n%v\n%v", value.Lesson1Start, value.Lesson2End, value.Discipline, value.Teacher, cabinet)
			method.SendMessage(
				ctx,
				method_dto.SendMessage[int64]{
					ChatID: update.Message.From.Id,
					Text:   str,
				},
			)
		}
	}
}
