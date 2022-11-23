package pushover

import (
	"os"
	"strings"
	"time"

	"github.com/gregdel/pushover"
	xenv "github.com/mt1976/appFrame/environment"
	xlogs "github.com/mt1976/appFrame/logs"
	xsys "github.com/mt1976/appFrame/system"
	"github.com/spf13/viper"
)

var cTokens cConfig
var cAppPort string
var cHostName string
var cAppName string

type cConfig struct {
	PushoverKey   string `mapstructure:"pushoverkey"`
	PushoverToken string `mapstructure:"pushovertoken"`
}

func init() {
	cTokens, _ = getConfig()
	cHostName = xsys.Get().Hostname
	cAppName = xenv.ApplicationName()
	cAppPort = xenv.ApplicationHTTPPort()
}

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

func build(messageBody string, messageTitle string, priority int) (*pushover.Pushover, *pushover.Recipient, *pushover.Message) {
	app := pushover.New(cTokens.PushoverKey)

	recipient := pushover.NewRecipient(cTokens.PushoverToken)

	messageTitle = messageTitle + " [" + cHostName + "]"

	message := new(messageTitle, messageBody, priority)
	return app, recipient, message
}

func send(app *pushover.Pushover, message *pushover.Message, recipient *pushover.Recipient) {
	xlogs.Info("Sending Pushover Message")
	_, err := app.SendMessage(message, recipient)
	if err != nil {
		xlogs.Panic(err)
	}
}
