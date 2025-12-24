package main

import (
	"fmt"
	"os"
)

func main() {

	var args []string
	var mapp = make(map[string]string)

	args = os.Args

	// logic is always
	// 0 idx = gh-ndk , 1 idx = --command , 2 idx data , 3 idx --command , 4 data
	// we can see that 2 , 4, 6 are going to be data
	// and 1, 3, 5, are going to be commands
	// it means we can say that , odd idxs ones are command , even idxs onces are data of that previous command

	var cmd string
	var data string

	for idx, str := range args {
		if idx > 0 && idx%2 == 1 {
			// these are commands
			fmt.Printf("Idx: %v Command: %s\n", idx, str)
			cmd = str
		} else if idx > 0 {
			// these are data of previous commands
			fmt.Printf("Idx: %v Data: %s\n", idx, str)
			data = str
			mapp[cmd] = data
		}
	}

	fmt.Println("Map: ", mapp)
}
