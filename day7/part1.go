package day7

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/javorszky/adventofcode2022/inputs"
	"github.com/rs/zerolog"
)

const atMostSize = 100000

func Task1(l zerolog.Logger) {
	localLogger := l.With().Int("day", 7).Int("part", 1).Logger()

	gog, err := inputs.ReadIntoLines("day7/input1.txt")
	if err != nil {
		localLogger.Err(err).Msg("could not read input file")
		os.Exit(1)
	}

	root := buildDirectory(gog, localLogger)
	//100,000
	dirsAtMost100k := filterDirsAtMost(root, atMostSize)
	sumSize := 0

	for _, d100k := range dirsAtMost100k {
		sumSize += d100k.size()
	}

	solution := sumSize
	s := localLogger.With().Int("solution", solution).Logger()
	s.Info().Msgf("The total space for all directories double counted that are at most 100k is %d", solution)
}

func buildDirectory(commands []string, log zerolog.Logger) *directory {
	localLogger := log.With().Str("func", "buildDirectory").Logger()
	root := newDirectory("/", nil)
	currentDir := root

	for i, l := range commands {
		switch {
		case strings.HasPrefix(l, "$ cd"):
			dirToMoveTo := strings.TrimPrefix(l, "$ cd ")
			switch dirToMoveTo {
			case "/":
				currentDir = root
			case "..":
				p, err := currentDir.upDir()
				if err != nil {
					localLogger.Err(err).Msgf("line %d: moving up a directory from %s, %v", i, currentDir.name, currentDir.parent)
					os.Exit(1)
				}
				currentDir = p
			default:
				d, err := currentDir.cdIntoDirectory(dirToMoveTo)
				if err != nil {
					localLogger.Err(err).Msgf("line %d: moving into directory %d failed from current dir %s", i, dirToMoveTo, currentDir.name)
				}
				currentDir = d
			}
		// this is going to be a move command
		case l == "$ ls":
			//localLogger.Info().Msgf("line %d: listing directory contents: %s", l)
			// this is going to be a list command
		case strings.HasPrefix(l, "dir "):
			dirName := strings.TrimPrefix(l, "dir ")
			err := currentDir.addDirectory(dirName)
			if err != nil {
				localLogger.Err(err).Msgf("line %d: adding directory %s to current directory %s. %#v", i, dirName, currentDir.name, currentDir.directories)
				os.Exit(1)
			}
			//localLogger.Info().Msgf("line %d: this is a directory: %s", l)
		// this is a directory listing
		default:
			parts := strings.Split(l, " ")
			if len(parts) != 2 {
				localLogger.Err(errors.New("splitting file declaration did not result in two parts")).Msgf("line %d: file line: %s", i, l)
				os.Exit(1)
			}

			size, err := strconv.Atoi(parts[0])
			if err != nil {
				localLogger.Err(err).Msgf("line %d: parsing file size for line %s: %s", i, l, parts[0])
				os.Exit(1)
			}

			err = currentDir.addFile(parts[1], size)
			if err != nil {
				localLogger.Err(err).Msgf("line %d: adding file %s to currect directory %s", i, l, currentDir.name)
				os.Exit(1)
			}
		}
	}

	return root
}

func filterDirsAtMost(root *directory, atMost int) []*directory {
	out := make([]*directory, 0)

	var getSize func(d *directory)

	getSize = func(d *directory) {
		for _, dir := range d.directories {
			dirSize := dir.size()
			if dirSize <= atMost {
				out = append(out, dir)
			}

			getSize(dir)
		}
	}

	getSize(root)

	return out
}

func filterDirsAtLeast(root *directory, atLeast int) []*directory {
	out := make([]*directory, 0)

	var getSize func(d *directory)

	getSize = func(d *directory) {
		for _, dir := range d.directories {
			dirSize := dir.size()
			if dirSize >= atLeast {
				out = append(out, dir)
			}

			getSize(dir)
		}
	}

	getSize(root)

	return out
}
