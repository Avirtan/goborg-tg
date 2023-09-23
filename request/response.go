package request

import (
	"TGoBot/dto"
	"encoding/json"
	"io"
	"net/http"
)

func ResponseHandler(response *http.Response) (*dto.Response, error) {
	var responeJson dto.Response
	body, _ := io.ReadAll(response.Body)
	err := json.Unmarshal(body, &responeJson)
	if err != nil {
		return nil, err
	}
	response.Body.Close()
	return &responeJson, nil
}
