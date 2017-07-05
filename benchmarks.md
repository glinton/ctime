# Benchmarks (simple)

#### Comparing `ctime` to `touch` when "updating" 1000 files

Control file info
```
$ stat file1000
  File: 'file1000'
  Size: 0         	Blocks: 0          IO Block: 4096   regular empty file
Device: 10305h/66309d	Inode: 13766558    Links: 1
Access: (0644/-rw-r--r--)  Uid: ( 1000/   fives)   Gid: ( 1000/   fives)
Access: 2017-07-05 17:22:07.074249642 -0600
Modify: 2017-07-05 17:22:07.074249642 -0600
Change: 2017-07-05 17:22:49.645918416 -0600
 Birth: -
```

Timing 1000 `touch` commands
```
# touch -c -r file1 file1 && touch -c -r file10 file10 && ...
$ time bash files

real	0m0.656s
user	0m0.036s
sys		0m0.112s
```

Verify it worked
```
$ stat file1000
  File: 'file1000'
  Size: 0         	Blocks: 0          IO Block: 4096   regular empty file
Device: 10305h/66309d	Inode: 13766558    Links: 1
Access: (0644/-rw-r--r--)  Uid: ( 1000/   fives)   Gid: ( 1000/   fives)
Access: 2017-07-05 17:22:07.074249642 -0600
Modify: 2017-07-05 17:22:07.074249642 -0600
Change: 2017-07-05 17:26:48.576203529 -0600
 Birth: -
```

Timing ctime updating the same 1000 files
```
$ time ctime *

real	0m0.007s
user	0m0.004s
sys		0m0.004s
```

##### `ctime` is more than 25 times faster than `touch`!

Verify it worked
```
$ stat file1000
  File: 'file1000'
  Size: 0         	Blocks: 0          IO Block: 4096   regular empty file
Device: 10305h/66309d	Inode: 13766558    Links: 1
Access: (0644/-rw-r--r--)  Uid: ( 1000/   fives)   Gid: ( 1000/   fives)
Access: 2017-07-05 17:22:07.074249642 -0600
Modify: 2017-07-05 17:22:07.074249642 -0600
Change: 2017-07-05 17:27:03.696101787 -0600
 Birth: -
```

Timing 1000 `touch` commands (linked with `;` rather than `&&`)
```
# touch -c -r file1 file1; touch -c -r file10 file10; ...
$ time bash files

real	0m0.646s
user	0m0.040s
sys	0m0.100s
```

Verify it worked
```
$ stat file1000
  File: 'file1000'
  Size: 0         	Blocks: 0          IO Block: 4096   regular empty file
Device: 10305h/66309d	Inode: 13766558    Links: 1
Access: (0644/-rw-r--r--)  Uid: ( 1000/   fives)   Gid: ( 1000/   fives)
Access: 2017-07-05 17:22:07.074249642 -0600
Modify: 2017-07-05 17:22:07.074249642 -0600
Change: 2017-07-05 17:29:51.971010684 -0600
 Birth: -
```

*Tests run on laptop running Ubuntu 16.10
