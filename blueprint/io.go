package blueprint

import (
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"io"
	"log"
	"os"
)

type GenericForm struct {
	Blueprint *Blueprint `json:"blueprint,omitempty"`
	Book      *Book      `json:"blueprint_book,omitempty"`
}

func (g *GenericForm) WriteJson(output io.Writer) (err error) {
	indent, err := json.MarshalIndent(g, "", "  ")
	if err == nil {
		_, err = output.Write(indent)
	}
	return
}

func (g *GenericForm) Write(output io.Writer) error {

	_, err := output.Write([]byte("0"))
	if err != nil {
		return err
	}
	b := base64.NewEncoder(base64.StdEncoding, output)
	z := zlib.NewWriter(b)
	j := json.NewEncoder(z)
	err = j.Encode(g)
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

func (g *GenericForm) CreateFile(path string) error {
	output, err := os.Create(path)
	if err != nil {
		return err
	}
	return g.Write(output)
}

func (g *GenericForm) AddBlueprint(b *Blueprint) {

	// if we have an existing book, add the blueprint to it
	if g.Book != nil {
		g.Book.AddBlueprint(b)
		return
	}

	// no book, but existing blueprint, move both to book
	if g.Blueprint != nil {
		previous := g.Blueprint
		g.Blueprint = nil
		g.Book = &Book{}
		g.Book.AddBlueprint(previous)
		g.AddBlueprint(b)
		return
	}

	// assign single blueprint
	g.Blueprint = b
}

func Read(io io.ReadCloser) *GenericForm {

	version := []byte{0}

	_, err := io.Read(version)
	if err != nil {
		panic(err)
	}

	if version[0] != 48 {
		log.Fatal("not a blueprint version I understand", version[0])
	}

	m := &GenericForm{}

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
