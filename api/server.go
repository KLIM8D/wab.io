package api

import (
	"fmt"
	"github.com/KLIM8D/wab.io/logs"
	"github.com/KLIM8D/wab.io/utils"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"
)

const (
	DATA_PATH_PREFIX = "web/assets/"
)

func (self *WebServer) StartServer() error {
	factory = utils.NewFactory(self.Config.Redis.Host)
	base = self.Config.Web.Base

	//Routes
	//goji.Use(self.handleRoute)
	goji.Get("/img/*", handleImage)
	goji.Get("/css/*", handleCss)
	goji.Get("/*", self.handleIndex)
	goji.Post("/", shortenUrl)

	goji.Serve()

	return nil
}

func handleCss(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, filepath.Join(DATA_PATH_PREFIX, "css", r.URL.Path[5:]))
}

func handleImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "image/jpeg")
	w.Header().Set("cache-control", "public, max-age=259200")
	http.ServeFile(w, r, filepath.Join(DATA_PATH_PREFIX, "img", r.URL.Path[5:]))
}

func (self *WebServer) handleRoute(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		h.ServeHTTP(w, r)

		elapsed := time.Since(start)
		ru, ra, rf := r.RequestURI, r.RemoteAddr, r.Form
		reqInfo := fmt.Sprintf("URI: %q USER:%q FORM:%v", ru, ra, rf)

		logs.Trace.Printf("Request handled: %q, elapsed %d ns\n",
			reqInfo, elapsed.Nanoseconds())
	}
	return http.HandlerFunc(fn)
}

func (self *WebServer) handleIndex(c web.C, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && len(r.RequestURI) > 1 {
		redirectUrl(c, w, r)
	} else {
		if content, err := ioutil.ReadFile("web/index.html"); err != nil {
			fmt.Fprintf(w, `An error occurred`)
			logs.Error.Println("Error: ", err.Error())
		} else {
			fmt.Fprintf(w, string(content))
		}
	}
}
