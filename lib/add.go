package lib

import (
	"errors"
	"os"
	"path/filepath"
)

type Addopts struct {
	Mine  bool `short:"m" long:"keep-mine" description:"Keep the current file in favor of one previously in dotsync"`
	Other bool `short:"o" long:"keep-other" description:"Keep the file in dotsync instead of the current one"`
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

		_, err = os.Stat(pathInDS)
		switch {
		case a.Mine && a.Other:
			return errors.New("Conflicting options: --keep-mine and --keep-other")
		case a.Mine || os.IsNotExist(err):
			os.Remove(pathInDS)
			os.Rename(path, pathInDS)
		case a.Other:
			os.Remove(path)
		default:
			return errors.New("A file with that name is already in dotsync, use options --keep-mine or --keep-other")
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
