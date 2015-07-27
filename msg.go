package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jangler/impulse"
	"github.com/jangler/minipkg/tool"
)

func msg(args []string) {
	file, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}

	module, err := impulse.ReadModule(file)
	if err != nil {
		log.Fatal(err)
	}

	if module.Message != "" {
		text := strings.Replace(module.Message, "\r", "\n", -1)
		fmt.Println(text[:len(text)-1]) // strip NUL byte from end
	}
}

func init() {
	cmd := &tool.Command{
		Name:    "msg",
		Summary: "print the song message from an IT file",
		Usage:   "<file>",
		Description: `
Print the song message from the given IT file.
`,
		Function: msg,
		MinArgs:  1,
		MaxArgs:  1,
	}

	cmd.FlagSet = flag.NewFlagSet(cmd.Name, flag.ExitOnError)
	cmd.FlagSet.Usage = tool.UsageFunc(cmd)

	tool.Commands[cmd.Name] = cmd
}
