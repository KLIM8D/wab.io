package utils

import (
	"github.com/satori/go.uuid"
	"runtime"
	"strings"
)

var (
	NewKey  chan string
	workers = uint8(runtime.NumCPU() / 2)
)

/*
Creates a buffered channel which holds 32 uuid's
The goroutine is non-blocking as long as the channel is not full
*/
func GenerateKeys() {
	NewKey = make(chan string, 32)
	if workers < 1 {
		workers = 1
	}

	for i := uint8(0); i < workers; i++ {
		go func() {
			for {
				u := uuid.NewV4()
				s := strings.Replace(u.String(), "-", "", -1)
				NewKey <- s
			}
		}()
	}
}
