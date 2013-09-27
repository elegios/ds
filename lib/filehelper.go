package lib

import (
	"encoding/json"
	"io"
	"os"
	"os/user"
	"path/filepath"
)

func dsFolderPath() string {
	return filepath.Join(homePath(), ".ds-folder")
}

func listFileName() string {
	hostname, _ := os.Hostname()
	return filepath.Join(dsFolderPath(), hostname)
}

func listFile() (*os.File, error) {
	file, err := os.Open(listFileName())
	if err != nil {
		return nil, err
	}

	return file, nil
}

func readListFile() (listfile map[string]string, err error) {
	file, err := listFile()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	err = dec.Decode(&listfile)
	if err != nil {
		return nil, err
	}

	return
}

func writeListFile(listfile map[string]string) (err error) {
	file, err := os.Create(listFileName())
	if err != nil {
		return
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	err = enc.Encode(listfile)
	return
}

func relativeToHome(path string) (relpath string) {
	path, _ = filepath.Abs(path)
	relpath, _ = filepath.Rel(homePath(), path)
	return
}

func homePath() string {
	u, _ := user.Current()
	return u.HomeDir
}

func cp(oldname, newname string) (err error) {
	in, err := os.Open(oldname)
	if err != nil {
		return
	}
	defer in.Close()

	out, err := os.Create(newname)
	if err != nil {
		return
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return
}
