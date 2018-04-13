package outputs

import (
	"fmt"
	"log"
)

type File struct {
	Path    string
}

func (f *File) Send(msg string) (err error) {
	err = nil
	out := fmt.Sprintf("output file plugin loaded, \"reading\" to %s", f.Path)
	log.Println(out, msg)
	return
}

func init() {
	log.Println("initializing file")
	Register("file", &File{})
}
