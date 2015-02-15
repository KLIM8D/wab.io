package utils

import (
	"github.com/KLIM8D/wab.io/logs"
	"runtime"
	"testing"
	"time"
)

var (
	urls = []string{"http://www.google.com/", "https://github.com", "https://github.com/KLIM8D/wab.io",
		"https://github.com/KLIM8D/wab.io/commit/3abca64ac19a2076be3717ddee350cd584bfdbe0#diff-a084b794bc0759e7a6b77810e01874f2", "https://github.com/KLIM8D/wab.io/blob/0012fdf2c97a6b561450f437e163566f3abea8c3/.gitignore",
		"https://github.com/KLIM8D/wab.io/pulls",
		"https://github.com/KLIM8D/wab.io/graphs/contributors?from=2015-01-12&to=2015-01-15&type=c",
		"https://github.com/KLIM8D/wab.io/blob/07dacf62edf9b9584e06b0d2200daa7555788daa/.travis.yml"}
)

const (
	BENCHDEBUG = false
)

func init() {
	if BENCHDEBUG {
		logs.Mode = logs.DebugMode
	} else {
		logs.Mode = logs.Quiet
	}

	logs.Initialize()
	Shortener2()
}

// Test the format of the shorten URLs
// Pre-condition: none
// Post-condition: n unique shortened urls are generated
func TestUrlShortener(t *testing.T) {
	t.Log("### TestUrlShortener ###")

	for i := 0; i < 10; i++ {
		for range urls {
			u := <-ShortUrls
			t.Logf("ShortenedURL: %q\n", u)
		}
	}
}

func BenchmarkShortener(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	for i := 0; i < b.N; i++ {
		for _, v := range urls {
			Shortener(v)
		}
	}
}

func BenchmarkShortener2(b *testing.B) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	b.StopTimer()
	time.Sleep(5 * time.Second)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		for range urls {
			<-ShortUrls
		}
	}
}
