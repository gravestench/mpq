package main

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/gravestench/mpq"
)

type composite []*mpq.MPQ

func (a composite) Open(internalPath string) (fs.File, error) {
	for _, archive := range a {
		file, err := archive.Open(internalPath)
		if err != nil {
			continue
		}

		fmt.Printf("found file in %q\n", filepath.Base(archive.Path()))

		return file, nil
	}

	return nil, fmt.Errorf("file not found")
}
