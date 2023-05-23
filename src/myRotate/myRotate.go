package main

import (
	"day_2/myRotate/Commandline"
	"day_2/myRotate/Parser"
	"log"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		commandline := Commandline.New(os.Args[1:])
		if err := Parser.Parse(commandline); err != nil {
			log.Fatal(err)
		}
		if err := commandline.Zip(); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("incorrect command usage")
	}
}
