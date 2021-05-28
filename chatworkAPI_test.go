package main

import (
	"log"

	"github.com/wangqinru/notification-go/chatwork"
)

func chatwork_test() {
	apiInfo := chatwork.ApiInfo{Token: "hogehogehogehoge", RoomId: "hogehogehogehoge"}
	res, err := chatwork.MessagesPost(apiInfo, "hogehogehogehoge")

	if err != nil {
		log.Fatalf("Call Chat Work API messages Error %s \n", err.Error())
	} else {
		log.Printf("Call Chat Work API Response : %+v\n", res)
	}
}
