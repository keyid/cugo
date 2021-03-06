// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package touch

import (
	"os"
	"syscall"
	"time"

	ex "github.com/jcmdln/cugo/lib/exists"
)

// Touch ...
func (opt *Options) Touch(operands []string) error {
	var (
		operand    string
		finfo      os.FileInfo
		fstat      *syscall.Stat_t
		rinfo      os.FileInfo
		rstat      *syscall.Stat_t
		accessTime time.Time
		modifyTime time.Time
		date       time.Time
		err        error
	)

	for _, operand = range operands {
		if !opt.Create {
			if _, err = os.Create(operand); err != nil {
				return err
			}
		}

		if ex.Exists(operand) != nil {
			finfo, _ = os.Stat(operand)
			fstat = finfo.Sys().(*syscall.Stat_t)

			if len(opt.Reference) > 0 {
				if rinfo, err = os.Stat(opt.Reference); err != nil {
					return err
				}

				rstat = rinfo.Sys().(*syscall.Stat_t)
				accessTime = time.Unix(rstat.Atim.Sec, rstat.Atim.Nsec)
				modifyTime = rinfo.ModTime()
			} else if len(opt.Date) > 0 {
				if date, err = time.Parse(time.RFC3339Nano, opt.Date); err != nil {
					return err
				}

				accessTime = date
				modifyTime = date
			} else {
				accessTime = time.Unix(fstat.Atim.Sec, fstat.Atim.Nsec)
				modifyTime = finfo.ModTime()
			}

			if err = os.Chtimes(operand, accessTime, modifyTime); err != nil {
				return err
			}
		}
	}

	return nil
}
