# go-step-sequencer

[![Join the chat at https://gitter.im/kellydunn/go-step-sequencer](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/kellydunn/go-step-sequencer?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

A step sequencer implemented in Golang using portaudio and libsndfile wrappers.

## Dependencies

This project requires three dependent libraries, which are thinly wrapped by a golang libraries.  In order for the project to successfully compile, you will need to install the development libraries of portaudio, libmpg123 and libsndfile for your platform. To learn more about how to install the native libraries of these dependencies, please visit their official sites:

  - [portaudio](http://www.portaudio.com/)
  - [libsndfile](http://mega-nerd.com/libsndfile/)
  - [libmpg123](http://www.mpg123.org/index.shtml)

## Build

This project is built using a [Makefile](./Makefile). Targets include 'build', 'test' and 'run'.

```bash
$ make
```

Upon a successful build, the binary will exist in the root level directory of this project.  You should be able to invoke it locally with the following command:

```bash
$ ./go-step-sequencer
```

## Usage

Use the `--help` flag to get more information on how to use the utility.

```bash
$ go-step-sequencer --help
Usage of go-step-sequencer:
  -kit="kits": -kit=path/to/kits
  -pattern="patterns/pattern_1.splice": -pattern=path/to/pattern.splice
```

The step sequencer was made to take a `pattern` and a `kit` as command line flags so that you can swap out different types of kits and patterns.  A typical use of the command looks like this:

```bash
$ go-step-sequencer --pattern path/to/pattern.splice --kit path/to/kits
```

The default pattern is found at `patterns/pattern_1.splice` and the default kit is located at `kits/0.808-alpha`.  Running `go-step-sequnencer` without specifying a `--pattern` or a `--kit` will run the default pattern with the default kit:   

```bash
$ go-step-sequencer
loaded sample: kits/0.808-alpha/kick.wav
loaded sample: kits/0.808-alpha/snare.wav
loaded sample: kits/0.808-alpha/clap.wav
loaded sample: kits/0.808-alpha/hh-open.wav
loaded sample: kits/0.808-alpha/hh-close.wav
loaded sample: kits/0.808-alpha/cowbell.wav
Saved with HW Version: 0.808-alpha
Tempo: 120
(0) kick        |x---|x---|x---|x---|
(1) snare       |----|x---|----|x---|
(2) clap        |----|x-x-|----|----|
(3) hh-open     |--x-|--x-|x-x-|--x-|
(4) hh-close    |x---|x---|----|x--x|
(5) cowbell     |----|----|--x-|----|
```

You should be able to hear the drum track out of your speakers now!

## Notes

The synchronization protocol for the step sequencer borrows heavily from the MIDI Beat Clock protocol, which uses 24 Pulses Per Quarter note for a specified Tempo.  More information on this technique is found [here](http://www.blitter.com/~russtopia/MIDI/~jglatt/tech/midispec/seq.htm).

The sampler uses a "Playhead" seeking strategy for playback similar to wavetable sythesizers.

For playback, the sampler mixes all of the tracks encoded in the pattern and outputs them into a single Portaudio Stream.
