package lib

import (
	"fmt"
	"os"
	"path/filepath"
)

type Addopts struct {
}

var addopts Addopts

func init() {
	Flag.AddCommand("add", "Adds the given file to dotsync", "TODO", &addopts)
}

func (a *Addopts) Execute(args []string) error {
	if len(args) != 1 {
		fmt.Println("Add needs exactly one argument, the file to be added to dotsync")
		os.Exit(1)
	}

	path, erp := filepath.Abs(args[0])
	d(erp)
	pathInDS := filepath.Join(dsFolderPath(), filepath.Base(path))

	os.Rename(path, pathInDS) //TODO: deal with errors by renaming files if necessary
	os.Symlink(pathInDS, path)

	listfile, erp := readListFile()
	d(erp)

	listfile[filepath.ToSlash(relativeToHome(path))] = filepath.Base(pathInDS)

	d(writeListFile(listfile))

	return nil
}
