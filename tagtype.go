package rpm

//go:generate stringer -type TagType

type TagType int32

const (
	Null TagType = iota
	Char
	Int8
	Int16
	Int32
	Int64
	String
	Bin
	StringArray
	I18nString
)
