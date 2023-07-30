package logs

import "testing"

func TestInfo(t *testing.T) {
	Start()
	SetPath("/Volumes/External/matttownsend/Documents/GitHub/mtChathamRateImport")
	Info("Testing1")
	ToFile("test")
	Info("Testing2")
	ToConsole()
	Info("Testing3")
	ToFileAndConsole("test")
	Info("Testing4")
	SetPath("")
	Info("Testing5")
	ToConsole()
	Info("Testing6")
}
