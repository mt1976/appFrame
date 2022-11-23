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

var SYS xsys.SystemInfo

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
		URL:         "http://" + SYS.Hostname + ":" + port + "/",
		URLTitle:    SYS.Hostname,
		Timestamp:   time.Now().Unix(),
		Retry:       60 * time.Second,
		Expire:      time.Hour,
		DeviceName:  strings.ReplaceAll(SYS.Hostname, ".", "_"),
		CallbackURL: "http://" + SYS.Hostname + ":" + port + "/ACKNotification",
		Sound:       pushover.SoundCosmic,
	}
}

func init() {
	SYS = xsys.Get()
}
