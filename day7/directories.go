package day7

import "github.com/pkg/errors"

var (
	errDirExists    = errors.New("directory already exists")
	errDirNotExists = errors.New("directory does not exist")
	errOnRoot       = errors.New("already at root, cannot go up one dir")
	errFileExists   = errors.New("file already exists")
)

type directory struct {
	name        string
	parent      *directory
	directories map[string]*directory
	files       map[string]int
}

func newDirectory(name string, parent *directory) *directory {
	return &directory{
		name:        name,
		parent:      parent,
		directories: make(map[string]*directory),
		files:       make(map[string]int),
	}
}

func (d *directory) addDirectory(name string) error {
	if _, ok := d.directories[name]; ok {
		return errDirExists
	}

	d.directories[name] = newDirectory(name, d)
	return nil
}

func (d *directory) cdIntoDirectory(name string) (*directory, error) {
	cdDir, ok := d.directories[name]
	if !ok { // requested directory not found
		return nil, errDirNotExists
	}

	return cdDir, nil
}

func (d *directory) upDir() (*directory, error) {
	if d.parent == nil {
		return nil, errOnRoot
	}

	return d.parent, nil
}

func (d *directory) rootDir() *directory {
	pd := d.parent
	for {
		if pd == nil {
			return d
		}

		pd = d.parent
	}
}

func (d *directory) addFile(name string, size int) error {
	if _, ok := d.files[name]; ok {
		return errFileExists
	}

	d.files[name] = size
	return nil
}

func (d *directory) size() int {
	s := 0
	for _, dir := range d.directories {
		s += dir.size()
	}

	for _, fileSize := range d.files {
		s += fileSize
	}

	return s
}
