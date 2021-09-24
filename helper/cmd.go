package helper

import (
	"fmt"
	"os"
	"os/exec"

	log "github.com/shubham-gaur/kubectl++/logger"
)

func ExecKubectlCmd(args ...string) string {
	cmd := exec.Command("kubectl", args...)
	stdout, err := cmd.Output()
	if err != nil {
		log.Warning.Printf("😖 Could not execute command properly %v", err)
	} else {
		log.Info.Printf("[✔️ ] Command executed successfully 😄 ")
	}
	return string(stdout)
}

func RunKubectlCmd(args ...string) {
	cmd := exec.Command("kubectl", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Printf("%5v﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎\n", "")
	err := cmd.Run()
	fmt.Printf("%5v﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊\n", "")
	if err != nil {
		log.Warning.Printf("😖 Could not execute command properly ", err)
	} else {
		log.Info.Printf("[✔️ ] Command executed successfully 😄 ")
	}
}
