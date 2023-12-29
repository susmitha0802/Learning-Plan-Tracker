package alerts

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var SendEmail = func (recieverEmail string, subject string, content string) {
	err := godotenv.Load("/home/susmitha/Desktop/pet-project/lpt/backend/.env")
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	SENDGRID_API_KEY := os.Getenv("SENDGRID_API_KEY")

	senderEmail := "susmitha.papani@beautifulcode.in"
	from := mail.NewEmail("From", senderEmail)
	to := mail.NewEmail("To", recieverEmail)

	message := mail.NewSingleEmail(from, subject, to, "", content)
	client := sendgrid.NewSendClient(SENDGRID_API_KEY)

	response, err := client.Send(message)

	if err != nil {
		log.Fatalf("Error:%v", err)
	} else {
		fmt.Println("response.StatusCode", response.StatusCode)
		fmt.Println("response.Body", response.Body)
		fmt.Println("response.Headers", response.Headers)
	}
}
