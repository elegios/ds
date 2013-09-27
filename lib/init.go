package lib

import (
	"errors"
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

func (i *Initopts) Execute(args []string) (err error) {
	if len(args) != 1 {
		return errors.New("init needs exactly one argument, the ds-folder")
	}

	path, err := filepath.Abs(args[0])
	if err != nil {
		return
	}
	err = os.MkdirAll(path, permissions)
	if err != nil {
		return
	}

	err = os.Symlink(path, dsFolderpath)
	if err != nil {
		return
	}

	err = writeListFile(map[string]string{})

	return
}
