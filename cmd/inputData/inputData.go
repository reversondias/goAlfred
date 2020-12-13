package inputData

import (
	"bufio"
	"fmt"
	"os"
)

func InputString() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter text: ")
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	line := scanner.Text()
	return line
}
