package main

import (
	"log"

	"github.com/wangqinru/notification-go/awsses"
)

func awsses_test() {

	mailInfo := awsses.EmailInfo{RecipientTo: []string{"hogehoge@gmail.com"}, Subject: "Test Subject", TextBody: "Test TextBody"}

	output, err := awsses.SendEmail(mailInfo)

	if err != nil {
		log.Fatalf("Call Chat Work API messages Error %s \n", err.Error())
	} else {
		log.Printf("Call Chat Work API Response : %+v\n", output.GoString())
	}
}
