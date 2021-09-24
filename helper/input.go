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

func TakeIntInput(in *int, optLen int) {
	for {
		log.Info.Print(fmt.Sprintln("Please enter: 👉 \033[A"))
		fmt.Print("\033[40;2C")
		fmt.Scan(in)
		if *in < 0 || *in > optLen-1 {
			log.Info.Print(fmt.Sprintln("😵 Not a valid input... 😕 Please try again!"))
		} else {
			return
		}
	}
}

func RetStrBufInput() string {
	log.Info.Print(fmt.Sprintln("Please enter: 👉 \033[A"))
	fmt.Print("\033[40;2C")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	txt := scanner.Text()
	return txt
}

func Default(options map[int]string) {
	log.Info.Println("🤗 Please select the required option...")
	fmt.Printf("%5v﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎\n", "")
	for opt := 0; opt < len(options); opt++ {
		fmt.Printf("%10v👉 [%-2v]: To %v\n", "", opt, options[opt])
	}
	fmt.Printf("%5v﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊\n", "")
}
