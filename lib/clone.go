package lib

import (
	"errors"
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

func (c *Cloneopt) Execute(args []string) (err error) {
	if len(args) != 2 {
		return errors.New("clone needs exactly two arguments, the ds-folder and the profile to be cloned")
	}

	dsFolder, err := filepath.Abs(args[0])
	if err != nil {
		return
	}
	err = os.Symlink(dsFolder, dsFolderpath)
	if err != nil {
		return
	}

	err = cp(filepath.Join(dsFolder, args[1]), listFilepath)
	if err != nil {
		return
	}

	listfile, err := readListFile()
	if err != nil {
		return
	}

	errs := make([]error, 0)
	for relpath, inDSpath := range listfile {
		abspath := filepath.Join(homepath, filepath.FromSlash(relpath))
		//Remove previous one if it exists, then add a symlink
		os.Remove(abspath)
		err = os.MkdirAll(filepath.Dir(abspath), permissions)
		if err != nil {
			errs = append(errs, err)
			err = nil
			continue
		}
		err = os.Symlink(filepath.Join(dsFolderpath, inDSpath), abspath)
		if err != nil {
			errs = append(errs, err)
			err = nil
		}
	}

	if len(errs) > 0 {
		errColl := ErrorCollection(errs)
		return &errColl
	}

	return nil
}
