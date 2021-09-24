package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

var (
	Error    *log.Logger
	Warning  *log.Logger
	Info     *log.Logger
	Debug    *log.Logger
	Critical *log.Logger
)

var (
	Black   = Color("\033[1;30m%s\033[0m")
	Red     = Color("\033[1;31m%s\033[0m")
	Green   = Color("\033[1;32m%s\033[0m")
	Yellow  = Color("\033[1;33m%s\033[0m")
	Purple  = Color("\033[1;34m%s\033[0m")
	Magenta = Color("\033[1;35m%s\033[0m")
	Teal    = Color("\033[1;36m%s\033[0m")
	White   = Color("\033[1;37m%s\033[0m")
	IGreen  = Color("\033[1;92m%s\033[0m")
	IYellow = Color("\033[1;93m%s\033[0m")
	ICyan   = Color("\033[1;96m%s\033[0m")
)

var (
	info     = ICyan
	warn     = IYellow
	error    = Red
	debug    = IGreen
	critical = Red
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func init() {
	// file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	// 	log.Info.Print(err)
	// }
	Info = log.New(os.Stdout, info("[kubectl++|INFO.....] "), log.Lmsgprefix)
	Warning = log.New(os.Stdout, warn("[kubectl++|WARNING?!] "), log.Lmsgprefix)
	Error = log.New(os.Stderr, error("[kubectl++|ERROR___X] "), log.Lmsgprefix)
	Debug = log.New(os.Stdout, debug("[kubectl++|DEBUG__??] "), log.Ldate|log.Ltime|log.Lshortfile)
	Critical = log.New(os.Stdout, critical("[kubectl++|CRITICAL_‼] "), log.Ldate|log.Ltime|log.Llongfile)
}

func PrintDebugMap(funcName string, m interface{}) {
	data, err := json.MarshalIndent(m, "", "\t|")
	if err != nil {
		log.Printf("logger:: json parse err: ", err)
		return
	}
	Debug.Printf("%v:\n%v", funcName, string(data))
}

func PrintSpecial(funcName string, output string) {
	fmt.Printf("%5v﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎﹎\n", "")
	fmt.Print(output)
	fmt.Printf("%5v﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊﹊\n", "")
	Info.Printf("[✔️ ] Command executed successfully 😄 ")
}

func GetCurrentFunctionName() string {
	funcName := strings.Split(getFrame(1).Function, "/")
	return fmt.Sprintf("%-10v |", funcName[len(funcName)-1])
}

func GetCallerFunctionName() string {
	return getFrame(2).Function
}

func getFrame(skipFrames int) runtime.Frame {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}
	return frame
}
