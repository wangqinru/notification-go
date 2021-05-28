package awsses

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

type EmailInfo struct {
	Sender       string   `json:"sender"`
	RecipientTo  []string `json:"to"`
	RecipientCc  []string `json:"cc"`
	RecipientBcc []string `json:"bcc"`
	Subject      string   `json:"subject"`
	HtmlBody     string   `json:"htmlbody"`
	TextBody     string   `json:"textbody"`
	CharSet      string   `json:"charset"`
}

func CreateRecipientList(addresses []string) []*string {

	recipientList := []*string{}

	if len(addresses) > 0 {
		for _, addr := range addresses {
			recipientList = append(recipientList, aws.String(addr))
		}
	}

	return recipientList
}

func CreateEmailContent(charset string, data string) *ses.Content {
	content := &ses.Content{}

	if len(data) > 0 {
		data = strings.Replace(data, "       ", "\n", -1)
		content.SetCharset(charset)
		content.SetData(data)
	}

	return content
}

func SendEmail(info EmailInfo) (*ses.SendEmailOutput, error) {
	// Create a new session in the us-west-2 region.
	// Replace us-west-2 with the AWS Region you're using for Amazon SES.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)

	if err != nil {
		log.Fatalf("create session failed : %s", err.Error())
		return nil, err
	}

	// Create an SES session.
	svc := ses.New(sess)

	// Assemble the email.
	input := &ses.SendEmailInput{}

	// Set Email Destination
	destination := &ses.Destination{}

	if len(info.RecipientTo) > 0 {
		destination.SetToAddresses(CreateRecipientList(info.RecipientTo))
	}
	if len(info.RecipientBcc) > 0 {
		destination.SetBccAddresses(CreateRecipientList(info.RecipientBcc))
	}
	if len(info.RecipientCc) > 0 {
		destination.SetCcAddresses(CreateRecipientList(info.RecipientCc))
	}

	input.Destination = destination

	// Set Email Body
	message := &ses.Message{}
	body := &ses.Body{}

	if len(info.HtmlBody) > 0 {
		body.SetHtml(CreateEmailContent(info.CharSet, info.HtmlBody))
	}

	body.SetText(CreateEmailContent(info.CharSet, info.TextBody))
	message.SetBody(body)
	message.SetSubject(CreateEmailContent(info.CharSet, info.Subject))

	input.Message = message

	// Set Email Source
	input.SetSource(info.Sender)

	// Attempt to send the email.
	result, err := svc.SendEmail(input)

	// Display error messages if they occur.
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ses.ErrCodeMessageRejected:
				log.Fatalln(ses.ErrCodeMessageRejected, aerr.Error())
			case ses.ErrCodeMailFromDomainNotVerifiedException:
				log.Fatalln(ses.ErrCodeMailFromDomainNotVerifiedException, aerr.Error())
			case ses.ErrCodeConfigurationSetDoesNotExistException:
				log.Fatalln(ses.ErrCodeConfigurationSetDoesNotExistException, aerr.Error())
			default:
				log.Fatalln(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			log.Fatalln(err.Error())
		}

		return nil, err
	}

	log.Printf("Send Email : %#v", info)
	log.Println(result)

	return result, nil
}
