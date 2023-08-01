package logs

import "testing"

func TestInfo(t *testing.T) {
	log1 := New()
	log1.Start()
	//log1.SetPath("/Volumes/External/matttownsend/Documents/GitHub/mtChathamRateImport")
	log1.Info("Log1 Testing1")
	log1.ToFile("test")
	log1.Info("Log1 Testing2")
	log1.ToConsole()
	log1.Info("Log1 Testing3")
	log1.ToFileAndConsole("test")
	log1.Info("Log1 Testing4")
	log1.SetPath("")
	log1.Info("Log1 Testing5")
	log1.ToConsole()
	log1.Info("Log1 Testing6")
}
