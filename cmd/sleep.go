// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package cmd

import (
	"github.com/jcmdln/cugo/lib/help"
	"github.com/jcmdln/cugo/src/sleep"
	"github.com/jcmdln/flagger"
)

type sleepCmd struct {
	name        string
	usage       string
	description string

	help bool
}

func (u *sleepCmd) Prepare(flags *flagger.Flags) {
	u.name, u.usage = "sleep", "DURATION ..."
	u.description = "Suspend execution for an interval of time"

	flags.BoolVar(&u.help, "Show help output", "-h", "--help")
}

func (u *sleepCmd) Action(s []string, flags *flagger.Flags) error {
	var (
		data []string
		err  error
	)

	if data, err = flags.Parse(s); err != nil {
		return err
	}

	if u.help {
		help.Help(u.name, u.usage, u.description, flags)
	}

	sleep.Sleep(data)

	return nil
}

func init() {
	Command.Add("sleep", &sleepCmd{})
}
