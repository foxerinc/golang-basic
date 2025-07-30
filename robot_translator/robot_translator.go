package main


import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Wellcome To Robot Translator")
	inputYourCommand := "RRL"
	stringSLice := strings.Split(inputYourCommand, "")
	stringSLice = append(stringSLice, "")
	output := ""

	counter := 0
	prev := ""
	for _, command := range stringSLice {
		move := ""
		if counter > 1 {
			move += "s"
		}
		if prev != command && prev != "" {
			switch prev {
			case "R":
				output += fmt.Sprintf("Move right %d time%s\n", counter,move)
				counter = 0
			case "L":
				output += fmt.Sprintf("Move left %d time%s\n", counter,move)
				counter = 0
			case "A":
				output += fmt.Sprintf("Move advance %d time%s\n", counter,move)
				counter = 0
			}
		}
		switch command {
		case "R":
			counter += 1
			prev = "R"
		case "L":
			counter += 1
			prev = "L"

		case "A":
			counter += 1
			prev = "A"
		}
	}
	println(output)
}