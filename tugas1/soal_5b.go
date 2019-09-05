package main

import "flag"
import "fmt"

func main() {

	wordPtr := flag.String("word", "foo", "a string")

	
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	flag.Parse()

	fmt.Print("word:", *wordPtr)
	fmt.Print("numb:", *numbPtr)
	fmt.Print("fork:", *boolPtr)
	fmt.Print("svar:", svar)
	fmt.Print("tail:", flag.Args())
}