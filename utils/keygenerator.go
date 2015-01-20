package utils

import (
	"github.com/satori/go.uuid"
	"strings"
)

var (
	NewKey chan string
)

func GenerateKeys() {
	NewKey = make(chan string)
	go func() {
		for {
			u := uuid.NewV4()
			s := strings.Replace(u.String(), "-", "", -1)
			NewKey <- s
		}
	}()
}
