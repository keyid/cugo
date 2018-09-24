package pwd

import (
	"fmt"
	"os"
	"path/filepath"

	er "github.com/jcmdln/cugo/lib/error"
)

var (
	L bool
	P bool
)

func Pwd() {
	var (
		cwd string
		dir string
		err error
	)

	if !L || P {
		cwd, err = os.Getwd()
		er.Error("cugo", err)

		dir, err = filepath.EvalSymlinks(cwd)
		er.Error("cugo", err)
	} else {
		dir = os.Getenv("PWD")
	}

	fmt.Printf("%s\n", dir)
}
