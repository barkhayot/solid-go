package main

import "fmt"

/*
Dependency Inversion Principle (DIP):
- High-level modules should not depend on low-level modules. Both should depend on abstractions.
- Abstractions should not depend on details. Details should depend on abstractions.
*/

// Define an abstraction (interface) for sending messages.
type MessageSender interface {
	Send(to string, message string)
}

// Concrete implementation of EmailService.
type EmailService struct{}

func (e *EmailService) Send(to string, message string) {
	fmt.Printf("Sending Email to %s: %s\n", to, message)
}

// Concrete implementation of SMSService.
type SMSService struct{}

func (s *SMSService) Send(to string, message string) {
	fmt.Printf("Sending SMS to %s: %s\n", to, message)
}

// NotificationService depends on the interface, not a specific implementation.
type NotificationService struct {
	messageSender MessageSender
}

// Constructor function for NotificationService.
func NewNotificationService(sender MessageSender) *NotificationService {
	return &NotificationService{messageSender: sender}
}

// Uses the MessageSender interface, making it flexible.
func (n *NotificationService) Notify(to string, message string) {
	n.messageSender.Send(to, message)
}

// func main() {
// 	emailService := &EmailService{}
// 	smsService := &SMSService{}

// 	// We can inject different implementations without modifying NotificationService.
// 	notificationByEmail := NewNotificationService(emailService)
// 	notificationBySMS := NewNotificationService(smsService)

// 	notificationByEmail.Notify("user@example.com", "Hello via Email!")
// 	notificationBySMS.Notify("1234567890", "Hello via SMS!")
// }
