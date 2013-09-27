package lib

import (
	"fmt"
	"os"
	"path/filepath"
)

type Initopts struct {
}

var initopts Initopts

func init() {
	Flags.AddCommand("init",
		"Creates a dotsync folder",
		"TODO",
		&initopts)
}

func (i *Initopts) Execute(args []string) error {
	if len(args) != 1 {
		fmt.Println("Init needs exactly one argument, the path to the folder to create")
		os.Exit(1) //TODO: return with error instead
	}

	path, erp := filepath.Abs(args[0])
	d(erp)
	d(os.MkdirAll(path, permissions))

	d(os.Symlink(path, dsFolderPath()))

	writeListFile(map[string]string{})

	return nil
}
