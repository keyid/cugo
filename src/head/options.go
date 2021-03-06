// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package head

type Options struct {
	Number int
}

type Opts func(*Options)

func Number(num int) Opts {
	return func(opt *Options) {
		opt.Number = num
	}
}
