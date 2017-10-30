package main

import (
	"fmt"
	"github.com/markedhero/flagit"
	"os"
)

func main() {
	verbose := false
	printAll := false
	val := 128
	temp := "default"

	flags := flagit.NewFlag()
	flags.Bool(&printAll, []string{"-a"}, "Should I print all of the flag test data")
	flags.Bool(&verbose, []string{"-b", "--boolean"}, "Boolean Flag: if flagged then set to true")
	flags.Int(&val, []string{"-i", "--integer", "-c"}, "Integer Flag: next value expected to be integer")
	flags.String(&temp, []string{"-s", "--string"}, "String Flag: next value expected to be string")

	fmt.Println()
	fmt.Println("Start Value:")
	fmt.Println("  Boolean: ", verbose)
	fmt.Println("  Integer: ", val)
	fmt.Println("  String: ", temp)

	data, err := flags.Parse(os.Args[1:])
	if err == flagit.ErrNoFlags {
		fmt.Println("No flags passed!")
	} else if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println()
	fmt.Println("After Parsing:")
	fmt.Println("  Boolean: ", verbose)
	fmt.Println("  Integer: ", val)
	fmt.Println("  String: ", temp)
	fmt.Println()
	fmt.Println("Returned Data: ")
	for i := range data {
		fmt.Println(data[i])
	}

	if printAll {
		fmt.Println()
		fmt.Println("Printing all Flags: ")
		flags.PrintUsage()
		fmt.Printf("\n")
		fmt.Println("Printing Boolean Flag")
		flags.PrintUsageOf("-b")
		fmt.Println()
		fmt.Println("Printing Integer Flag")
		flags.PrintUsageOf("-i")
		fmt.Println()
		fmt.Println("Printing String Flag")
		flags.PrintUsageOf("--string")
		fmt.Println()
		fmt.Println("Printing invalid Flag")
		flags.PrintUsageOf("--invalid")
	}
}
