package services

import (
	"os"
	"os/exec"
	"time"

	help "github.com/shubham-gaur/kubectl++/helper"
	log "github.com/shubham-gaur/kubectl++/logger"
)

func CollectLogsForPod() {

}

func CollectLogForContainer() {
	ctIndex, pdIndex, nsIndex := GetTaggedContainers()
	namespace := namespacesSt.namespaces[nsIndex]
	pod := podSt.pods[pdIndex]
	container := containerSt.containers[ctIndex]

	// open the out file for writing
	outfile, err := os.Create("./" + container + ".log")
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	log.Info.Println("🤔 For how long I should collect logs ??? Please be generous time in seconds 😅")
	var timeout int
	help.TakeIntInput(&timeout)

	var cmd *exec.Cmd
	log.Info.Println("🤔 Collecting log of " + container + " for " + pod + " pod in " + namespace + " namespace")
	log.Info.Println("😓 Executing command: timeout 3 kubectl -n " + namespace + " logs -f " + pod + " -c " + container + " > " + container + ".log")
	cmd = exec.Command("kubectl", "logs", "-f", pod, "-n", namespace, "-c", container)
	cmd.Stdin = os.Stdin
	cmd.Stdout = outfile
	cmd.Stderr = os.Stderr
	done := make(chan error, 1)
	err = cmd.Start()
	if err != nil {
		log.Warning.Println("😖 Could not execute properly ", err)
	}
	go func() {
		done <- cmd.Wait()
	}()
	go func() {
		for {
			select {
			case <-time.After(time.Duration(timeout) * time.Second):
				if err := cmd.Process.Kill(); err != nil {
					log.Critical.Println("😢 failed to kill process: ", err)
				}
				log.Info.Printf("I get tired after fetching logs for %v seconds 😰", timeout)
				log.Warning.Println("Limited support for logging 😢")
				return
			case err := <-done:
				if err != nil {
					log.Critical.Println("process finished with error ", err)
				}
				log.Info.Println("[✔️ ] Command executed successfully 😄 ")
			}
		}
	}()
}
