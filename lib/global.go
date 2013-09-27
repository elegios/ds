package lib

import (
	"github.com/jessevdk/go-flags"
)

const (
	permissions = 0700
)

var Flags = flags.NewParser(nil, flags.Default)

func d(err error) {
	if err != nil {
		panic(err)
	}
}
