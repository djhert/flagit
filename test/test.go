package main

import (
	"fmt"
	"github.com/markedhero/flagit"
	"os"
)

func main() {
	verbose := false
	recursive := false
	parent := false
	val := 128
	temp := "default"

	flags := new(flagit.Flag)
	flags.AddBoolFlag(&verbose, "-v", "--verbose")
	flags.AddBoolFlag(&recursive, "-r", "--recursive")
	flags.AddBoolFlag(&parent, "-p", "--parent")
	flags.AddIntFlag(&val, "-i", "--integer")
	flags.AddStringFlag(&temp, "-s", "--string")

	fmt.Println("Before")
	fmt.Println("Verbose: ", verbose)
	fmt.Println("recursive: ", recursive)
	fmt.Println("string: ", temp)
	err := flags.ParseFlags(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println()
	fmt.Println("After")
	fmt.Println("Verbose: ", verbose)
	fmt.Println("recursive: ", recursive)
	fmt.Println("Value: ", val)
	fmt.Println("string: ", temp)
}
