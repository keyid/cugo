// Copyright 2018 Johnathan C Maudlin
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// remove directories
//
// SYNOPSIS
//
//     rmdir [-p] directory ...
//
// DESCRIPTION
//
// rmdir removes the directory entry specified by each directory
// argument, provided it is empty.
//
// The options are as follow:
//
//     -p        Remove empty parent directories.
//
// Arguments are processed in the provided order and will exit on any
// error, leaving following directories intact.
//
// SEE ALSO
//
//     https://golang.org/pkg/os/#Stat
//     https://golang.org/pkg/os/#Remove
//     https://golang.org/pkg/path/filepath/#Walk
//
// REFERENCES
//
//     http://man.openbsd.org/rmdir
//     http://pubs.opengroup.org/onlinepubs/9699919799/utilities/rmdir.html
package rmdir
