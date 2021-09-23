package services

import (
	"fmt"
	"os/exec"
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
	cmd := exec.Command("kubectl", "get", "namespaces")
	stdout, err := cmd.Output()
	if err != nil {
		panic(err)
	}
	nsStr = string(stdout)
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
	log.Info.Println("🤔 Which namespace you are looking for?")
	log.Info.Println("😀 Below list might help... Available namespaces 👇 ")
	fmt.Printf("%5v﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎\n", "")
	for ns := 0; ns < namespacesSt.numberOfNamespaces; ns++ {
		nsMap[ns] = namespacesSt.namespaces[ns]
		fmt.Printf("%10v👉 Press [%-3v]: %v\n", "", ns, namespacesSt.namespaces[ns])
	}
	fmt.Printf("%5v﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊\n", "")
	var nsIndex int
	help.TakeIntInput(&nsIndex)
	return nsIndex
}
