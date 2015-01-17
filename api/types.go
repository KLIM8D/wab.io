package api

import (
	"github.com/KLIM8D/wab.io/utils"
	"net/http"
)

type RequestHandler func(w http.ResponseWriter, r *http.Request)

type WebServer struct {
	Config  *utils.Configuration
	Factory *utils.RedisConf
}

type Base struct {
	ConfigFileName string
	Server         *WebServer
}
