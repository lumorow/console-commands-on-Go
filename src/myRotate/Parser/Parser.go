package Parser

import (
	"day_2/myRotate/Commandline"
	"fmt"
	"os"
)

func Parse(cl *Commandline.Commandline) error {
	cl.Path = make(map[string]int)
	for _, arg := range cl.Args {
		switch arg {
		case "-a":
			if cl.Flags.A {
				return fmt.Errorf("flag %s already exist", arg)
			}
			cl.Flags.A = true
		default:
			if arg[0] != '-' {
				fileInfo, err := os.Stat(arg)
				if os.IsNotExist(err) {
					return err
				}
				if !fileInfo.IsDir() {
					fmt.Errorf("%s - is not a directory", arg)
				}
				if cl.Arch == "" && cl.Flags.A {
					cl.Arch = arg
				} else {
					cl.Path[arg]++
				}
			} else {
				return fmt.Errorf("unknown flag %s", arg)
			}
		}
	}
	return nil
}
