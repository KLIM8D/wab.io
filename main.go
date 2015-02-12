package main

import (
	"flag"
	"fmt"
	"github.com/KLIM8D/wab.io/api"
	"github.com/KLIM8D/wab.io/logs"
	"github.com/KLIM8D/wab.io/utils"
	"log"
	"os"
	"os/signal"
)

var debug = flag.Int("mode", 2, "\n Options:\n 0 = Debug \n 1 = Warning \n 2 = Info \n 3 = Quiet \n")

func main() {
	flag.Parse()
	switch *debug {
	case logs.DebugMode:
		fmt.Println("Debugging...")
		logs.Mode = logs.DebugMode
		break
	case logs.WarningMode:
		logs.Mode = logs.WarningMode
		break
	case logs.InfoMode:
		logs.Mode = logs.InfoMode
		break
	case logs.Quiet:
		logs.Mode = logs.Quiet
		break
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

	b := &api.Base{ConfigFileName: "config.json"}
	b.Init()

	//Start keygenerator
	utils.GenerateKeys()

	b.Server.StartServer()
}
