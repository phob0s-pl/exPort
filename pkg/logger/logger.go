package logger

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func Configure() {
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
}
