package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	signatures, err := os.Open("readme.md")
	if err != nil {
		fmt.Println(err)
		return
	}

	reader := bufio.NewReader(signatures)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		fmt.Println(line)
	}

	defer signatures.Close()
}
