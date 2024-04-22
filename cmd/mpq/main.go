package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// parse cli options and use the values
	opt := parseCliOptions()
	archives := getArchiveComposite(opt)
	destination := reformatInternalPath(opt)

	// try to open the file from one of the archives
	f, err := archives.Open(opt.selector)
	if err != nil {
		log.Fatalf("searching archives for file: %v", err)
	}

	// extract and write the file data to destination path
	if err = mkdirAllAndWriteFileData(destination, f); err != nil {
		log.Fatalf("writing file: %v", err)
	}
}

func reformatInternalPath(opt options) (formatted string) {
	formattedInternalPath := strings.ReplaceAll(opt.selector, "\\", string(os.PathSeparator))
	return filepath.Join(opt.destination, formattedInternalPath)
}

func mkdirAllAndWriteFileData(destination string, f io.Reader) error {
	data, err := io.ReadAll(f)
	if err != nil {
		return fmt.Errorf("reading file data: %v", err)
	}

	_ = os.MkdirAll(filepath.Dir(destination), 0o755)

	if err := os.WriteFile(destination, data, 0o755); err != nil {
		return fmt.Errorf("extracting file: %v\n", err)
	}

	return nil
}
