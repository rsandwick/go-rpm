package rpm // import "rs3.io/go/rpm"

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
)

const Magic = "\xed\xab\xee\xdb"
const HeaderMagic = "\x8e\xad\xe8"

var (
	ErrBadMagic        = errors.New("bad RPM magic number")
	ErrBadHeaderMagic  = errors.New("bad RPM header magic number")
	ErrInvalidTagValue = errors.New("invalid tag value")
	ErrShortArrayTag   = errors.New("short array-type tag")
	ErrTagNotFound     = errors.New("tag not found")
)

type Head struct {
	lead
	Sig  *Header
	Hdr  *Header
	Size int64
}

type lead struct {
	Magic  [4]byte
	Major  byte
	Minor  byte
	Type   int16
	Arch   int16
	Name   [66]byte
	OS     int16
	SigVer int16

	Reserved [16]byte
}

type Header struct {
	header
	Index  []indexEntry
	tagmap map[Tag]int
	data   []byte
}

type header struct {
	Magic    [3]byte
	Version  byte
	Reserved [4]byte
	NIndex   int32 //uint32
	Size     int32 //uint32
}

type indexEntry struct {
	Tag    Tag
	Type   TagType
	Offset int32 //uint32
	Count  int32 //uint32
}

func readHeader(r io.ReadSeeker) (*Header, error) {
	var h header
	if err := binary.Read(r, binary.BigEndian, &h); err != nil {
		return nil, err
	}
	if string(h.Magic[:]) != HeaderMagic {
		return nil, ErrBadHeaderMagic
	}
	index := make([]indexEntry, h.NIndex)
	tagmap := make(map[Tag]int)
	for i := 0; i < int(h.NIndex); i++ {
		binary.Read(r, binary.BigEndian, &index[i])
		tagmap[index[i].Tag] = i
	}
	sz := h.Size
	if sz%8 != 0 {
		sz += 8 - (sz % 8)
	}
	data := make([]byte, sz)
	if _, err := io.ReadFull(r, data); err != nil {
		return nil, err
	}

	return &Header{h, index, tagmap, data}, nil
}

func (h *Header) Get(t Tag) (interface{}, error) {
	i, ok := h.tagmap[t]
	if !ok {
		return nil, ErrTagNotFound
	}
	v := h.Index[i]
	switch v.Type {

	case Bin:
		return h.data[v.Offset : v.Offset+v.Count], nil

	case Int8:
		x := make([]int8, v.Count)
		r := bytes.NewReader(h.data[v.Offset:])
		err := binary.Read(r, binary.BigEndian, x)
		return x, err

	case Int16:
		x := make([]int16, v.Count)
		r := bytes.NewReader(h.data[v.Offset:])
		err := binary.Read(r, binary.BigEndian, x)
		return x, err

	case Int32:
		x := make([]int32, v.Count)
		r := bytes.NewReader(h.data[v.Offset:])
		err := binary.Read(r, binary.BigEndian, x)
		return x, err

	case Int64:
		x := make([]int64, v.Count)
		r := bytes.NewReader(h.data[v.Offset:])
		err := binary.Read(r, binary.BigEndian, x)
		return x, err

	case I18nString, String:
		buf := h.data[v.Offset:]
		i := bytes.IndexByte(buf, 0)
		if i < 0 {
			return nil, ErrInvalidTagValue
		}
		return string(buf[:i]), nil

	case StringArray:
		n := int(v.Count)
		x := make([]string, n)
		y := bytes.SplitN(h.data[v.Offset:], []byte{0}, n+1)
		if len(y) < n {
			return nil, ErrShortArrayTag
		}
		for i := 0; i < n; i++ {
			x[i] = string(y[i])
		}
		return x, nil

	default:
		return nil, nil
	}
}

func (h *Header) ActualSize() int {
	return 16 + (16 * len(h.Index)) + len(h.data)
}

func ReadHead(r io.ReadSeeker) (*Head, error) {
	var l lead
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		return nil, err
	}
	if string(l.Magic[:]) != Magic {
		return nil, ErrBadMagic
	}

	sig, err := readHeader(r)
	if err != nil {
		return nil, err
	}

	hdr, err := readHeader(r)
	if err != nil {
		return nil, err
	}

	offs, err := r.Seek(0, io.SeekCurrent)
	return &Head{l, sig, hdr, offs}, err
}
