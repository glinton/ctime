// Ctime is meant to be a linux only inode event trigger. It will update the
// ctime only (to current time), allowing listeners to be notified, but not
// break certain text editors where the file may be open.
//
// ctime's results are similar to `touch -cr file file` where no new file is
// created, and the atime and mtime remain "unmodified."
//
// Usage
//
// Running is as simple as specifying a single file:
//  ctime /path/to/file
// Multiple files may also be specified:
//  ctime file1 file2 file3 file4
//
package main

import (
	"os"
	"syscall"
)

// updateCtime updates the ctime of a file. It does so by reading the current
// atime/mtime from the specified file. (Similar to how `touch -cr file file`
// works)
func updateCtime(fileName string) (err error) {
	times := syscall.Stat_t{}

	err = syscall.Stat(fileName, &times)
	// ignore this error
	if err != nil {
		os.Stderr.WriteString("Failed to stat '" + fileName + "': " + err.Error())
		return nil
	}

	err = syscall.UtimesNano(fileName, []syscall.Timespec{times.Atim, times.Mtim})
	if err != nil {
		os.Stderr.WriteString("Failed to utime: " + err.Error())
		return err
	}
	return nil
}

// main accepts file names as arguments to update the ctime on
func main() {
	if len(os.Args) >= 2 {
		for i := range os.Args[1:] {
			updateCtime(os.Args[i+1])
		}
	}
}
