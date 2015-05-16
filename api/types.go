package api

import (
	"github.com/KLIM8D/wab.io/utils"
	"github.com/zenazn/goji/web"
	"net/http"
)

type RequestHandler func(web.C, http.ResponseWriter, *http.Request)

type WebServer struct {
	Config  *utils.Configuration
	Factory *utils.RedisConf
}

type Base struct {
	ConfigFileName string
	Server         *WebServer
}
