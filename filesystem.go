package mpq

import (
	"fmt"
	"io/fs"
	"time"
)

// we implement interfaces that allow the mpq archive to act like a filesystem

// MPQ must implement fs.FS
var _ fs.FS = &MPQ{}

func (mpq *MPQ) Open(name string) (fs.File, error) {
	stream, err := mpq.readFileStream(name)
	if err != nil {
		return nil, fmt.Errorf("getting file stream: %v", err)
	}

	return stream, nil
}

// MpqDataStream must implement fs.File
var _ fs.File = &MpqDataStream{}

func (m *MpqDataStream) Stat() (fs.FileInfo, error) {
	return &fileInfo{
		mpq:  m.stream.MPQ,
		path: "",
	}, nil
}

var _ fs.FileInfo = &fileInfo{}

type fileInfo struct {
	mpq    *MPQ
	stream *MpqDataStream
	path   string
}

func (f *fileInfo) Name() string {
	return f.path
}

func (f *fileInfo) Size() int64 {
	return int64(f.stream.stream.Size)
}

func (f *fileInfo) Mode() fs.FileMode {
	return 0o755
}

func (f *fileInfo) ModTime() time.Time {
	return time.Time{}
}

func (f *fileInfo) IsDir() bool {
	return false
}

func (f *fileInfo) Sys() any {
	return f.mpq
}
