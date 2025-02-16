package mail_handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"sync"
	"time"
	"tiny-letter/email/pkg/constants"

	"gopkg.in/gomail.v2"
)

// Email represents an email structure
type Email struct {
	To      string
	Subject string
	Body    string
}

// Simulated email queue
var emailQueue = []Email{
	{"nkskl6@gmail.com", "Welcome!", "Thanks for signing up!"},
	{"nkskl6@gmail.com", "Test subject", "Test body!!!"},
	{"nkskl6@gmail.com", "New subject", "New body!!!"},
	{"nkskl6@gmail.com", "New subject 2", "New body 2!!!"},
	{"nkskl6@gmail.com", "New subject 3", "New body 3!!!"},
}

// readTemplate reads the HTML template file from the given path
func readTemplate(templatePath string) (string, error) {
	data, err := os.ReadFile(templatePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// sendEmail simulates sending an email
func sendEmail(email Email, templatePath string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", constants.HostEmail)
	m.SetHeader("To", email.To)
	m.SetHeader("Subject", email.Subject)

	// Read the HTML template from a file
	emailTemplate, err := readTemplate(templatePath)
	if err != nil {
		return err
	}

	// Parse and execute the template with dynamic values
	tmpl, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		return err
	}

	var bodyBuffer bytes.Buffer
	if err := tmpl.Execute(&bodyBuffer, nil); err != nil {
		return err
	}

	m.SetBody("text/html", bodyBuffer.String()) // Set HTML content

	// Configure SMTP
	d := gomail.NewDialer("smtp.gmail.com", 587, constants.HostEmail, constants.EmailAppPassword)

	// Simulate sending delay
	time.Sleep(500 * time.Millisecond)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	fmt.Println("Sent email to:", email.To)
	return nil
}

// batchEmailSender processes emails in batches with rate limiting
func batchEmailSender(emails []Email) {
	var wg sync.WaitGroup
	ticker := time.NewTicker(constants.RateLimit * time.Second)
	defer ticker.Stop()

	batchSize := constants.BatchSize
	for i := 0; i < len(emails); i += batchSize {
		batch := emails[i:min(i+batchSize, len(emails))]
		for _, email := range batch {
			wg.Add(1)
			go func(email Email) {
				defer wg.Done()
				err := sendEmail(email, fmt.Sprintf("%s%s", constants.TemplatePath, "subscriber_welcome.html"))
				if err != nil {
					log.Println("Failed to send email:", err)
				}
			}(email)
		}
		wg.Wait()
		<-ticker.C // Wait for the next tick instead of blocking with Sleep
	}
}

// min function for slicing batches
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
