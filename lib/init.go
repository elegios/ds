package lib

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Initopts struct {
}

var initopts Initopts

func init() {
	Flag.AddCommand("init", "Creates a dotsync folder", "TODO", &initopts)
}

func (i *Initopts) Execute(args []string) error {
	if len(args) != 1 {
		fmt.Println("Init needs exactly one argument, the path to the folder to create")
		os.Exit(1) //TODO: return with error instead
	}

	path, erp := filepath.Abs(args[0])
	d(erp)
	os.MkdirAll(path, permissions)

	erp = os.Symlink(path, dsFolderPath())
	d(erp)

	file, erp := os.Create(listFileName()) //TODO: use helpermethod instead
	d(erp)
	defer file.Close()

	json.NewEncoder(file).Encode(map[string]string{})

	return nil
}
