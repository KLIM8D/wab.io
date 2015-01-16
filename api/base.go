package api

import (
	"encoding/json"
	"github.com/klim8d/wab.io/lib"
	"io/ioutil"
)

func (self *Base) Init() *Base {
	if b, err := self.readConfiguration(); err != nil {
		panic(err)
	} else {
		return b
	}
}

func (self *Base) readConfiguration() (*Base, error) {
	if content, err := ioutil.ReadFile(self.ConfigFileName); err != nil {
		return nil, err
	} else {
		var conf lib.Configuration
		if err = json.Unmarshal(content, &conf); err != nil {
			return nil, err
		} else {
			self.Server = &WebServer{Config: &conf}
			return self, nil
		}
	}
}
