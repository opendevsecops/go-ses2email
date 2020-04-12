package main

import (
	"flag"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"log"
)

type Payload struct {
	Text string `json:"text"`
}

func main() {
	fromPtr := flag.String("from", "", "From email address")
	toPtr := flag.String("to", "", "To email address")
	subjectPtr := flag.String("subject", "", "Email subject")
	htmlPtr := flag.String("html", "", "The html message to send")
	textPtr := flag.String("text", "", "The text message to send")

	flag.Parse()

	if len(*fromPtr) > 0 && len(*toPtr) > 0 && len(*subjectPtr) > 0 && len(*htmlPtr) > 0 && len(*textPtr) > 0 {
		svc := ses.New(session.New())

		input := &ses.SendEmailInput{
			Destination: &ses.Destination{
				CcAddresses: []*string{},
				ToAddresses: []*string{
					aws.String(*toPtr),
				},
			},
			Message: &ses.Message{
				Body: &ses.Body{
					Html: &ses.Content{
						Charset: aws.String("UTF-8"),
						Data:    aws.String(*htmlPtr),
					},
					Text: &ses.Content{
						Charset: aws.String("UTF-8"),
						Data:    aws.String(*textPtr),
					},
				},
				Subject: &ses.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(*subjectPtr),
				},
			},
			Source: aws.String(*fromPtr),
		}

		_, err := svc.SendEmail(input)

		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				default:
					log.Fatal(aerr)
				}
			} else {
				log.Fatal(err)
			}

			return
		}
	} else {
		log.Fatal("missing parameters")
	}
}
