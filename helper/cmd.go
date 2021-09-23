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
		panic(err)
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
		log.Warning.Println("😖 Could not execute command properly ", err)
	} else {
		log.Info.Println("[✔️ ] Command executed successfully 😄 ")
	}
}
