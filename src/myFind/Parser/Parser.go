package Parser

import (
	"day_2/myFind/Commandline"
	"fmt"
	"os"
)

func Parse(cl *Commandline.Commandline) error {
	cl.Path = cl.Args[len(cl.Args)-1]
	if _, err := os.Stat(cl.Path); os.IsNotExist(err) {
		return err
	}
	for _, arg := range cl.Args[:len(cl.Args)-1] {
		switch arg {
		case "-sl":
			if cl.Flags.Sl {
				return fmt.Errorf("flag %s already exist", arg)
			}
			cl.Flags.Sl = true
		case "-d":
			if cl.Flags.D {
				return fmt.Errorf("flag %s already exist", arg)
			}
			cl.Flags.D = true
		case "-f":
			if cl.Flags.F {
				return fmt.Errorf("flag %s already exist", arg)
			}
			cl.Flags.F = true
		case "-ext":
			if cl.Flags.Ext {
				return fmt.Errorf("flag %s already exist", arg)
			}
			cl.Flags.Ext = true
		default:
			if cl.Flags.Ext && arg[0] != '-' {
				cl.Ext = arg
			} else {
				return fmt.Errorf("unknown flag %s", arg)
			}
		}
	}
	if cl.Flags.Ext == true && cl.Ext == "" {
		return fmt.Errorf("missing search extension after: %s", "-ext")
	}
	return nil
}
