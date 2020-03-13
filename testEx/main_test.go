package testex

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("before test")
	retCode := m.Run()
	log.Println("after test")
	os.Exit(retCode)
}
