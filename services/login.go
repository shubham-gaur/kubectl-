package services

import (
	"os"
	"os/exec"

	log "github.com/shubham-gaur/kubectl++/logger"
)

func Login(args ...int) {
	var ctIndex, pdIndex, nsIndex int
	if len(args) > 0 {
		ctIndex = args[0]
		pdIndex = args[1]
		nsIndex = args[2]
	} else {
		ctIndex, pdIndex, nsIndex = GetTaggedContainers()
	}
	namespace := namespacesSt.namespaces[nsIndex]
	pod := podSt.pods[pdIndex]
	container := containerSt.containers[ctIndex]
	var cmd *exec.Cmd
	log.Info.Printf("🤔 %v namespace provided; will login inside %v container in %v pod", namespace, container, pod)
	log.Info.Println("😓 Executing command: kubectl exec")
	cmd = exec.Command("kubectl", "-n", namespace, "exec", "-it", pod, "-c", container, "--", "bash")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Warning.Println("😖 Could not login properly ", err)
	}
}
