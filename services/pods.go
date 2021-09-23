package services

import (
	"fmt"
	"os/exec"
	"strings"

	help "github.com/shubham-gaur/kubectl++/helper"
	log "github.com/shubham-gaur/kubectl++/logger"
)

var podSt struct {
	pods         []string
	numverOfPods int
}

var podStr string

func fetchPods(ns string) {
	var cmd *exec.Cmd
	if ns == "" {
		ns = "default"
	}
	cmd = exec.Command("kubectl", "get", "pods", "-n", ns)
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	podStr = string(stdout)
	podList := strings.Fields(podStr)
	podSt.pods = []string{}
	for pdIndex := 5; pdIndex < len(podList); pdIndex = pdIndex + 5 {
		podSt.pods = append(podSt.pods, podList[pdIndex])
	}
	podSt.numverOfPods = len(podSt.pods)
}

func GetTaggedPods() (int, int) {
	nsIndex := GetTaggedNamespaces()
	fetchPods(namespacesSt.namespaces[nsIndex])
	pdMap := make(map[int]string)
	log.Info.Println("🤔 Which pod you are looking for?")
	log.Info.Println("😀 Below list might help... Available pods 👇 ")
	fmt.Printf("%5v﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎\n", "")
	for pd := 0; pd < podSt.numverOfPods; pd++ {
		pdMap[pd] = podSt.pods[pd]
		fmt.Printf("%10v👉 Press [%-3v]: %v\n", "", pd, podSt.pods[pd])
	}
	fmt.Printf("%5v﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊\n", "")
	var pdIndex int
	help.TakeIntInput(&pdIndex)
	return pdIndex, nsIndex
}

func DisplayPods() {
	nsIndex := GetTaggedNamespaces()
	fetchPods(namespacesSt.namespaces[nsIndex])
	log.Info.Println("🤔 Namespace provided; will display pods in " + namespacesSt.namespaces[nsIndex] + " namespace")
	log.Info.Println("😓 Executing command: kubectl get pods -n " + namespacesSt.namespaces[nsIndex])
	log.PrintSpecial(log.GetCurrentFunctionName(), podStr)
}
