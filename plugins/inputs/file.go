package inputs

import (
	"fmt"
	"log"
)

type File struct {
	Path    string
	message string
}

func (f *File) Start() (out string, err error) {
	err = nil
	out = fmt.Sprintf("input file plugin loaded, looking at %s", f.Path)
	log.Println(out)
	f.message = fmt.Sprintf("%s contains foobars!", f.Path)
	return
}

func (f *File) Message() (string) {
	return f.message
}

func init() {
	log.Println("initializing file")
	Register("file", &File{})
}
