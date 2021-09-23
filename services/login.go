package services

import (
	help "github.com/shubham-gaur/kubectl++/helper"
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
	log.Info.Printf("🤔 %v namespace provided; will login inside %v container in %v pod", namespace, container, pod)
	log.Info.Printf("😓 Executing command: kubectl -n %v exec -it %v -c %v --bash", namespace, pod, container)
	kargs := []string{"-n", namespace, "exec", "-it", pod, "-c", container, "--", "bash"}
	help.RunKubectlCmd(kargs...)
}
