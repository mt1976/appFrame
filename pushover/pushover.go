package pushover

import (
	"os"
	"strings"
	"time"

	"github.com/gregdel/pushover"
	xenv "github.com/mt1976/appFrame/environment"
	xlogger "github.com/mt1976/appFrame/logs"
	xsys "github.com/mt1976/appFrame/system"
	"github.com/spf13/viper"
)

var cTokens cConfig
var cAppPort string
var cHostName string
var cAppName string
var xlogs xlogger.XLogger

type cConfig struct {
	PushoverKey   string `mapstructure:"pushoverkey"`
	PushoverToken string `mapstructure:"pushovertoken"`
}

// The init function initializes global variables for the application.
func init() {
	cTokens, _ = getConfig()
	cHostName = xsys.Get().Hostname
	cAppName = xenv.ApplicationName()
	cAppPort = xenv.ApplicationHTTPPort()
	xlogs = xlogger.New()
}

// The getConfig function reads a configuration file named "notifications.env" from the "config"
// directory in the current working directory and unmarshals it into a struct named cConfig.
func getConfig() (config cConfig, err error) {
	// get current os directory path
	pwd, _ := os.Getwd()

	//fmt.Println(pwd)
	viper.AddConfigPath(pwd + "/config/")
	viper.SetConfigName("notifications")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		xlogs.Fatal(err)
		return
	}

	err = viper.Unmarshal(&config)
	//spew.Dump(config)

	return config, err
}

// The function "new" creates a new pushover.Message object with the given title, body, and priority,
// along with other optional parameters.
func new(title string, body string, priority int) *pushover.Message {

	return &pushover.Message{
		Message:     body,
		Title:       title,
		Priority:    priority,
		URL:         "http://" + cHostName + ":" + cAppPort + "/",
		URLTitle:    cHostName,
		Timestamp:   time.Now().Unix(),
		Retry:       60 * time.Second,
		Expire:      time.Hour,
		DeviceName:  strings.ReplaceAll(cHostName, ".", "_"),
		CallbackURL: "http://" + cHostName + ":" + cAppPort + "/ACK",
		Sound:       pushover.SoundCosmic,
	}
}

// The function builds a push notification message with a title, body, and priority using the Pushover
// library in Go.
func build(messageBody string, messageTitle string, priority int) (*pushover.Pushover, *pushover.Recipient, *pushover.Message) {
	app := pushover.New(cTokens.PushoverKey)

	recipient := pushover.NewRecipient(cTokens.PushoverToken)

	messageTitle = messageTitle + " [" + cHostName + "]"

	message := new(messageTitle, messageBody, priority)
	return app, recipient, message
}

// The function sends a Pushover message using the provided app, message, and recipient.
func send(app *pushover.Pushover, message *pushover.Message, recipient *pushover.Recipient) {
	xlogs.Info("Sending Pushover Message")
	_, err := app.SendMessage(message, recipient)
	if err != nil {
		xlogs.Panic(err)
	}
}
