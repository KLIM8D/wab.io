package utils

import (
	"bytes"
	"github.com/OneOfOne/xxhash/native"
	"io"
)

func Shortener(url string) uint32 {
	h := xxhash.New32()
	r := bytes.NewReader([]byte(url))
	io.Copy(h, r)

	return h.Sum32()
}
