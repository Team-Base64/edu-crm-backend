package model

import "io"

type File interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}

type Attach struct {
	Dest string
	File File
}
