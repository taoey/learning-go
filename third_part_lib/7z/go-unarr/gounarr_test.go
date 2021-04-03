package go_unarr

import (
	"github.com/gen2brain/go-unarr"
	"testing"
)

func Test01(t *testing.T) {
	a, err := unarr.NewArchive("01.7z")
	if err != nil {
		panic(err)
	}
	defer a.Close()

	err2 := a.Extract("./")
	if err2 != nil {
		panic(err2)
	}
}
