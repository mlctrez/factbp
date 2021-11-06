package blueprint

import (
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"os"
)

// Container represents an item ( book or blueprint ) stored in a blueprint string or an item in a blueprint book.
//
// The Index field on the root container should not be set. It will be cleared when any of the write
// methods are invoked.
type Container struct {
	Blueprint *Blueprint `json:"blueprint,omitempty"`
	Book      *Book      `json:"blueprint_book,omitempty"`
	Index     *uint64    `json:"index,omitempty"`
}

// WriteJson writes an indented representation of the container.
func (c *Container) WriteJson(output io.Writer) (err error) {
	indent, err := json.MarshalIndent(c, "", "  ")
	if err == nil {
		_, err = output.Write(indent)
	}
	return
}

// Write compresses the blueprint json data using zlib, encodes it using base64, and prepends the version.
//
// Reference: https://wiki.factorio.com/Blueprint_string_format
func (c *Container) Write(output io.Writer) error {
	if c.Index!=nil {
		c.Index=nil
	}

	_, err := output.Write([]byte("0"))
	if err != nil {
		return err
	}
	b := base64.NewEncoder(base64.StdEncoding, output)
	z := zlib.NewWriter(b)
	j := json.NewEncoder(z)
	err = j.Encode(c)
	if err != nil {
		return err
	}
	err = z.Flush()
	if err != nil {
		return err
	}
	err = z.Close()
	if err != nil {
		return err
	}
	err = b.Close()
	if err != nil {
		return err
	}

	if closer, ok := output.(io.Closer); ok {
		return closer.Close()
	}

	return nil
}

// CreateFile is shorthand for os.Create followed by Write.
func (c *Container) CreateFile(path string) error {
	output, err := os.Create(path)
	if err != nil {
		return err
	}
	return c.Write(output)
}

// Read decodes the versioned, base64 encoded, zlib compressed json data.
func Read(io io.ReadCloser) *Container {

	version := []byte{0}

	_, err := io.Read(version)
	if err != nil {
		panic(err)
	}

	if version[0] != 48 {
		log.Fatal("not a blueprint version I understand", version[0])
	}

	m := &Container{}

	reader, err := zlib.NewReader(base64.NewDecoder(base64.StdEncoding, io))
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(reader).Decode(&m)
	if err != nil {
		panic(err)
	}

	return m

}

// ReadAsMap is a convenience method for extracting the blueprint data from an encoded blueprint string.
func ReadAsMap(io io.ReadCloser) map[string]interface{} {
	version := []byte{0}

	_, err := io.Read(version)
	if err != nil {
		panic(err)
	}

	if version[0] != 48 {
		log.Fatal("not a blueprint version I understand", version[0])
	}

	m := make(map[string]interface{})

	reader, err := zlib.NewReader(base64.NewDecoder(base64.StdEncoding, io))
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(reader).Decode(&m)
	if err != nil {
		panic(err)
	}

	return m
}
