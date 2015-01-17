package api

import (
	"fmt"
	"github.com/KLIM8D/wab.io/logs"
	"github.com/KLIM8D/wab.io/utils"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (self *WebServer) StartServer() error {
	var err error

	factory = utils.NewFactory(self.Config.Redis.Host)
	base = self.Config.Web.Base

	//Routes
	http.HandleFunc("/", self.handleRoute())
	p := fmt.Sprintf(":%d", self.Config.Web.Port)

	log.Printf("Server started. Listening on port %d\n", self.Config.Web.Port)
	if err = http.ListenAndServe(p, nil); err != nil {
		return err
	}

	return nil
}

func (self *WebServer) handleRoute() RequestHandler {
	if logs.Mode == logs.DebugMode {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			self.handleIndex(w, r)

			elapsed := time.Since(start)
			reqInfo := fmt.Sprintf("URI: %q USER:%q FORM:%v", r.RequestURI, r.RemoteAddr, r.Form)
			logs.Trace.Printf("Request handled: %q, elapsed %d ns\n", reqInfo, elapsed.Nanoseconds())
		}
	} else {
		return self.handleIndex
	}
}

func (self *WebServer) handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		shortenUrl(w, r)
	} else if r.Method == "GET" && len(r.RequestURI) > 1 {
		redirectUrl(w, r)
	} else {
		if content, err := ioutil.ReadFile("web/index.html"); err != nil {
			fmt.Fprintf(w, "<html><head></head><body><span>An error occurred</span></body></html>")
			logs.Error.Println("Error: ", err.Error())
		} else {
			fmt.Fprintf(w, string(content))
		}
	}
}
