package postmark

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var client = &http.Client{}

func handler(receiver, templateAlias string, model map[string]interface{}) error {
	emailBody := map[string]interface{}{
		"From":          "GreatApe <no-reply@greatape.stream>",
		"To":            receiver,
		"TemplateAlias": templateAlias,
		"TemplateModel": model,
	}

	requestBody, _ := json.Marshal(emailBody)
	buffer := bytes.NewBuffer(requestBody)

	r, err := http.NewRequest("POST", "https://api.postmarkapp.com/email/withTemplate", buffer)
	if err != nil {
		return err
	}

	token := os.Getenv("POSTMARK_TOKEN")

	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Accept", "application/json")
	r.Header.Add("X-Postmark-Server-Token", token)

	resp, err := client.Do(r)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error sending email to %s", receiver)
	}

	return nil
}
