package main

import (
	"day_2/myWc/Commandline"
	"day_2/myWc/Parser"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		commandline := Commandline.New(os.Args[1:])
		if err := Parser.Parse(commandline); err != nil {
			log.Fatal(err)
		}
		if err := commandline.CheckFlags(); err != nil {
			log.Fatal(err)
		}
		if err := commandline.Print(); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("incorrect command usage")
	}
}
