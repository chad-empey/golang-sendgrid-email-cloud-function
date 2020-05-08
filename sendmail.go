package p

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type requestBody struct {
	From             string `json:"from"`
	FromName         string `json:"fromName"`
	To               string `json:"to"`
	ToName           string `json:"toName"`
	Subject          string `json:"subject"`
	HTML             string `json:"html"`
	PlainTextContent string `json:"plainText"`
}

// SendEmail is an HTTP Cloud Function that sends an email powered by SendGrid's API.
func SendEmail(w http.ResponseWriter, r *http.Request) {
	body := new(requestBody)

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	from := mail.NewEmail(body.FromName, body.From)
	subject := body.Subject
	to := mail.NewEmail(body.ToName, body.To)
	plainTextContent := body.PlainTextContent
	htmlContent := body.HTML
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		w.Write([]byte(response.Body))
		return
	}
}
