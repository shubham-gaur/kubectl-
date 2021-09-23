package helper

import (
	"bufio"
	"fmt"
	"os"

	log "github.com/shubham-gaur/kubectl++/logger"
)

func TakeStrInput(in *string) {
	log.Info.Print(fmt.Sprintln("Please enter: 👉 \033[A"))
	fmt.Print("\033[40;2C")
	fmt.Scan(in)
}

func TakeIntInput(in *int) {
	log.Info.Print(fmt.Sprintln("Please enter: 👉 \033[A"))
	fmt.Print("\033[40;2C")
	fmt.Scan(in)
}

func RetStrBufInput() string {
	log.Info.Print(fmt.Sprintln("Please enter: 👉 \033[A"))
	fmt.Print("\033[40;2C")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	txt := scanner.Text()
	return txt
}

func Default(options map[int]string) int {
	log.Info.Println("🤗 Please select the required option...")
	fmt.Printf("%5v﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎\n", "")
	for opt, val := range options {
		fmt.Printf("%10v 👉 [%-3v]: To [%v]\n", "", opt, val)
	}
	fmt.Printf("%5v﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊\n", "")
	var opt int
	TakeIntInput(&opt)
	return opt
}
