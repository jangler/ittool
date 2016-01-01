package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jangler/impulse"
	"github.com/jangler/minipkg/tool"
)

func title(args []string) {
	for _, arg := range args {
		file, err := os.Open(arg)
		if err != nil {
			log.SetPrefix(log.Prefix() + arg + ": ")
			log.Fatal(err)
		}

		module, err := impulse.ReadModule(file)
		if err != nil {
			log.SetPrefix(log.Prefix() + arg + ": ")
			log.Fatal(err)
		}

		fmt.Println(module.SongName)
	}
}

func init() {
	cmd := &tool.Command{
		Name:    "title",
		Summary: "print song titles of IT files",
		Usage:   "<file>...",
		Description: `
Print the song titles from the given IT files.
`,
		Function: title,
		MinArgs:  1,
		MaxArgs:  -1,
	}

	cmd.FlagSet = flag.NewFlagSet(cmd.Name, flag.ExitOnError)
	cmd.FlagSet.Usage = tool.UsageFunc(cmd)

	tool.Commands[cmd.Name] = cmd
}
