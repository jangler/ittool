package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jangler/impulse"
	"github.com/jangler/minipkg/tool"
	"github.com/jangler/minipkg/wave"
)

var itsDumpFlag bool

func dumpSample(source string, index int, sample *impulse.Sample) error {
	ext := "wav"
	if itsDumpFlag {
		ext = "its"
	}
	filename := fmt.Sprintf("%s-%03d.%s", source[:len(source)-3], index, ext)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	if itsDumpFlag {
		return sample.Write(file)
	} else {
		waveFile := wave.File{
			Channels:       1,
			SampleRate:     int(sample.Speed),
			BytesPerSample: 1,
			Data:           sample.Data,
		}
		if sample.Flags&impulse.StereoSample != 0 {
			waveFile.Channels = 2
		}
		if sample.Flags&impulse.Quality16Bit != 0 {
			waveFile.BytesPerSample = 2
		}
		return waveFile.Write(file)
	}
}

func dumpFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	module, err := impulse.ReadModule(file)
	if err != nil {
		return err
	}

	for i, sample := range module.Samples {
		if sample.Length == 0 {
			continue
		}
		if err := dumpSample(filename, i+1, sample); err != nil {
			log.Printf("%s: %v\n", filename, err)
		}
	}
	return nil
}

func dump(args []string) {
	for _, arg := range args {
		if err := dumpFile(arg); err != nil {
			log.Printf("%s: %v\n", arg, err)
		}
	}
}

func init() {
	cmd := &tool.Command{
		Name:    "dump",
		Summary: "dump samples from an IT file",
		Usage:   "[<option>] <file>...",
		Description: `
Dump all samples from the given IT files in WAV format.
`,
		Function: dump,
		MinArgs:  1,
		MaxArgs:  -1,
		HasOpts:  true,
	}

	cmd.FlagSet = flag.NewFlagSet(cmd.Name, flag.ExitOnError)
	cmd.FlagSet.Usage = tool.UsageFunc(cmd)
	cmd.FlagSet.BoolVar(&itsDumpFlag, "its", itsDumpFlag,
		"dump ITS files instead of WAV files")

	tool.Commands[cmd.Name] = cmd
}
