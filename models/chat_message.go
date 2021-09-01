package models

type ChatMessage struct {
	EmailSender  string `json:"email_sender"`
	EmailReciver string `json:"email_reciver"`
	Message      string `json:"message"`
}
