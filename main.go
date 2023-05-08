package main

//go:generate sh injectGitVars.sh

import (
	"fmt"

	osc "github.com/hypebeast/go-osc/osc"
	log "github.com/s00500/env_logger"
)

func main() {
	log.Info("Hello, world!")

	OSC_Server("127.0.0.1", 8000)


	for {









		// err := "" // replace later
		// if err != nil {
		// 	log.Error("Error: %s", err)
		// 	break
		// }
	}

	log.Error("Program terminated")
}

func OSC_Server(ip string, port int) {
	addr := fmt.Sprintf("%s:%d", ip, port)
	d := osc.NewStandardDispatcher()
	d.AddMsgHandler("/message/address", func(msg *osc.Message) {
		osc.PrintMessage(msg)
	})

	server := &osc.Server{
		Addr: addr,
		Dispatcher:d,
	}
	server.ListenAndServe()
}		


