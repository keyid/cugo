// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package basename

import (
	"fmt"
	"path/filepath"
)

func Basename(operand string) error {
	fmt.Printf("%s\n", filepath.Base(operand))
	return nil
}
