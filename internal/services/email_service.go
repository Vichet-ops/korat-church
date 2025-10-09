package services

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
)

type EmailService struct{}

// NewEmailService creates a new email service
func NewEmailService() *EmailService {
	return &EmailService{}
}

// ContactEmailData holds the data for contact form emails
type ContactEmailData struct {
	Name    string
	Email   string
	Subject string
	Message string
}

// SendContactNotification sends an email notification to the church admin
func (es *EmailService) SendContactNotification(data ContactEmailData) error {
	// Get email configuration from environment
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	fromEmail := os.Getenv("SMTP_FROM")
	toEmail := os.Getenv("ADMIN_EMAIL")

	// If email is not configured, just log and return success
	if smtpHost == "" || smtpPort == "" {
		fmt.Println("⚠️  Email not configured. Contact form submission logged only.")
		fmt.Printf("📧 Contact Form Submission:\n")
		fmt.Printf("   Name: %s\n", data.Name)
		fmt.Printf("   Email: %s\n", data.Email)
		fmt.Printf("   Subject: %s\n", data.Subject)
		fmt.Printf("   Message: %s\n", data.Message)
		return nil
	}

	// Create email template
	emailTemplate := `
<!DOCTYPE html>
<html>
<head>
	<style>
		body { font-family: Arial, sans-serif; line-height: 1.6; color: #333; }
		.container { max-width: 600px; margin: 0 auto; padding: 20px; }
		.header { background: #2563eb; color: white; padding: 20px; text-align: center; }
		.content { background: #f9f9f9; padding: 20px; border: 1px solid #ddd; }
		.field { margin-bottom: 15px; }
		.label { font-weight: bold; color: #555; }
		.footer { margin-top: 20px; padding: 20px; text-align: center; font-size: 12px; color: #777; }
	</style>
</head>
<body>
	<div class="container">
		<div class="header">
			<h2>New Contact Form Submission</h2>
		</div>
		<div class="content">
			<div class="field">
				<div class="label">From:</div>
				<div>{{.Name}} ({{.Email}})</div>
			</div>
			<div class="field">
				<div class="label">Subject:</div>
				<div>{{.Subject}}</div>
			</div>
			<div class="field">
				<div class="label">Message:</div>
				<div>{{.Message}}</div>
			</div>
		</div>
		<div class="footer">
			<p>This email was sent from the church website contact form.</p>
		</div>
	</div>
</body>
</html>
	`

	// Parse template
	tmpl, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse email template: %w", err)
	}

	// Execute template
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to execute email template: %w", err)
	}

	// Prepare email message
	subject := fmt.Sprintf("Contact Form: %s", data.Subject)
	headers := make(map[string]string)
	headers["From"] = fromEmail
	headers["To"] = toEmail
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"UTF-8\""

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body.String()

	// Send email
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
	addr := fmt.Sprintf("%s:%s", smtpHost, smtpPort)
	
	err = smtp.SendMail(addr, auth, fromEmail, []string{toEmail}, []byte(message))
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	fmt.Println("✅ Contact form notification email sent successfully")
	return nil
}

// SendAutoReply sends an automatic reply to the person who submitted the form
func (es *EmailService) SendAutoReply(data ContactEmailData) error {
	// Get email configuration from environment
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	fromEmail := os.Getenv("SMTP_FROM")
	churchName := os.Getenv("CHURCH_NAME")

	if churchName == "" {
		churchName = "Muang Thai Korat Church"
	}

	// If email is not configured, skip auto-reply
	if smtpHost == "" || smtpPort == "" {
		return nil
	}

	// Create auto-reply template
	autoReplyTemplate := `
<!DOCTYPE html>
<html>
<head>
	<style>
		body { font-family: Arial, sans-serif; line-height: 1.6; color: #333; }
		.container { max-width: 600px; margin: 0 auto; padding: 20px; }
		.header { background: #2563eb; color: white; padding: 20px; text-align: center; }
		.content { background: #f9f9f9; padding: 20px; border: 1px solid #ddd; }
		.footer { margin-top: 20px; padding: 20px; text-align: center; font-size: 12px; color: #777; }
	</style>
</head>
<body>
	<div class="container">
		<div class="header">
			<h2>Thank You for Contacting Us</h2>
		</div>
		<div class="content">
			<p>Dear {{.Name}},</p>
			<p>Thank you for reaching out to us. We have received your message and will get back to you as soon as possible.</p>
			<p><strong>Your message:</strong></p>
			<p style="background: white; padding: 15px; border-left: 3px solid #2563eb;">
				{{.Message}}
			</p>
			<p>If you have any urgent matters, please feel free to call us directly.</p>
			<p>Blessings,<br>` + churchName + ` Team</p>
		</div>
		<div class="footer">
			<p>This is an automated message. Please do not reply to this email.</p>
		</div>
	</div>
</body>
</html>
	`

	// Parse template
	tmpl, err := template.New("autoreply").Parse(autoReplyTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse auto-reply template: %w", err)
	}

	// Execute template
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("failed to execute auto-reply template: %w", err)
	}

	// Prepare email message
	subject := "Thank you for contacting us"
	headers := make(map[string]string)
	headers["From"] = fromEmail
	headers["To"] = data.Email
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=\"UTF-8\""

	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body.String()

	// Send email
	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
	addr := fmt.Sprintf("%s:%s", smtpHost, smtpPort)
	
	err = smtp.SendMail(addr, auth, fromEmail, []string{data.Email}, []byte(message))
	if err != nil {
		// Don't fail the whole operation if auto-reply fails
		fmt.Printf("⚠️  Failed to send auto-reply: %v\n", err)
		return nil
	}

	fmt.Println("✅ Auto-reply email sent successfully")
	return nil
}

