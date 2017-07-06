[![GoDoc](https://godoc.org/github.com/glinton/ctime?status.svg)](https://godoc.org/github.com/glinton/ctime)

# ctime

Ctime is meant to be a linux only inode event trigger. It will only update the ctime (to system time), allowing listeners to be notified, but not break certain text editors where the file may be open.

ctime's results are similar to `touch -cr file file` where no new file is created, and the atime and mtime remain "unmodified."

### Usage

Running is as simple as specifying a single file:
```
ctime /path/to/file
```

Multiple files may also be specified:
```
ctime file1 file2 file3 file4
```
