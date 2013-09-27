package lib

import (
	"fmt"
	"os"
	"path/filepath"
)

type Remopts struct {
	DeleteMode bool `short:"d" long:"delete" description:"Delete the file, don't just stop syncing it"`
}

var remopts Remopts

func init() {
	Flag.AddCommand("rem",
		"Removes the given file from dotsync",
		"TODO",
		&remopts)
}

func (r *Remopts) Execute(args []string) error {
	if len(args) != 1 {
		fmt.Println("Rem needs exactly one argument, the file to remove from dotsync")
		os.Exit(1)
	}

	path, erp := filepath.Abs(args[0])
	d(erp)

	listfile, erp := readListFile()
	d(erp)

	key := filepath.ToSlash(relativeToHome(path))
	pathInDS := filepath.Join(dsFolderPath(), listfile[key])

	erp = os.Remove(path)
	d(erp)
	if r.DeleteMode {
		d(os.Rename(pathInDS, path))
	} else {
		d(cp(pathInDS, path))
	}

	delete(listfile, key)

	d(writeListFile(listfile))

	return nil
}
