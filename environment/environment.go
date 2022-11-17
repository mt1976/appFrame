package environment

import (
	"os"
	"strings"

	fio "github.com/mt1976/AppFrame/fileio"
	logs "github.com/mt1976/AppFrame/logs"
	str "github.com/mt1976/AppFrame/strings"
	sys "github.com/mt1976/AppFrame/system"
	"github.com/spf13/viper"
)

// PATH: application\environment.go
// Language: go

// Contains Basic Application Environment Information
type Environment struct {
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

// Contains Basic Application Environment Information
var ENV Environment

// Contains Overrides for Application Environment Information
var OVR map[string]string

// Contains Extra Information to be added to Application Environment Information
var EXT map[string]string

// Contains Nothing
var DUMMY map[string]string

func init() {

	Refresh()
}

func Refresh() {
	logs.Info("(Re)load environment")
	ENV, _ = getEnvironment()
	OVR, _ = fio.GetProperties("overrides.cfg")
	EXT, _ = fio.GetProperties("extra.cfg")
}

//TODO - Add Reload Function when the re-load button is pressed.

// GetEnvironment returns the environment variables
func getEnvironment() (env Environment, err error) {
	//err := errors.New("")
	pwd, _ := os.Getwd()
	//env := Environment{}
	FILENAME := "system"
	FILEEXTN := "env"
	FULLFILENAME := FILENAME + "." + FILEEXTN

	DUMMY, _ = fio.GetProperties(FULLFILENAME)
	//fmt.Println(pwd)
	FILEPATH := pwd + "/config/"

	if sys.IsRunningInDockerContainer() {
		FILEPATH = "/config/"
	}
	viper.AddConfigPath(FILEPATH)
	viper.SetConfigName(FILENAME)
	viper.SetConfigType(FILEEXTN)

	logs.WithFields(logs.Fields{"File": FULLFILENAME, "Path": FILEPATH}).Info("Environment File")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		logs.Fatal(err)
		return Environment{}, err
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		logs.Fatal(err)
		return Environment{}, err
	}
	if env.AdditionalServices {
		env.AdditionalServicesList = strings.Split(viper.GetString("additionalServicesList"), ",")
	}

	//spew.Dump(env)

	return env, nil
}

func Debug() {
	logs.WithFields(logs.Fields{"NAME": ENV.AppName, "VERSION": ENV.AppVersion}).Info(TextGet("Application"))
	logs.WithFields(logs.Fields{"URI": ENV.DockerURI, "PORT": ENV.DockerPORT, "PROTOCOL": ENV.DockerPROTOCOL}).Info(TextGet("Container"))
	logs.WithFields(logs.Fields{"URI": ENV.AppURI, "PORT": ENV.AppPORT, "PROTOCOL": ENV.AppPROTOCOL}).Info(TextGet("Application"))
	if sys.IsRunningInDockerContainer() {
		logs.WithFields(logs.Fields{"DOCKER": "true"}).Info(TextGet("Runtime"))
	} else {
		logs.WithFields(logs.Fields{"DOCKER": "false"}).Info(TextGet("Runtime"))

	}
}

func Get(prop map[string]string, inName string, what string) string {
	what = strings.ToLower(what)
	search := inName + what
	rVal := prop[search]
	//spew.Dump(prop)
	//log.WithFields(log.Fields{"NAME": inName, "WHAT": what, "VALUE": rVal}).Info("EnvironmentGet")
	//if rVal == "" {
	//		return inName
	//	}
	return rVal
}

func EnvironmentOverrideGet(orig string, inName string, what string) string {
	return environmentOverrideGet(orig, inName, what)
}

func EnvironmentExtraGet(orig string, inName string, what string) string {
	return environmentExtraGet(orig, inName, what)
}

func environmentOverrideGet(orig string, inName string, what string) string {
	//fmt.Printf("orig: %v\n", orig)
	//fmt.Printf("inName: %v\n", inName)
	//fmt.Printf("what: %v\n", what)
	val := Get(OVR, inName, what)

	//log.WithFields(log.Fields{"orig": orig, "in": inName, "what": what, "value": val}).Info("OVR")

	if val == "" {
		return orig
	}
	//fmt.Printf("val: %v\n", val)
	return val
}

func environmentExtraGet(orig string, inName string, what string) string {
	//retVal := orig
	out := Get(EXT, inName, what)
	if out == "" {
		out = orig
	}
	//log.Info("EXT: ", orig, ":", inName, ":", what, ":", out, ":")
	out = environmentOverrideGet(out, inName, what)
	//if out2 == "" {
	//log.Info("OVR: |", inName, "|", what, "|", out, "|", orig, "|")
	return out
}

func TextGet(in string) string {
	//log.Info("TextGet: ", in)
	//log.Info("TextGet: ", lowerFirst(in)+"TXT")
	out := environmentOverrideGet("", str.LowerFirst(in), "TXT")
	//log.Info("TextGet: In :", in)
	//log.Info("TextGet: Out :", out)
	return out
}

// Return AppName
func ApplicationName() string {
	return ENV.AppName
}

// Return HostName
func HostName() string {
	return sys.SystemInfoGet().Hostname
}

// Return Application HTTP Port
func ApplicationHTTPPort() string {
	return ENV.AppPORT
}
