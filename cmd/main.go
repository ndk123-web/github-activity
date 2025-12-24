package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	for idx, str := range args {
		fmt.Println(idx, str)
	}
}
