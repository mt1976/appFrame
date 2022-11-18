package environment

import (
	"os"
	"strings"

	xio "github.com/mt1976/appFrame/fileio"
	xlogs "github.com/mt1976/appFrame/logs"
	xsys "github.com/mt1976/appFrame/system"
	xtl "github.com/mt1976/appFrame/translate"
	"github.com/spf13/viper"
)

// PATH: application\environment.go
// Language: go

func init() {

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

	xlogs.WithFields(xlogs.Fields{"File": FULLFILENAME, "Path": FILEPATH}).Info("Environment File")

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
	if env.AdditionalServices() {
		env.additionalServicesList = strings.Split(viper.GetString("additionalServicesList"), ",")
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

func getValue(prop map[string]string, inName string, what string) string {
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
	xlogs.WithFields(xlogs.Fields{"NAME": Application.AppName, "VERSION": Application.AppVersion}).Info(xtl.Get("Application"))
	xlogs.WithFields(xlogs.Fields{"URI": Application.dockerURI, "PORT": Application.DockerPORT, "PROTOCOL": Application.DockerPROTOCOL}).Info(xtl.Get("Container"))
	xlogs.WithFields(xlogs.Fields{"URI": Application.AppURI, "PORT": Application.AppPORT, "PROTOCOL": Application.AppPROTOCOL}).Info(xtl.Get("Application"))
	if xsys.IsRunningInDockerContainer() {
		xlogs.WithFields(xlogs.Fields{"DOCKER": "true"}).Info(xtl.Get("Runtime"))
	} else {
		xlogs.WithFields(xlogs.Fields{"DOCKER": "false"}).Info(xtl.Get("Runtime"))

	}
}
