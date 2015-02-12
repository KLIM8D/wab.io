package utils

import (
	"github.com/satori/go.uuid"
	"strings"
)

var (
	NewKey chan string
)

/*
	Creates a buffered channel which holds 32 uuid's
	The goroutine is non-blocking as long as the channel is not full
*/
func GenerateKeys() {
	NewKey = make(chan string, 32)
	go func() {
		for {
			u := uuid.NewV4()
			s := strings.Replace(u.String(), "-", "", -1)
			NewKey <- s
		}
	}()
}
