package log4go

import (
	log "github.com/jeanphorn/log4go"
	"testing"
)

func Test00(t *testing.T) {
	log.LoadConfiguration("./example.json")
	log.LOGGER("324").Info("category Test info test ...")
	log.Close()
}
