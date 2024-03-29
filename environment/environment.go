package environment

import (
	"os"
	"strings"

	xio "github.com/mt1976/appFrame/fileio"
	xlogger "github.com/mt1976/appFrame/logs"
	xsys "github.com/mt1976/appFrame/system"
	xtl "github.com/mt1976/appFrame/translate"
	"github.com/spf13/viper"
)

// Contains Basic Application environment Information
type environment struct {
	DockerURI              string `mapstructure:"dockerURI"`
	DockerPORT             string `mapstructure:"dockerPORT"`
	DockerPROTOCOL         string `mapstructure:"dockerPROTOCOL"`
	AppName                string `mapstructure:"appName"`
	AppVersion             string `mapstructure:"appVersion"`
	AppURI                 string `mapstructure:"appURI"`
	AppPORT                string `mapstructure:"appPORT"`
	AppPROTOCOL            string `mapstructure:"appPROTOCOL"`
	AppGlyph               string `mapstructure:"appGLYPH"`
	AppGlyphColor          string `mapstructure:"appGLYPHCOLOR"`
	UserName               string `mapstructure:"userName"`
	UserPassword           string `mapstructure:"userPassword"`
	UserSecret             string `mapstructure:"userSecret"`
	AppTemplate            string `mapstructure:"appTemplate"`
	AdditionalServices     bool   `mapstructure:"additionalServices"`
	AdditionalServicesList []string
}

var xlogs xlogger.XLogger
var t *xtl.Translator

// PATH: application\environment.go
// Language: go

func init() {
	xlogs = xlogger.New()
	t = xtl.New()
	Refresh()
}

//TODO - Add Reload Function when the re-load button is pressed.

// GetEnvironment returns the environment variables
func getEnvironment() (env environment, err error) {
	//err := errors.New("")
	pwd, _ := os.Getwd()
	//env := Environment{}
	FILENAME := "system"
	FILEEXTN := "env"
	FULLFILENAME := FILENAME + "." + FILEEXTN

	DUMMY, _ = xio.GetPropertiesFile(FULLFILENAME)
	//fmt.Println(pwd)
	FILEPATH := pwd + "/config/"

	if xsys.IsRunningInDockerContainer() {
		FILEPATH = "/config/"
	}
	viper.AddConfigPath(FILEPATH)
	viper.SetConfigName(FILENAME)
	viper.SetConfigType(FILEEXTN)

	//xlogs.WithFields(xlogs.Fields{"File": FULLFILENAME, "Path": FILEPATH}).Info("Environment")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		xlogs.Fatal(err)
		return environment{}, err
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		xlogs.Fatal(err)
		return environment{}, err
	}
	if env.AdditionalServices {
		env.AdditionalServicesList = strings.Split(viper.GetString("additionalServicesList"), ",")
	}

	//spew.Dump(env)

	return env, nil
}

func getOverride(orig string, inName string, what string) string {
	//fmt.Printf("orig: %v\n", orig)
	//fmt.Printf("inName: %v\n", inName)
	//fmt.Printf("what: %v\n", what)
	val := getValue(Overrides, inName, what)

	//log.WithFields(log.Fields{"orig": orig, "in": inName, "what": what, "value": val}).Info("OVR")

	if val == "" {
		return orig
	}
	//fmt.Printf("val: %v\n", val)
	return val
}

func getExtra(orig string, inName string, what string) string {
	//retVal := orig
	out := getValue(Extras, inName, what)
	if out == "" {
		out = orig
	}
	//log.Info("EXT: ", orig, ":", inName, ":", what, ":", out, ":")
	out = getOverride(out, inName, what)
	//if out2 == "" {
	//log.Info("OVR: |", inName, "|", what, "|", out, "|", orig, "|")
	return out
}

func getValue(prop map[string]string, inProperty string, what string) string {

	search := inProperty + what

	low_var := strings.ToLower(search)
	rtn_var := prop[low_var]
	env_var := "APP_" + strings.ToUpper(search)

	xlogs.Info("GetApplicationProperty : " + low_var + " " + env_var)

	env_value := os.Getenv(env_var)
	if env_value != "" {
		xlogs.Info(inProperty, rtn_var, env_var, env_value)
		return env_value
	}

	xlogs.Info("GetApplicationProperty", rtn_var)

	return rtn_var
}

func getConfig(orig string, inName string, what string) string {
	//fmt.Printf("orig: %v\n", orig)
	//fmt.Printf("inName: %v\n", inName)
	//fmt.Printf("what: %v\n", what)
	val := getValue(Config, inName, what)

	//log.WithFields(log.Fields{"orig": orig, "in": inName, "what": what, "value": val}).Info("OVR")

	if val == "" {
		return orig
	}
	//fmt.Printf("val: %v\n", val)
	return val
}

func refresh() {
	xlogs.Info("(Re)load environment")
	Application, _ = getEnvironment()
	Config, _ = xio.GetPropertiesFile("application.cfg")
	Overrides, _ = xio.GetPropertiesFile("overrides.cfg")
	Extras, _ = xio.GetPropertiesFile("extra.cfg")
}

func debug() {
	xlogs.WithFields(xlogger.Fields{"NAME": Name(), "VERSION": Version()}).Info(t.Get("Application"))
	xlogs.WithFields(xlogger.Fields{"URI": DockerURI(), "PORT": DockerPort(), "PROTOCOL": DockerProtocol()}).Info(t.Get("Container"))
	xlogs.WithFields(xlogger.Fields{"URI": URI(), "PORT": Port(), "PROTOCOL": Protocol()}).Info(t.Get("Application"))
	if xsys.IsRunningInDockerContainer() {
		xlogs.WithFields(xlogger.Fields{"DOCKER": "true"}).Info(t.Get("Runtime"))
	} else {
		xlogs.WithFields(xlogger.Fields{"DOCKER": "false"}).Info(t.Get("Runtime"))

	}
}
