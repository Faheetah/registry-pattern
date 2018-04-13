package inputs

import (
	"fmt"
	"log"
)

type Http struct {
	Port    int
	message string
}

func (h *Http) Start() (out string, err error) {
	err = nil
	out = fmt.Sprintf("input http plugin loaded, listening on %d", h.Port)
	log.Println(out)
	h.message = fmt.Sprintf("read some data from port: %d", h.Port)
	return
}

func (h *Http) Message() (string) {
	return h.message
}

func init() {
	log.Println("initializing http")
	Register("http", &Http{})
}
