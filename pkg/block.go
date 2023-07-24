package pkg

import (
	"io"
	"strings"
)

// Block represents an entry in the block table
type Block struct { // 16 bytes
	FilePosition         uint32
	CompressedFileSize   uint32
	UncompressedFileSize uint32
	Flags                FileFlag
	// Local Stuff...
	FileName       string
	EncryptionSeed uint32
}

// HasFlag returns true if the specified flag is present
func (b *Block) HasFlag(flag FileFlag) bool {
	return (b.Flags & flag) != 0
}

func (b *Block) calculateEncryptionSeed(fileName string) {
	fileName = fileName[strings.LastIndex(fileName, `\`)+1:]
	seed := hashString(fileName, 3)
	b.EncryptionSeed = (seed + b.FilePosition) ^ b.UncompressedFileSize
}

//nolint:gomnd // number
func (mpq *MPQ) readBlockTable() error {
	if _, err := mpq.file.Seek(int64(mpq.header.BlockTableOffset), io.SeekStart); err != nil {
		return err
	}

	blockData, err := decryptTable(mpq.file, mpq.header.BlockTableEntries, "(block table)")
	if err != nil {
		return err
	}

	for n, i := uint32(0), uint32(0); i < mpq.header.BlockTableEntries; n, i = n+4, i+1 {
		mpq.blocks = append(mpq.blocks, &Block{
			FilePosition:         blockData[n],
			CompressedFileSize:   blockData[n+1],
			UncompressedFileSize: blockData[n+2],
			Flags:                FileFlag(blockData[n+3]),
		})
	}

	return nil
}
