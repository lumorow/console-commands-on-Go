package Parser

import (
	"day_2/myWc/Commandline"
	"fmt"
	"os"
)

func Parse(cl *Commandline.Commandline) error {
	cl.Path = make(map[string]int)
	for _, arg := range cl.Args {
		switch arg {
		case "-l":
			if cl.Flags.L {
				return fmt.Errorf("flag %s already exist", arg)
			}
			cl.Flags.L = true
		case "-m":
			if cl.Flags.M {
				return fmt.Errorf("flag %s already exist", arg)
			}
			cl.Flags.M = true
		case "-w":
			if cl.Flags.W {
				return fmt.Errorf("flag %s already exist", arg)
			}
			cl.Flags.W = true
		default:
			if arg[0] != '-' {
				if _, err := os.Stat(arg); os.IsNotExist(err) {
					return err
				}
				cl.Path[arg]++
			} else {
				return fmt.Errorf("unknown flag %s", arg)
			}
		}
	}
	return nil
}
