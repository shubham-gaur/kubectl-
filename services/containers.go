package services

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	help "github.com/shubham-gaur/kubectl++/helper"
	log "github.com/shubham-gaur/kubectl++/logger"
)

var containerSt struct {
	containers         []string
	numberOfContainers int
}

var ctrStr string

func fetchContainers(pdIndex int, nsIndex int) {
	cmd := exec.Command("kubectl", "get", "pods", podSt.pods[pdIndex], "-n", namespacesSt.namespaces[nsIndex], "-o", "jsonpath={.spec.containers[*].name}")
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	ctrStr = string(stdout)
	containerSt.containers = strings.Fields(ctrStr)
	containerSt.numberOfContainers = len(containerSt.containers)
}

func GetTaggedContainers() (int, int, int) {
	pdIndex, nsIndex := GetTaggedPods()
	fetchContainers(pdIndex, nsIndex)
	ctMap := make(map[int]string)
	log.Info.Println("🤔 Which container you are looking for?")
	log.Info.Println("😀 Below list might help... Available containers 👇 ")
	fmt.Printf("%5v﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎\n", "")
	for ct := 0; ct < containerSt.numberOfContainers; ct++ {
		ctMap[ct] = containerSt.containers[ct]
		fmt.Printf("%10v👉 Press [%-3v]: %v\n", "", ct, containerSt.containers[ct])
	}
	fmt.Printf("%5v﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊\n", "")
	var ctIndex int
	help.TakeIntInput(&ctIndex)
	return ctIndex, pdIndex, nsIndex
}

func MarkContainer() {
	var ctIndex, pdIndex, nsIndex int
	ctIndex, pdIndex, nsIndex = GetTaggedContainers()
	options := make(map[int]string)
	options[1] = "exec command"
	options[2] = "display logs"
	options[3] = "login"
	options[-1] = "exit"
	options[-2] = "return"
	var opt int
	for {
		switch options[opt] {
		case "exec command":
			executeCmd(ctIndex, pdIndex, nsIndex)
			opt = 0
		case "display logs":
			displayLogs(ctIndex, pdIndex, nsIndex)
			opt = 0
		case "login":
			Login(ctIndex, pdIndex, nsIndex)
			opt = 0
		case "return":
			fmt.Printf("%5v﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎\n", "")
			for ct := 0; ct < containerSt.numberOfContainers; ct++ {
				fmt.Printf("%10v👉 Press [%-3v]: %v\n", "", ct, containerSt.containers[ct])
			}
			fmt.Printf("%5v﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊\n", "")
			help.TakeIntInput(&ctIndex)
			log.Info.Println("😀 Active container being set to 👉 ", containerSt.containers[ctIndex])
			opt = 0
		case "exit":
			return
		default:
			opt = help.Default(options)
		}
	}
}

func executeCmd(ctIndex int, pdIndex int, nsIndex int) {
	log.Info.Println("🤔 What command to execute?")
	c := help.RetStrBufInput()
	args := []string{"-n", namespacesSt.namespaces[nsIndex], "exec", "-it", podSt.pods[pdIndex], "-c", containerSt.containers[ctIndex], "--"}
	args = append(args, strings.Fields(c)...)
	cmd := exec.Command("kubectl", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Printf("%5v﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎\n", "")
	err := cmd.Run()
	if err != nil {
		log.Warning.Println("😖 Could not execute properly ", err)
	}
	fmt.Printf("%5v﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊\n", "")
}

func displayLogs(ctIndex int, pdIndex int, nsIndex int) {
	log.Info.Println("🤔 Following logs are found for ", containerSt.containers[ctIndex])
	cmd := exec.Command("kubectl", "-n", namespacesSt.namespaces[nsIndex], "logs", podSt.pods[pdIndex], "-c", containerSt.containers[ctIndex])
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Warning.Println("😖 Could not execute properly ", err)
	}
}

func DisplayContainers() {
	pdIndex, nsIndex := GetTaggedPods()
	fetchContainers(pdIndex, nsIndex)
	log.Info.Printf("🤔 No namespace provided; will display containers in %v pod for %v namespace", podSt.pods[pdIndex], namespacesSt.namespaces[nsIndex])
	log.Info.Println("😓 Executing command: kubectl get pods -n " + namespacesSt.namespaces[nsIndex])
	log.PrintSpecial(log.GetCurrentFunctionName(), ctrStr)
}
