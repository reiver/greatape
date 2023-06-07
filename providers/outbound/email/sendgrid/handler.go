package sendgrid

import (
	"fmt"
	"net/http"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func handler(receiver, content string, _ map[string]interface{}) error {
	from := mail.NewEmail("Admin", "admin@site")
	subject := "Subject"
	to := mail.NewEmail(receiver, receiver)

	plainTextContent := content
	htmlContent := fmt.Sprintf("<p>%s</p>", content)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	if response, err := client.Send(message); err != nil {
		return err
	} else if response.StatusCode != http.StatusCreated {
		return fmt.Errorf("%d", response.StatusCode)
	}

	return nil
}
