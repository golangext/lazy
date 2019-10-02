package lazy

import "fmt"

type sprintf struct {
	format string
	a      []interface{}
}

func Sprintf(format string, a ...interface{}) fmt.Stringer {
	return &sprintf{format, a}
}

func (f *sprintf) String() string {
	return fmt.Sprintf(f.format, f.a...)
}

type errorf struct {
	format    string
	a         []interface{}
	formatted bool
	res       string
}

func Errorf(format string, a ...interface{}) error {
	return &errorf{format, a, false, ""}
}

func (f *errorf) Error() string {
	if !f.formatted {
		f.res = fmt.Sprintf(f.format, f.a...)
		f.formatted = true
	}
	return f.res
}

type lazybytes struct {
	b []byte
	s string
}

func String(b []byte) fmt.Stringer {
	return &lazybytes{b, ""}
}

func (b *lazybytes) String() string {
	if len(b.b) != len(b.s) {
		b.s = string(b.b)
	}
	return b.s
}
