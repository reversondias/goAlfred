package inputData

import (
	"fmt"
)

func InputString() string {
	fmt.Print("Enter text: ")
	var input string
	fmt.Scanln(&input)
	return input
}
