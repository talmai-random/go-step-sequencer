package main

import (
	"drum"
	"flag"
	"fmt"
	"github.com/talmai-random/portmidi"
	"sequencer"
	"time"
)

// Entry point for go-step-sequencer
// Parses command line flags, which can be either:
//   - --pattern A filepath for the splice pattern on the filesystem.
//   - --kit A directory that contains all the samples for the tracks contained in the pattern.
func main() {
	var patternPath string
	var kitPath string

	flag.StringVar(
		&patternPath,
		"pattern",
		"patterns/pattern_1.splice",
		"-pattern=path/to/pattern.splice",
	)

	flag.StringVar(
		&kitPath,
		"kit",
		"kits",
		"-kit=path/to/kits",
	)

	flag.Parse()

	pattern, err := drum.DecodeFile(patternPath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for _, track := range pattern.Tracks {
		filepath := kitPath + "/" + pattern.Version + "/" + track.Name + ".wav"

		track.Buffer, err = sequencer.LoadSample(filepath)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		track.Playhead = len(track.Buffer)

		fmt.Printf("loaded sample: %s\n", filepath)
	}

	fmt.Printf("%s\n", pattern)

	s, err := sequencer.NewSequencer()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	s.Pattern = pattern
	s.Timer.SetTempo(s.Pattern.Tempo)

	s.Start()

	portmidi.Initialize()

	fmt.Printf("Found %d devices\n", portmidi.CountDevices())
	fmt.Printf("Default Input Device info: %+v\n", portmidi.Info(portmidi.DefaultInputDeviceID()))
	fmt.Printf("Default Output Device info: %+v\n", portmidi.Info(portmidi.DefaultOutputDeviceID()))

	deviceID := portmidi.DefaultInputDeviceID()
	in, err := portmidi.NewInputStream(deviceID, 1024)
	if err != nil {
		fmt.Printf("An error occurred: %+v", err)
	}
	defer in.Close()

	for {
		time.Sleep(time.Second)
		fmt.Printf("Found %d devices\n", portmidi.CountDevices())
		fmt.Printf("Device info: %+v\n", portmidi.Info(portmidi.DefaultInputDeviceID()))
		events, err := in.Read(10)
		if err != nil {
			fmt.Printf("An error occurred: %+v", err)
		}
		for _, evt := range events {
			fmt.Printf("Event read: %+v\n", evt)
		}

		out, err := portmidi.NewOutputStream(portmidi.DefaultOutputDeviceID(), 1024, 0)
		if err != nil {
			fmt.Printf("An error occurred: %+v", err)
		}

		// note on events to play C major chord
		out.WriteShort(0x90, 61, 90)
		out.WriteShort(0x90, 64, 100)
		out.WriteShort(0x90, 67, 110)

		out.Close()
	}

}
