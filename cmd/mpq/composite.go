package main

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/gravestench/mpq"
)

func getArchiveComposite(o options) (c archiveComposite) {
	for _, path := range o.archivePaths {
		matches, err := filepath.Glob(path)
		if err != nil {
			continue
		}

		for _, match := range matches {
			mpqArchive, err := mpq.New(match)
			if err != nil {
				log.Fatalf("opening mpq archive: %v", err)
			}

			c = append(c, mpqArchive)
		}
	}

	return
}

type archiveComposite []*mpq.MPQ

func (a archiveComposite) Open(internalPath string) (fs.File, error) {
	for _, archive := range a {
		file, err := archive.Open(internalPath)
		if err != nil {
			continue
		}

		return file, nil
	}

	return nil, fmt.Errorf("file not found")
}
