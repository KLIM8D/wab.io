package api

import (
	"github.com/KLIM8D/wab.io/lib"
	"net/http"
)

type RequestHandler func(w http.ResponseWriter, r *http.Request)

type WebServer struct {
	Config  *lib.Configuration
	Factory *lib.RedisConf
}

type Base struct {
	ConfigFileName string
	Server         *WebServer
}