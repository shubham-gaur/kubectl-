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
	log.Info.Printf("🤔 Which pod you are looking for?")
	log.Info.Printf("😀 Below list might help... Available pods 👇 ")
	fmt.Printf("%5v﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎\n", "")
	for pd := 0; pd < podSt.numverOfPods; pd++ {
		pdMap[pd] = podSt.pods[pd]
		fmt.Printf("%10v👉 Press [%-2v]: %v\n", "", pd, podSt.pods[pd])
	}
	fmt.Printf("%5v﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊\n", "")
	var pdIndex int
	help.TakeIntInput(&pdIndex, podSt.numverOfPods)
	return pdIndex, nsIndex
}

func DisplayPods() {
	nsIndex := GetTaggedNamespaces()
	fetchPods(namespacesSt.namespaces[nsIndex])
	log.Info.Printf("🤔 %v namespace provided; will display pods in %v namespace", namespacesSt.namespaces[nsIndex], namespacesSt.namespaces[nsIndex])
	log.Info.Printf("😓 Executing command: kubectl get pods -n %v", namespacesSt.namespaces[nsIndex])
	log.PrintSpecial(log.GetCurrentFunctionName(), podStr)
}

func MarkPod() {
	log.Info.Printf("🤔 Mark active pod for operations...")
	var pdIndex int
	var nsIndex int
	pdIndex, nsIndex = GetTaggedPods()
	defaultPod := podSt.pods[pdIndex]
	defaultNamespace := namespacesSt.namespaces[nsIndex]
	options := make(map[int]string)
	options[0] = "select another pod"
	options[1] = "show all pod in current namespace"
	options[2] = "show all pod in another namespace"
	options[3] = "show all containers in a pod"
	options[4] = "describe pod"
	options[5] = "create pod"
	options[6] = "delete pod"
	options[7] = "return to main"
	var opt int
	for {
		log.Info.Printf("😀 Refreshing %v pod status...", defaultPod)
		kargs := []string{"get", "pod", defaultPod, "-n", defaultNamespace, "-o", "jsonpath={.status.phase}"}
		status := help.ExecKubectlCmd(kargs...)
		if status == "Active" {
			log.Info.Printf("😀 [Set Pod]|👉 %v pod |👉 %v namespace |👌 %v", defaultPod, defaultNamespace, status)
		} else {
			if status == "" {
				status = "Error"
			}
			log.Info.Printf("😦 [Set Pod]|👉 %v pod |👉 %v namespace |❌ %v", defaultPod, defaultNamespace, status)
			log.Error.Printf("🙏 |♻️ Please change pod!")
		}
		help.Default(options)
		help.TakeIntInput(&opt, len(options))
		switch options[opt] {
		case options[0]:
			pdIndex, nsIndex = GetTaggedPods()
			log.Info.Printf("😀 Active pod being set to 👉 %v pod 👉 %v namespace", defaultPod, defaultNamespace)
		case options[1]:
			log.Info.Printf("🤔 Displaying all pods in %v namespace", defaultNamespace)
			kargs := []string{"get", "pods", "-n", defaultNamespace}
			help.RunKubectlCmd(kargs...)
		case options[2]:
			log.Info.Printf("🤔 Displaying all containers in %v pod %v namespace", defaultPod, defaultNamespace)
			kargs := []string{"get", "pods", defaultPod, "-n", defaultNamespace, "-o", "jsonpath={.spec.containers[*].name}"}
			help.RunKubectlCmd(kargs...)
		case options[3]:
			log.Info.Printf("🤔 Displaying all containers in %v pod %v namespace", defaultPod, defaultNamespace)
			DisplayContainers(pdIndex, nsIndex)
		case options[4]:
			log.Info.Printf("🤔 Describing %v pod in %v namespace", defaultPod, defaultNamespace)
			kargs := []string{"describe", "pod", defaultPod, "-n", defaultNamespace}
			help.RunKubectlCmd(kargs...)
		case options[5]:
			log.Info.Printf("😅 Good things take time ⌛ to commit! Coming soon...")
		case options[6]:
			var confirm string
			log.Info.Printf("😲 You are going to delete a pod..?")
			log.Info.Printf("🙄 Deleting %v pod... This action is irreversible. Please confirm by typing 'y' for yes?", defaultPod)
			help.TakeStrInput(&confirm)
			if confirm == "y" {
				kargs := []string{"delete", "pod", defaultPod, "-n", defaultNamespace}
				help.RunKubectlCmd(kargs...)
				log.Info.Printf("[✔️ ] %v pod delete operation successfull in %v namespace 😄 ", defaultPod, defaultNamespace)
			} else {
				log.Info.Printf("😐 Confirmation failed please try again?")
			}
		case options[7]:
			return
		}
	}
}
