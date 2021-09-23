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
		log.Info.Println("😊 Thank you! Please visit again...😊")
		os.Exit(0)
	}()

	options := make(map[int]string)
	options[1] = "display namespaces"
	options[2] = "display pods"
	options[3] = "login"
	options[4] = "collect logs"
	options[5] = "select container"
	options[-1] = "exit"

	log.Info.Println("😄 Initializing kubectl++ 😄")
	var opt int
	for {
		switch options[opt] {
		case "display pods":
			services.DisplayPods()
			opt = 0
		case "display namespaces":
			services.DisplayNamespaces()
			opt = 0
		case "login":
			services.Login()
			opt = 0
		case "collect logs":
			services.CollectLogForContainer()
			opt = 0
		case "select container":
			services.MarkContainer()
			opt = 0
		case "exit":
			return
		default:
			opt = help.Default(options)
		}
	}
}
