# walk cli tool
the `walk` tool has two main goals: descending into a directory tree
to look for files that match a specified criteria and executing an action on these files.

* `-root`: where; the root of the directory tree to start the search. The default is the current dir
* `-list`: action; list files found by the tool. when specified, no other action will be executed
* `-ext` : filter; file extension to search. when specified, the tool will only match files with this extension
* `-size`: filter; minimum file size in bytes. when specified, the tool will only match files whose size is larger than this value
---
* testdata Dir tree used for test cases
``` 
$tree testdata
testdata
├── dir2
│ └── script.sh
└── dir.log  size=12 bytes

```
---
* showcase various scenarios; all files size is zero except file1.txt =12bytes

```
/tmp/testdir/
├── file1.txt
├── logs
│  ├── log1.log
│  ├── log2.log
│  └── log3.log
└── text
    ├── text1.txt
    ├── text2.txt
    └── text3.txt

```
1. `$ go run . -root /tmp/testdir/ `<br>
if no options are set, then by default list all the path directories and files
``` 
/tmp/testdir/file1.txt
/tmp/testdir/logs/log1.log
/tmp/testdir/logs/log2.log
/tmp/testdir/logs/log3.log
/tmp/testdir/text/text1.txt
/tmp/testdir/text/text2.txt
/tmp/testdir/text/text3.txt
```
2. `$go run . -root /tmp/testdir/ -size 0`<br>
list all files with minimum size 0, means all files
``` 
/tmp/testdir/file1.txt
/tmp/testdir/logs/log1.log
/tmp/testdir/logs/log2.log
/tmp/testdir/logs/log3.log
/tmp/testdir/text/text1.txt
/tmp/testdir/text/text2.txt
/tmp/testdir/text/text3.txt
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
/tmp/testdir/logs/log2.log
/tmp/testdir/logs/log3.log

```
6. `$go run . -root /tmp/testdir/ -ext txt` <br>
list all files extension .txt
```  
/tmp/testdir/file1.txt
/tmp/testdir/text/text1.txt
/tmp/testdir/text/text2.txt
/tmp/testdir/text/text3.txt
```
---