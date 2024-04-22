package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gravestench/mpq"
)

func main() {
	opt := parseCliOptions()
	archives := getArchiveComposite(opt)

	f, err := archives.Open(opt.selector)
	if err != nil {
		log.Fatalf("searching archives for file: %v", err)
	}

	data, err := io.ReadAll(f)
	if err != nil {
		log.Fatalf("reading file data: %v", err)
	}

	formattedInternalPath := strings.ReplaceAll(opt.selector, "\\", string(os.PathSeparator))
	outPath := filepath.Join(opt.destination, formattedInternalPath)

	if err = mkdirAllAndWriteFile(outPath, data); err != nil {
		log.Fatalf("writing file: %v", err)
	}

	fmt.Printf("extracted to %q\n", outPath)
}

func getArchiveComposite(o options) (c composite) {
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

func mkdirAllAndWriteFile(destination string, data []byte) error {
	_ = os.MkdirAll(filepath.Dir(destination), 0o755)

	if err := os.WriteFile(destination, data, 0o755); err != nil {
		log.Fatalf("extracting file: %v\n", err)
	}

	return nil
}
