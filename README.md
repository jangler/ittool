ittool
======
A command-line tool for manipulating Impulse Tracker modules.

Installation
------------
Install via the [go command](http://golang.org/cmd/go/):

	go get -u github.com/jangler/ittool

Usage
-----
	Usage: ittool <cmd> [<arg>]...

	A command-line tool for manipulating Impulse Tracker modules.

	If not enough command-line arguments are specified for a command,
	remaining arguments are read from standard input. For help regarding a
	specific command, see 'ittool <cmd> -h'.

	Commands:
	  dump   dump samples from an IT file
	  msg    print the song message from an IT file
	  title  print song titles of IT files
