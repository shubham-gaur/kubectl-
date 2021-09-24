package services

import (
	"fmt"
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
	kargs := []string{"get", "pods", podSt.pods[pdIndex], "-n", namespacesSt.namespaces[nsIndex], "-o", "jsonpath={.spec.containers[*].name}"}
	ctrStr = help.ExecKubectlCmd(kargs...)
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
		fmt.Printf("%10v👉 Press [%-2v]: %v\n", "", ct, containerSt.containers[ct])
	}
	fmt.Printf("%5v﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊\n", "")
	var ctIndex int
	help.TakeIntInput(&ctIndex, containerSt.numberOfContainers)
	return ctIndex, pdIndex, nsIndex
}

func MarkContainer() {
	var ctIndex, pdIndex, nsIndex int
	ctIndex, pdIndex, nsIndex = GetTaggedContainers()
	options := make(map[int]string)
	options[0] = "return"
	options[1] = "login container"
	options[2] = "exec command in container"
	options[3] = "display logs of container"
	options[4] = "return to main"
	var opt int
	for {
		help.Default(options)
		help.TakeIntInput(&opt, len(options))
		switch options[opt] {
		case options[0]:
			fmt.Printf("%5v﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎\n", "")
			for ct := 0; ct < containerSt.numberOfContainers; ct++ {
				fmt.Printf("%10v👉 Press [%-2v]: %v\n", "", ct, containerSt.containers[ct])
			}
			fmt.Printf("%5v﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊\n", "")
			help.TakeIntInput(&ctIndex, containerSt.numberOfContainers)
			log.Info.Println("😀 Active container being set to 👉 ", containerSt.containers[ctIndex])
		case options[1]:
			Login(ctIndex, pdIndex, nsIndex)
		case options[2]:
			executeCmd(ctIndex, pdIndex, nsIndex)
		case options[3]:
			displayLogs(ctIndex, pdIndex, nsIndex)
		case options[4]:
			return
		}
	}
}

func executeCmd(ctIndex int, pdIndex int, nsIndex int) {
	log.Info.Println("🤔 What command to execute?")
	c := help.RetStrBufInput()
	kargs := []string{"-n", namespacesSt.namespaces[nsIndex], "exec", "-it", podSt.pods[pdIndex], "-c", containerSt.containers[ctIndex], "--"}
	kargs = append(kargs, strings.Fields(c)...)
	help.RunKubectlCmd(kargs...)
}

func displayLogs(ctIndex int, pdIndex int, nsIndex int) {
	log.Info.Println("🤔 Following logs are found for ", containerSt.containers[ctIndex])
	kargs := []string{"-n", namespacesSt.namespaces[nsIndex], "logs", podSt.pods[pdIndex], "-c", containerSt.containers[ctIndex]}
	help.RunKubectlCmd(kargs...)
}

func DisplayContainers() {
	pdIndex, nsIndex := GetTaggedPods()
	fetchContainers(pdIndex, nsIndex)
	log.Info.Printf("🤔 No namespace provided; will display containers in %v pod for %v namespace", podSt.pods[pdIndex], namespacesSt.namespaces[nsIndex])
	log.Info.Println("😓 Executing command: kubectl get pods -n " + namespacesSt.namespaces[nsIndex])
	log.PrintSpecial(log.GetCurrentFunctionName(), ctrStr)
}
