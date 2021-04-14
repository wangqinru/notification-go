package chatwork

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	ChatWorkAPIBaseURL = "https://api.chatwork.com/%s/%s/%s/%s"
	Version            = "v2"
	FormType           = "application/x-www-form-urlencoded"
	MultiFormType      = "multipart/form-data"
)

type ApiInfo struct {
	Token  string
	RoomId string
}

type Response struct {
	Result string
	Error  error
}

func (res *Response) BuildResponse(result string, status int) *Response {
	if status != http.StatusOK {
		res.Error = errors.New(result)
		res.Result = ""
	} else {
		res.Result = result
		res.Error = nil
	}

	return res
}

func requestChatWorkAPI(apiURL string, method string, body io.Reader, token string, contentType string) ([]byte, int, error) {
	var request *http.Request
	var err error

	request, err = http.NewRequest(method, apiURL, body)

	if err != nil {
		log.Fatalf("Error at ChatWork API request:%#v", err)
		return nil, 0, err
	}

	request.Header.Set("Content-Type", contentType)
	request.Header.Set("X-ChatWorkToken", token)

	// timeoutを暫定60秒に設定する
	client := &http.Client{Timeout: time.Duration(60) * time.Second}

	response, err := client.Do(request)
	if err != nil {
		log.Fatalf("Error at ChatWork API response:%#v", err)
		return nil, 0, err
	}

	defer response.Body.Close()

	resbody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error at read ChatWork API response:%#v", err)
		return nil, 0, err
	}

	return resbody, response.StatusCode, nil
}
