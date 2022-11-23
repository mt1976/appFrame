package pushover

import (
	"github.com/gregdel/pushover"
)

func Emergency(messageTitle string, messageBody string) {
	// Build Message
	app, recipient, message := build(messageBody, messageTitle, pushover.PriorityEmergency)
	// Send the message to the recipient
	send(app, message, recipient)
}

func Normal(messageTitle string, messageBody string) {
	// Build Message
	app, recipient, message := build(messageBody, messageTitle, pushover.PriorityNormal)
	// Send the message to the recipient
	send(app, message, recipient)
}

func WithURL(messageTitle string, messageBody string, url string) {
	// Build Message
	app, recipient, message := build(messageBody, messageTitle, pushover.PriorityNormal)
	// Add URI/URL
	message.URL = message.URL + url
	// Send the message to the recipient
	send(app, message, recipient)
}

func High(messageTitle string, messageBody string) {
	// Build Message
	app, recipient, message := build(messageBody, messageTitle, pushover.PriorityHigh)
	// Send the message to the recipient
	send(app, message, recipient)
}

func Low(messageTitle string, messageBody string) {
	//cfg, _ := notification_GetConfig()
	app, recipient, message := build(messageBody, messageTitle, pushover.PriorityLow)
	// Send the message to the recipient
	send(app, message, recipient)
}
