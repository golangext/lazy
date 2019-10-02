package lazy

import (
	"fmt"
	"io"
)

type Reference interface {
	Elem() interface{}
}

type Byter interface {
	Bytes() []byte
}

type Debyter interface {
	FromBytes(b []byte)
}

type Stringer fmt.Stringer

type Destringer interface {
	FromString(s string)
}

type Serializable interface {
	Stringer
	Byter
	io.Reader
	io.Closer
	Pipe(io.Writer) error
}

type Deserializable interface {
	Destringer
	Debyter
	io.Writer
	io.Closer
	Pipe(io.Reader) error
}
