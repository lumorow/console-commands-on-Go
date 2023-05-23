package Commandline

import (
	"day_2/myFind/Flags"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type Commandline struct {
	Args  []string
	Ext   string
	Path  string
	Flags Flags.Flags
}

func New(args []string) *Commandline {
	return &Commandline{
		Args: args,
	}
}

func (cl *Commandline) CheckFlags() error {
	if cl.Flags.Ext && (cl.Flags.D || cl.Flags.Sl || !cl.Flags.F) {
		return fmt.Errorf("invalid combination with flag: %s\ncombinate flag: %s only with %s", "-ext", "-ext", "-f")
	}
	return nil
}

func (cl *Commandline) Print() error {
	filepath.Walk(cl.Path, func(wPath string, info os.FileInfo, err error) error {
		if wPath == cl.Path {
			return nil
		}
		cl.condition(info, wPath)
		return nil
	})
	return nil
}

func (cl *Commandline) condition(info os.FileInfo, wPath string) error {
	mode := info.Mode()
	if cl.Flags.Sl && mode&fs.ModeSymlink != 0 {
		printSymlink(wPath)
		return nil
	}
	if cl.Flags.Ext && !info.IsDir() {
		cl.printFileExt(wPath)
		return nil
	}
	if cl.Flags.D && mode.IsDir() || cl.Flags.F && !mode.IsDir() && mode&fs.ModeSymlink == 0 {
		printFileDir(wPath)
	}
	return nil
}

func (cl *Commandline) printFileExt(wPath string) {
	extension := filepath.Ext(wPath)
	if extension != "" {
		if extension[1:] == cl.Ext {
			fmt.Println(wPath)
		}
	}
}

func printFileDir(wPath string) {
	fmt.Println(wPath)
}

func printSymlink(wPath string) {
	symbolicLink, err := os.Readlink(wPath)
	if err != nil {
		fmt.Println(wPath, "-> [broken]")
	} else {
		fmt.Println(wPath, "->", symbolicLink)
	}
}
