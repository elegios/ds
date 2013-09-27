package lib

import (
	"fmt"
	"os"
	"path/filepath"
)

type Cloneopt struct {
}

var cloneopt Cloneopt

func init() {
	Flags.AddCommand("clone",
		"Clones the supplied ds-folder with the given profile",
		"TODO",
		&cloneopt)
}

func (c *Cloneopt) Execute(args []string) error {
	if len(args) != 2 {
		fmt.Println("clone needs exactly two arguments, the ds-folder and the profile to be cloned")
		os.Exit(1)
	}

	dsFolder, erp := filepath.Abs(args[0])
	d(erp)
	d(os.Symlink(dsFolder, dsFolderPath()))

	d(cp(filepath.Join(dsFolder, args[1]), listFileName()))

	listfile, erp := readListFile()
	d(erp)

	home := homePath()
	dsFolder = dsFolderPath()
	for relpath, inDSpath := range listfile {
		abspath := filepath.Join(home, relpath)
		//Remove previous one if it exists, then add a symlink
		os.Remove(abspath)
		d(os.Symlink(filepath.Join(dsFolder, inDSpath), abspath))
	}

	return nil
}
