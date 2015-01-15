package api

import (
	lib "../lib"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (self *WebServer) StartServer() error {
	var err error

	factory = lib.NewFactory(self.Config.Redis.Host)
	base = self.Config.Web.Base

	//Routes
	http.HandleFunc("/", self.handleIndex)
	p := fmt.Sprintf(":%d", self.Config.Web.Port)

	log.Printf("Server started. Listening on port %d\n", self.Config.Web.Port)
	if err = http.ListenAndServe(p, nil); err != nil {
		return err
	}

	return nil
}

func (self *WebServer) handleIndex(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		shortenUrl(res, req)
	} else if req.Method == "GET" && len(req.RequestURI) > 1 {
		redirectUrl(res, req)
	} else {
		if content, err := ioutil.ReadFile("web/index.html"); err != nil {
			fmt.Fprintf(res, "<html><head></head><body><span>An error occurred</span></body></html>")
			log.Println("Error: ", err)
		} else {
			fmt.Fprintf(res, string(content))
		}
	}
}
