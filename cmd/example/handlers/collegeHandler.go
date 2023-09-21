package handlers

import (
	"TGoBot/cmd/example/model"
	"TGoBot/dto"
	"TGoBot/method"
	"TGoBot/request"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

const url = "https://info.kma29.ru/api.php/"

type CollegeHandler struct {
}

func (c *CollegeHandler) Action(update *dto.Update, msgHandler *method.MethodHandler) {
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
			msgHandler.SendMessage(
				dto.SendMessage{
					ChatID: update.Message.From.Id,
					Text:   str,
				},
			)
		}
	}
}
