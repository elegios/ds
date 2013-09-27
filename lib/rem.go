package lib

import (
	"errors"
	"os"
	"path/filepath"
)

type Remopts struct {
	DeleteMode bool `short:"d" long:"delete" description:"Delete the file, don't just stop syncing it"`
}

var remopts Remopts

func init() {
	Flags.AddCommand("rem",
		"Removes the given file from dotsync",
		"TODO",
		&remopts)
}

func (r *Remopts) Execute(args []string) (err error) {
	if len(args) < 1 {
		return errors.New("Rem needs at least one argument: the file(s) to remove from dotsync")
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

		key := filepath.ToSlash(relativeToHome(path))
		pathInDS := filepath.Join(dsFolderpath, listfile[key])

		err = os.Remove(path)
		if err != nil {
			errs = append(errs, err)
			err = nil
			continue
		}
		if r.DeleteMode {
			err = os.Rename(pathInDS, path)
		} else {
			err = cp(pathInDS, path)
		}
		if err != nil {
			errs = append(errs, err)
			err = nil
			continue
		}

		delete(listfile, key)
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
