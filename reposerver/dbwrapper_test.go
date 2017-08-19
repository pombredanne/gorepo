package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	//setup
	os.Setenv("DATASTORE_EMULATOR_HOST", "localhost:10101")
	os.Setenv("DATASTORE_PROJECT_ID", "gorepo")
	retCode := m.Run()
	//tear down
	os.Exit(retCode)
}

func TestTmp(t *testing.T) {
	Tmp()
}
