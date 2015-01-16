package main

import (
	"flag"
	"fmt"
	"github.com/KLIM8D/wab.io/api"
	"github.com/KLIM8D/wab.io/logs"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
)

var debug = flag.Bool("debug", false, "Turn on debug info")

func main() {
	flag.Parse()
	if *debug {
		fmt.Println("Debugging...")
		logs.Output = os.Stderr
	} else {
		logs.Output = ioutil.Discard
	}

	logs.Initialize()

	// capture ctrl+c and perform clean-up
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for sig := range c {
			fmt.Println()
			log.Printf("captured %v, exiting...", sig)
			os.Exit(0)
		}
	}()

	//api.StartServer()
	b := &api.Base{ConfigFileName: "config.json"}
	b.Init()
	b.Server.StartServer()
}
