package repomd

import (
	"compress/bzip2"
	"compress/gzip"
	"io"
	"os"
	"strings"

	"github.com/xi2/xz"
)

type file struct {
	io.Reader
	closer *os.File
}

type zOpener func(io.Reader) (io.Reader, error)

func bz2Open(r io.Reader) (io.Reader, error) { return bzip2.NewReader(r), nil }
func gzOpen(r io.Reader) (io.Reader, error)  { return gzip.NewReader(r) }
func xzOpen(r io.Reader) (io.Reader, error)  { return xz.NewReader(r, 0) }

// Open opens the named file for reading with suffix-specified decompression.
func Open(name string) (io.ReadCloser, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	for _, t := range []struct {
		x  string
		fn zOpener
	}{
		{".bz2", bz2Open},
		{".gz", gzOpen},
		{".xz", xzOpen},
	} {
		if strings.HasSuffix(name, t.x) {
			if r, err := t.fn(f); err != nil {
				f.Close()
				return nil, err
			} else {
				return &file{r, f}, nil
			}
		}
	}
	return &file{f, f}, nil
}

// Close closes the file, rendering it unusable for I/O.
func (f *file) Close() error {
	if f.closer != nil {
		return f.closer.Close()
	}
	return nil
}
