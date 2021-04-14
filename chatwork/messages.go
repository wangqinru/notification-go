package chatwork

import (
	"fmt"
	"strings"
)

func MessagesPost(info ApiInfo, message string) (Response, error) {
	apiURL := fmt.Sprintf(ChatWorkAPIBaseURL, Version, "rooms", info.RoomId, "messages")

	body := strings.NewReader(fmt.Sprintf("body=%s", message))

	var response Response
	res, status, err := requestChatWorkAPI(apiURL, "POST", body, info.Token, FormType)

	if err != nil {
		return response, err
	}

	response.BuildResponse(string(res), status)

	return response, response.Error
}
