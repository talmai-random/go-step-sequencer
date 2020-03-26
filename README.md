# go-step-sequencer

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

## SPLICE file format
  - [Header](#header)
    - ["SPLICE" File Type](#splice_file_body_type)
    - [Splice File Body Size](#splice_file_body_size)
  - [Pattern](#pattern)
    - [Pattern Version String](#pattern_version_string)
    - [Pattern Tempo](#pattern_tempo)
  - [Tracks](#tracks)
    - [Track Id](#track_id)
    - [Track Name](#track_name)
    - [Track Step Sequence](#track_step_sequence)

<a name="header"></a>
### Header

The first section of any `.splice` file is the header sectoin.  It contains a File Type and a File Body Size.

<a name="splice_file_body_type"></a>
#### "SPLICE" File Type

The Splice file type is fairly simple: a 6-byte string with a static value of "SPLICE".  This identifies the file as a `.splice` file.

<a name="splice_file_body_size"></a>
#### Splice File Body Size

The Splice File Body Size is an unsigned 64-bit integer that describes the length of the Body of the `.splice` file in bytes. It should be noted that `pattern_5.splice` is the only example file that has extra, erroneous data encoded in it that occurs after the aforementioned Splice File Body Size. Note: This code discards any data in the file that occurs after the Splice File Body Size.

<a name="pattern"></a>
### Pattern

The First Part of the body encodes the Pattern Version followed by the Pattern Tempo.

<a name="pattern_version_string"></a>
#### Pattern Version String

The Pattern Version String is a 32-bytes in Length.  The unused bytes is discarded in order to display the version string correctly.

<a name="pattern_tempo"></a>
#### Pattern Tempo

The Pattern Tempo is encoded next in the file as a floating 32-bit number.

<a name="tracks"></a>
### Tracks

The next substantial section of the document is all of the Tracks in sequence.

<a name="track_id"></a>
#### Track Id

The Track Id is an unsigned 8-bit integer.

<a name="track_name"></a>
#### Track Name

The Track Name is comprised of two parts encoded in the file. First, is an unsigned 32-bit integer which describes the size of the track name in bytes.  This value is the number of bytes in which to read off of the file in order to receive the name of the track.

<a name="track_step_sequence"></a>
#### Track Step Sequence

The final component of each Track is a series of 16 bytes, each of which describes whether or not the corresponding track should be driggered for a given 16th note in the step sequencer.
## Notes

The synchronization protocol for the step sequencer borrows heavily from the MIDI Beat Clock protocol, which uses 24 Pulses Per Quarter note for a specified Tempo.  More information on this technique is found [here](http://www.blitter.com/~russtopia/MIDI/~jglatt/tech/midispec/seq.htm).

The sampler uses a "Playhead" seeking strategy for playback similar to wavetable sythesizers.

For playback, the sampler mixes all of the tracks encoded in the pattern and outputs them into a single Portaudio Stream.
