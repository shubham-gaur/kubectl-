package services

import (
	"fmt"
	"strings"

	help "github.com/shubham-gaur/kubectl++/helper"
	log "github.com/shubham-gaur/kubectl++/logger"
)

var namespacesSt struct {
	namespaces         []string
	numberOfNamespaces int
}
var nsStr string

func fetchNamespaces() {
	kargs := []string{"get", "namespaces"}
	nsStr = help.ExecKubectlCmd(kargs...)
	ns := strings.Fields(nsStr)
	namespacesSt.namespaces = []string{}
	for nsIndex := 3; nsIndex < len(ns); nsIndex = nsIndex + 3 {
		namespacesSt.namespaces = append(namespacesSt.namespaces, ns[nsIndex])
	}
	namespacesSt.numberOfNamespaces = len(namespacesSt.namespaces)
}

func DisplayNamespaces() {
	fetchNamespaces()
	log.PrintSpecial(log.GetCurrentFunctionName(), nsStr)
}

func GetTaggedNamespaces() int {
	fetchNamespaces()
	nsMap := make(map[int]string)
	log.Info.Printf("🤔 Which namespace you are looking for?")
	log.Info.Printf("😀 Below list might help... Available namespaces 👇 ")
	fmt.Printf("%5v﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎\n", "")
	for ns := 0; ns < namespacesSt.numberOfNamespaces; ns++ {
		nsMap[ns] = namespacesSt.namespaces[ns]
		fmt.Printf("%10v👉 Press [%-2v]: %v\n", "", ns, namespacesSt.namespaces[ns])
	}
	fmt.Printf("%5v﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊\n", "")
	var nsIndex int
	help.TakeIntInput(&nsIndex, namespacesSt.numberOfNamespaces)
	return nsIndex
}

func MarkNamespace() {
	log.Info.Printf("🤔 Mark active namespace for operations...")
	var nsIndex int
	nsIndex = GetTaggedNamespaces()
	defaultNamespace := namespacesSt.namespaces[nsIndex]
	options := make(map[int]string)
	options[0] = "select another namespace"
	options[1] = "show all resources in a namespace"
	options[2] = "create namespace"
	options[3] = "delete namespace"
	options[4] = "show all namespaces"
	options[5] = "return to main"
	var opt int
	for {
		log.Info.Printf("😀 Refreshing %v namespace status...", defaultNamespace)
		kargs := []string{"get", "ns", defaultNamespace, "-o", "jsonpath={.status.phase}"}
		status := help.ExecKubectlCmd(kargs...)
		if status == "Active" {
			log.Info.Printf("😀 [Set Namespace]|👉 %v namespace |👌 %v", defaultNamespace, status)
		} else {
			if status == "" {
				status = "Error"
			}
			log.Info.Printf("😦 [Set Namespace]|👉 %v namespace |❌ %v ", defaultNamespace, status)
			log.Error.Printf("🙏 |♻️ Please change namespace!")
		}

		help.Default(options)
		help.TakeIntInput(&opt, len(options))
		switch options[opt] {
		case options[0]:
			nsIndex = GetTaggedNamespaces()
			defaultNamespace = namespacesSt.namespaces[nsIndex]
			log.Info.Printf("😀 Active namespace being set to 👉 %v", defaultNamespace)
		case options[1]:
			log.Info.Printf("🤔 Displaying all resources in %v namespace", defaultNamespace)
			kargs := []string{"get", "all", "-n", defaultNamespace}
			help.RunKubectlCmd(kargs...)
		case options[2]:
			var nsName string
			log.Info.Printf("🤔 What is the namespace name you are looking for?")
			help.TakeStrInput(&nsName)
			kargs := []string{"create", "ns", nsName}
			help.RunKubectlCmd(kargs...)
			log.Info.Printf("[✔️ ] %v create operation successfully 😄 ", nsName)
		case options[3]:
			var confirm string
			log.Info.Printf("😲 You are going to delete a namespace..?")
			log.Info.Printf("🙄  Deleting %v namespace... This action is irreversible. Please confirm by typing 'y' for yes?", defaultNamespace)
			help.TakeStrInput(&confirm)
			if confirm == "y" {
				kargs := []string{"delete", "ns", defaultNamespace}
				help.RunKubectlCmd(kargs...)
				log.Info.Printf("[✔️ ] %v delete operation successfull 😄 ", defaultNamespace)
			} else {
				log.Info.Printf("😐 Confirmation failed please try again?")
			}
		case options[4]:
			DisplayNamespaces()
		case options[5]:
			return
		}
	}
}
