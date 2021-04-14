package main

import (
	"github.com/wangqinru/notification-go/chatwork"
	"log"
)

func main() {
	apiInfo := chatwork.ApiInfo{Token: "hogehogehogehoge", RoomId: "hogehogehogehoge"}
	res, err := chatwork.MessagesPost(apiInfo, "hogehogehogehoge")

	if err != nil {
		log.Fatalf("Call Chat Work API messages Error %s \n", err.Error())
	} else {
		log.Printf("Call Chat Work API Response : %+v\n", res)
	}
}
