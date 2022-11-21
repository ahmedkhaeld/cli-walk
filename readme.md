# walk a cli tool
the `walk` tool has two main goals: descending into a directory tree
to look for files that match a specified criteria and executing an action on these files.

* `-root`: source; the root of the directory tree to start the search. The default is the current dir
* `-list`: action; list files found by the tool. when specified, no other action will be executed
* `-del`: action; delete the  listed files
* `-archive`:action; archive file before deletion to free some space
* `-ext` : filter; file extension to search. when specified, the tool will only match files with this extension
* `-size`: filter; minimum file size in bytes. when specified, the tool will only match files whose size is larger than this value
---

####Filter by size and extension
`$go run . -root /tmp/testdir/ -ext log`<br>
`$go run . -root /tmp/testdir/ -size 0`
* showcase various scenarios; 
>all files size is zero except file1.txt =12bytes

```
/tmp/testdir/
├── file1.txt
├── logs
│  ├── log1.log
└── text
    ├── text1.txt
```
1. `$ go run . -root /tmp/testdir/ `<br>
if no options are set, then by default list all the path directories and files
``` 
/tmp/testdir/file1.txt
/tmp/testdir/logs/log1.log
/tmp/testdir/text/text1.txt

```
2. `$go run . -root /tmp/testdir/ -size 0`<br>
list all files with minimum size 0, means all files
``` 
/tmp/testdir/file1.txt
/tmp/testdir/logs/log1.log
/tmp/testdir/text/text1.txt

```
3. `$go run . -root /tmp/testdir/ -size 12`<br>
list files with minimum size 12 bytes<br>
>Note: file1.txt size is 12 bytes<br>
so, from 1 to 12 bytes will list the file1.txt
``` 
/tmp/testdir/file1.txt
```
4. `$go run . -root /tmp/testdir/ -size 13`<br>
      list files with minimum size 13 bytes; No files

5. `$go run . -root /tmp/testdir/ -ext log`<br>
list all file extension .log
``` 
/tmp/testdir/logs/log1.log

```
6. `$go run . -root /tmp/testdir/ -ext txt` <br>
list all files extension .txt
```  
/tmp/testdir/file1.txt
/tmp/testdir/text/text1.txt
```
---
### delete files
`go run . -root /tmp/testdir/ -ext .log -del`
<br> after deletion if we list<br>
`go run . -root /tmp/testdir/   -list`

``` 
/tmp/testdir/file1.txt
/tmp/testdir/text/text1.txt
```
##### feedback the user with the deleted files
store the deleted files in a log file
<br> using` go run . root /tmp/testdir/ -ext .txt -log deleted_files -del`

use `cat deleted_files.log` to display the logs

---
### archive files
copy the -root dir to be archived into -archive, only the file with .go extension 
`go run . -root /tmp/gomisc/ -ext .go -archive /tmp/gomisc_bkp`
