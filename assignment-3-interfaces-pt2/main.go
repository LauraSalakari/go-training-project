package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// fmt.Println(os.Args)
	//os.Args reads the list of arguments that are passed to the program when it is executed
	//it is a []string and stores the arguments passed in the command line
	// the args are stored in a temporary directory before the file is actually executed

	//os.Args[1] // this will give you the first string entered (index 0 is the temp directory location)

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
