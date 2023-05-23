package Commandline

import (
	"bufio"
	"day_2/myWc/Flags"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type Commandline struct {
	Args  []string
	Path  map[string]int
	Flags Flags.Flags
}

func New(args []string) *Commandline {
	return &Commandline{
		Args: args,
	}
}

func (cl *Commandline) CheckFlags() error {
	if cl.Flags.L && cl.Flags.W || cl.Flags.W && cl.Flags.M || cl.Flags.L && cl.Flags.M {
		return fmt.Errorf("invalid combination with flag\nuse only one flag of: %s, %s, %s", "-l", "-m", "-w")
	}
	return nil
}

func (cl *Commandline) Print() error {
	wg := new(sync.WaitGroup)
	if cl.Flags.L {
		cl.PrintL(wg)
	} else if cl.Flags.M {
		cl.PrintM(wg)
	} else {
		cl.PrintW(wg)
	}
	wg.Wait()
	return nil
}

func (cl *Commandline) PrintL(wg *sync.WaitGroup) error {
	for path, _ := range cl.Path {
		wg.Add(1)
		go lineCounter(path, wg)
	}
	return nil
}

func (cl *Commandline) PrintM(wg *sync.WaitGroup) error {
	for path, _ := range cl.Path {
		wg.Add(1)
		go symbCounter(path, wg)
	}
	return nil
}

func (cl *Commandline) PrintW(wg *sync.WaitGroup) error {
	for path, _ := range cl.Path {
		wg.Add(1)
		go wordCounter(path, wg)
	}
	return nil
}

func wordCounter(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(file)
	wordCount := 0
	for fileScanner.Scan() {
		stringSplit := strings.Split(fileScanner.Text(), " ")
		wordCount += len(stringSplit)
	}
	output(wordCount, s)
}

func symbCounter(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(file)
	symbCount := 0
	for fileScanner.Scan() {
		str := fileScanner.Text()
		symbCount += len([]rune(str))
		symbCount++
	}
	symbCount--
	output(symbCount, s)
}

func lineCounter(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	output(lineCount, s)
}

func output(count int, path string) {
	pathSplit := strings.Split(path, "/")
	fmt.Println(count, "\t", pathSplit[len(pathSplit)-1])
}
