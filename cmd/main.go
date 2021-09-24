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
		log.Info.Printf("😊 Thank you! Please visit again... 😊")
		os.Exit(0)
	}()

	options := make(map[int]string)
	options[0] = "return"
	options[1] = "perform operations on namespace"
	options[2] = "perform operations on pod"
	options[3] = "perform operations on container"
	options[4] = "login specific container"
	options[5] = "collect logs of container"
	options[6] = "exit"

	log.Info.Printf("😄 Initializing kubectl++ 😄")
	var opt int
	for {
		help.Default(options)
		help.TakeIntInput(&opt, len(options))
		switch options[opt] {
		case options[0]:
			log.Info.Printf("😊 Thank you! Please visit again... 😊")
			return
		case options[1]:
			services.MarkNamespace()
		case options[2]:
			services.MarkPod()
		case options[3]:
			services.MarkContainer()
		case options[4]:
			services.Login()
		case options[5]:
			services.CollectLogForContainer()
		case options[6]:
			log.Info.Printf("😊 Thank you! Please visit again... 😊")
			return
		}
	}
}
