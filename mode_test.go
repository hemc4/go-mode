package go-mode


import (
	"log"
	"os"
	"testing"
)

func TestSetModeFromEnv(t *testing.T) {
	value := os.Getenv("ENV_MODE")
	log.Println("current mode", value)
}
