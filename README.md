[![GoDoc](https://godoc.org/github.com/Anaminus/pkm?status.png)](https://godoc.org/github.com/Anaminus/pkm)

# PKM

PKM is a Go library for extracting data from the ROM files of various Pokemon
games. Currently, only games from generation III are targeted, implemented by
the [gen3](/gen3) sub-package.

## Testing

The gen3 sub-package has almost complete test coverage. Assuming a proper Go
installation, tests can be run by navigating to the `gen3` directory and
running `go test`.

In order to run tests successfully, you must have a file containing a ROM dump
of the English version of Pokemon Emerald (game code `BPEE`). This file must
be placed in the `gen3` directory, and must be named `rom.gba`. In the
interest of remaining legal, this repository will never provide or link to any
ROM files. Go find them yourself, scrub.
