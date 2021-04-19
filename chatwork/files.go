package chatwork

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/textproto"
	"strings"
)

func FilesPost(info ApiInfo, message string, fileName string, fileBody string) (Response, error) {
	apiURL := fmt.Sprintf(ChatWorkAPIBaseURL, Version, "rooms", info.RoomId, "files")
	// mlutipart Bodyを作る
	var buffer bytes.Buffer

	mwriter := multipart.NewWriter(&buffer)

	// ヘッダを準備する
	mheader := make(textproto.MIMEHeader)
	mheader.Set("Content-Type", "text/plain")
	mheader.Set("Content-Disposition", fmt.Sprintf("form-data; name=\"file\"; filename=\"%s\"", fileName))

	// パートを追加する
	partwriter, _ := mwriter.CreatePart(mheader)
	io.Copy(partwriter, strings.NewReader(fileBody))

	// パートを追加する
	mwriter.WriteField("message", message)

	// multipart/form-dataのContent-Typeを取得
	contentType := mwriter.FormDataContentType()
	mwriter.Close()

	var response Response
	res, status, err := requestChatWorkAPI(apiURL, "POST", &buffer, info.Token, contentType)

	if err != nil {
		// response.Error = err.Error()
		return response, err
	}

	response.BuildResponse(string(res), status)

	return response, response.Error
}

func FilesPostUseMultipartBody(info ApiInfo, contentType string, body []byte) (Response, error) {
	apiURL := fmt.Sprintf(ChatWorkAPIBaseURL, Version, "rooms", info.RoomId, "files")

	var response Response
	res, status, err := requestChatWorkAPI(apiURL, "POST", bytes.NewReader(body), info.Token, contentType)

	if err != nil {
		return response, err
	}

	response.BuildResponse(string(res), status)

	return response, response.Error
}
