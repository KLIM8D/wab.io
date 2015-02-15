package utils

import (
	"bytes"
	"github.com/KLIM8D/wab.io/logs"
	"github.com/OneOfOne/xxhash/native"
	"io"
	"strings"
	"sync/atomic"
)

const (
	ALPHABET = "abcdefghijklmnopqrstuvwxyz0123456789"
	BASE     = int64(len(ALPHABET))
)

var (
	ShortUrls chan string
	count     int64 = 0
)

func Shortener(url string) uint32 {
	h := xxhash.New32()
	r := bytes.NewReader([]byte(url))
	io.Copy(h, r)

	return h.Sum32()
}

func Shortener2() {
	ShortUrls = make(chan string, 1024)
	if workers < 1 {
		workers = 1
	}

	if logs.Mode == logs.DebugMode {
		logs.Trace.Printf("Spawning %d workers\n", workers)
	}
	for i := uint8(0); i < workers; i++ {
		go func() {
			for {
				k := atomic.AddInt64(&count, int64(1))
				if logs.Mode == logs.DebugMode {
					if k > 1000 {
						logs.Trace.Printf("Shortener2 - count(k): %d\n", k)
					}
				}
				ShortUrls <- Encode(k)
			}
		}()
	}
}

func Encode(i int64) string {
	if logs.Mode == logs.DebugMode {
		logs.Trace.Printf("Encode - arg(i): %d\n", i)
	}
	if i == 0 {
		return string(ALPHABET[0])
	}

	var buffer bytes.Buffer
	for i > 0 {
		buffer.WriteByte(ALPHABET[i%BASE])
		i = i / BASE
		if logs.Mode == logs.DebugMode {
			logs.Trace.Printf("Encode - loop(i): %d\n", i)
		}
	}

	f := reverseBytes(buffer.String())
	if logs.Mode == logs.DebugMode {
		logs.Trace.Printf("Encode - string(f): %q\n", f)
	}
	return f
}

func Decode(s string) int64 {
	i := int64(0)

	for j := 0; j < len(s); j++ {
		i = (i * BASE) + int64(strings.IndexByte(ALPHABET, s[j]))
	}

	return i
}

func reverseBytes(s string) string {
	r := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-1-i]
	}
	return string(r)
}
