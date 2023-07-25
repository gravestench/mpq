package mpq

import (
	"github.com/gravestench/mpq/pkg"
)

// these aliases are here so you can import from the repo root

type (
	MPQ        = pkg.MPQ
	Header     = pkg.Header
	Block      = pkg.Block
	FileRecord = pkg.FileRecord
	FileFlag   = pkg.FileFlag
	Hash       = pkg.Hash
	Stream     = pkg.Stream
	DataStream = pkg.DataStream
	PatchInfo  = pkg.PatchInfo
)

func New(filepath string) (*MPQ, error) {
	return pkg.New(filepath)
}

func FromFile(filepath string) (*MPQ, error) {
	return pkg.FromFile(filepath)
}

func CreateStream(mpq *MPQ, block *Block, fileName string) (*Stream, error) {
	return pkg.CreateStream(mpq, block, fileName)
}
