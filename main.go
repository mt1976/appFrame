package main

import (
	"fmt"
	"strings"

	xdl "github.com/mt1976/appFrame/dataloader"
	xlogs "github.com/mt1976/appFrame/logs"
	xsys "github.com/mt1976/appFrame/system"
	xtl "github.com/mt1976/appFrame/translate"
)

func main() {

	S := xsys.SYSINFO
	fmt.Printf("si: %v\n", S.String())

	L := xlogs.New()
	L.ToFileAndConsole("testingconfig")

	T := xtl.New()
	T.Verbose()
	// Test Config
	Config := *xdl.New("test", "", "")
	Config.Debug()

	tval, err := Config.GetString("test")
	if err != nil {
		panic(err)
	}

	L.WithField("test", tval).Info("Test Value")

	t0 := T.Get(tval)
	L.WithField("t0", t0).Info("Test Value")
	T.AddLocale("us")
	t1 := T.GetLocalised(strings.ToUpper(tval), "en")
	t2 := T.GetLocalised(strings.ToLower(tval), "jp")

	L.WithField("t1", t1).Info("Test Value")
	L.WithField("t2", t2).Info("Test Value")

	L.WithField("T.GetLocales()", T.GetLocales()).Info("Locales")

	Test := *xdl.New("test", "file", "")
	Test.Debug()
}
