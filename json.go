package lazy

import (
	"bytes"
	"encoding/json"
	"io"
)

func JSONMarshal(v interface{}) Serializable {
	return &jsonmarshal{v: v}
}

func JSONUnmarshal(v interface{}) Deserializable {
	return &jsonunmarshal{v: v}
}

type jsonmarshal struct {
	v interface{}
	b []byte
	s string
}

func (j *jsonmarshal) Close() error {
	j.b = nil
	return nil
}

func (j *jsonmarshal) Bytes() []byte {
	bytes, err := json.Marshal(j.v)

	if err != nil {
		panic(err)
	}

	return bytes
}

func (j *jsonmarshal) String() string {
	if len(j.s) == 0 {
		j.s = string(j.Bytes())
	}
	return j.s
}

func (j *jsonmarshal) Pipe(w io.Writer) error {
	j.Close()
	_, err := io.Copy(w, j)
	return err
}

func (j *jsonmarshal) Read(p []byte) (int, error) {
	if j.b == nil {
		j.b = j.Bytes()
	}

	cnt := copy(p, j.b)

	j.b = j.b[cnt:]

	if len(j.b) == 0 {
		return cnt, io.EOF
	}

	return cnt, nil
}

type jsonunmarshal struct {
	v interface{}
	b bytes.Buffer
}

func (j *jsonunmarshal) FromBytes(b []byte) {
	err := json.Unmarshal(b, &j.v)

	if err != nil {
		panic(err)
	}
}

func (j *jsonunmarshal) FromString(s string) {
	j.FromBytes([]byte(s))
}

func (j *jsonunmarshal) Pipe(r io.Reader) error {
	_, err := io.Copy(j, r)
	return err
}

func (j *jsonunmarshal) Write(p []byte) (n int, err error) {
	return j.b.Write(p)
}

func (j *jsonunmarshal) Close() error {
	err := json.Unmarshal(j.b.Bytes(), &j.v)
	j.b = bytes.Buffer{}
	return err
}
