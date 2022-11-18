package pushover

import (
	"os"
	"strings"
	"time"

	"github.com/gregdel/pushover"
	xenv "github.com/mt1976/appFrame/environment"
	xlogs "github.com/mt1976/appFrame/logs"
	"github.com/spf13/viper"
)

func notification_GetConfig() (config CONFIG, err error) {
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
	port := xenv.ApplicationHTTPPort()

	return &pushover.Message{
		Message:     body,
		Title:       title,
		Priority:    priority,
		URL:         "http://" + xenv.HostName() + ":" + port + "/",
		URLTitle:    xenv.HostName(),
		Timestamp:   time.Now().Unix(),
		Retry:       60 * time.Second,
		Expire:      time.Hour,
		DeviceName:  strings.ReplaceAll(xenv.HostName(), ".", "_"),
		CallbackURL: "http://" + xenv.HostName() + ":" + port + "/ACKNotification",
		Sound:       pushover.SoundCosmic,
	}
}
