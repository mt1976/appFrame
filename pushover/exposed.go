package pushover

import (
	"github.com/gregdel/pushover"
	xenv "github.com/mt1976/appFrame/environment"
	xlogs "github.com/mt1976/appFrame/logs"
)

type CONFIG struct {
	PushoverKey   string `mapstructure:"pushoverkey"`
	PushoverToken string `mapstructure:"pushovertoken"`
}

func Emergency(messageTitle string, messageBody string) {
	cfg, _ := notification_GetConfig()
	//fmt.Printf("cfg: %v\n", cfg)
	app := pushover.New(cfg.PushoverKey)

	// Create a new recipient
	recipient := pushover.NewRecipient(cfg.PushoverToken)

	// Create the message to send
	messageBody = messageBody + " - " + xenv.HostName()
	messageTitle = "[" + xenv.ApplicationName() + "] Notification - " + messageTitle + " - " + xenv.HostName()

	// NOTE Notification Message & Title
	message := new(messageTitle, messageBody, pushover.PriorityEmergency)

	// Send the message to the recipient
	_, err := app.SendMessage(message, recipient)
	if err != nil {
		xlogs.Panic(err)
	}
}

func Normal(messageTitle string, messageBody string) {
	cfg, _ := notification_GetConfig()
	app := pushover.New(cfg.PushoverKey)

	// Create a new recipient
	recipient := pushover.NewRecipient(cfg.PushoverToken)

	// Create the message to send
	messageBody = messageBody + " - " + xenv.HostName()
	messageTitle = "[" + xenv.ApplicationName() + "] Notification - " + messageTitle + " - " + xenv.HostName()

	// NOTE Notification Message & Title
	message := new(messageTitle, messageBody, pushover.PriorityNormal)

	// Send the message to the recipient
	_, err := app.SendMessage(message, recipient)
	if err != nil {
		xlogs.Panic(err)
	}
}

func WithURL(messageTitle string, messageBody string, url string) {
	cfg, _ := notification_GetConfig()
	app := pushover.New(cfg.PushoverKey)

	// Create a new recipient
	recipient := pushover.NewRecipient(cfg.PushoverToken)

	// Create the message to send
	messageBody = messageBody + " - " + xenv.HostName()
	messageTitle = "[" + xenv.ApplicationName() + "] Notification - " + messageTitle + " - " + xenv.HostName()

	// NOTE Notification Message & Title
	message := new(messageTitle, messageBody, pushover.PriorityNormal)
	message.URL = message.URL + url

	// Send the message to the recipient
	_, err := app.SendMessage(message, recipient)
	if err != nil {
		xlogs.Panic(err)
	}
}

func High(messageTitle string, messageBody string) {
	cfg, _ := notification_GetConfig()
	app := pushover.New(cfg.PushoverKey)

	// Create a new recipient
	recipient := pushover.NewRecipient(cfg.PushoverToken)

	// Create the message to send
	messageBody = messageBody + " - " + xenv.HostName()
	messageTitle = "[" + xenv.ApplicationName() + "] Notification - " + messageTitle + " - " + xenv.HostName()

	// NOTE Notification Message
	message := new(messageTitle, messageBody, pushover.PriorityHigh)

	// Send the message to the recipient
	_, err := app.SendMessage(message, recipient)
	if err != nil {
		xlogs.Panic(err)
	}
}

func Low(messageTitle string, messageBody string) {
	cfg, _ := notification_GetConfig()
	app := pushover.New(cfg.PushoverKey)

	// Create a new recipient
	recipient := pushover.NewRecipient(cfg.PushoverToken)

	// Create the message to send
	messageBody = messageBody + " - " + xenv.HostName()
	messageTitle = "[" + xenv.ApplicationName() + "] Notification - " + messageTitle + " - " + xenv.HostName()

	// NOTE Notification Message
	message := new(messageTitle, messageBody, pushover.PriorityLow)

	// Send the message to the recipient
	_, err := app.SendMessage(message, recipient)
	if err != nil {
		xlogs.Panic(err)
	}
}
