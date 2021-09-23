package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	help "github.com/shubham-gaur/kubectl++/helper"
	log "github.com/shubham-gaur/kubectl++/logger"
	"github.com/shubham-gaur/kubectl++/services"
)

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println()
		log.Info.Println("😊 Thank you! Please visit again... 😊")
		os.Exit(0)
	}()

	options := make(map[int]string)
	options[-1] = "return"
	options[0] = "exit"
	options[1] = "display all namespaces"
	options[2] = "display all pods in a namespace"
	options[3] = "login specific container"
	options[4] = "collect logs of container"
	options[5] = "perform operations on container"

	log.Info.Println("😄 Initializing kubectl++ 😄")
	var opt int
	for {
		help.Default(options)
		help.TakeIntInput(&opt)
		switch options[opt] {
		case options[-1]:
			log.Info.Println("😊 Thank you! Please visit again... 😊")
			return
		case options[0]:
			log.Info.Println("😊 Thank you! Please visit again... 😊")
			return
		case options[1]:
			services.DisplayNamespaces()
		case options[2]:
			services.DisplayPods()
		case options[3]:
			services.Login()
		case options[4]:
			services.CollectLogForContainer()
		case options[5]:
			services.MarkContainer()
		}
	}
}
