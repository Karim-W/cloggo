package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {
	errorFormat := color.New(color.BgRed, color.FgBlack).SprintFunc()
	warningFormat := color.New(color.BgYellow, color.FgBlack).SprintFunc()
	infoFormat := color.New(color.FgBlack, color.FgCyan).SprintFunc()
	lines := make(chan string, 1000)
	go func() {
		for line := range lines {
			txt := strings.ToUpper(line)
			if strings.Contains(txt, "ERROR") {
				fmt.Println(errorFormat(line))
			} else if strings.Contains(txt, "WARN") {
				fmt.Println(warningFormat(line))
			} else if strings.Contains(txt, "INFO") {
				fmt.Println(infoFormat(line))
			} else {
				fmt.Println(line)
			}
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines <- scanner.Text()
	}
}
