package request

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/Avirtan/goborg-tg/dto"
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

func ResponseHandlerToType[TypeData any](response *http.Response) (*TypeData, error) {
	var responeJson dto.Response
	body, _ := io.ReadAll(response.Body)
	err := json.Unmarshal(body, &responeJson)
	if err != nil {
		return nil, err
	}
	response.Body.Close()
	jsonData, _ := json.Marshal(responeJson.Result)
	data := new(TypeData)
	json.Unmarshal(jsonData, &data)
	return data, nil
}
