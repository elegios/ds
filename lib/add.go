package lib

import (
	"errors"
	"os"
	"path/filepath"
)

type Addopts struct {
}

var addopts Addopts

func init() {
	Flags.AddCommand("add",
		"Adds the given file to dotsync",
		"TODO",
		&addopts)
}

func (a *Addopts) Execute(args []string) (err error) {
	if len(args) < 1 {
		return errors.New("Add needs at least one argument: the file(s) to be added to dotsync")
	}

	listfile, err := readListFile()
	if err != nil {
		return
	}

	errs := make([]error, 0)
	for _, arg := range args {
		path, err := filepath.Abs(arg)
		if err != nil {
			errs = append(errs, err)
			err = nil
			continue
		}
		pathInDS := filepath.Join(dsFolderpath, filepath.Base(path))

		err = os.Rename(path, pathInDS) //TODO: deal with errors by renaming files if necessary
		if err != nil {
			errs = append(errs, err)
			err = nil
			continue
		}
		err = os.Symlink(pathInDS, path)
		if err != nil {
			errs = append(errs, err)
			err = nil
			continue
		}

		listfile[filepath.ToSlash(relativeToHome(path))] = filepath.Base(pathInDS)
	}

	err = writeListFile(listfile)
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		errColl := ErrorCollection(errs)
		return &errColl
	}

	return nil
}
