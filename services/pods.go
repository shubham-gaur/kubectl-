package services

import (
	"fmt"
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
	if ns == "" {
		ns = "default"
	}
	kargs := []string{"get", "pods", "-n", ns}
	podStr = help.ExecKubectlCmd(kargs...)
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
		fmt.Printf("%10v👉 Press [%-2v]: %v\n", "", pd, podSt.pods[pd])
	}
	fmt.Printf("%5v﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊\n", "")
	var pdIndex int
	help.TakeIntInput(&pdIndex)
	return pdIndex, nsIndex
}

func DisplayPods() {
	nsIndex := GetTaggedNamespaces()
	fetchPods(namespacesSt.namespaces[nsIndex])
	log.Info.Printf("🤔 %v namespace provided; will display pods in %v namespace", namespacesSt.namespaces[nsIndex], namespacesSt.namespaces[nsIndex])
	log.Info.Printf("😓 Executing command: kubectl get pods -n %v", namespacesSt.namespaces[nsIndex])
	log.PrintSpecial(log.GetCurrentFunctionName(), podStr)
}
