package lib

import (
	"errors"
	"fmt"
	"path/filepath"
)

type Listopts struct {
}

var listopts Listopts

func init() {
	Flags.AddCommand("list",
		"Lists all files in dotsync",
		"TODO",
		&listopts)
}

func (l *Listopts) Execute(args []string) (err error) {
	if len(args) > 0 {
		return errors.New("list takes no arguments")
	}

	listfile, err := readListFile()
	if err != nil {
		return
	}

	length := 0
	for _, inDSname := range listfile {
		if length < len(inDSname) {
			length = len(inDSname)
		}
	}
	for relpath, inDSname := range listfile {
		fmt.Printf(fmt.Sprintf("%% %ds <- %%s\n", length), inDSname, filepath.Join("~", relpath))
	}

	return nil
}
