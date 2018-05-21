package outputs

import (
	"fmt"
	"log"
)

type Http struct {
	Url string
}

func (h *Http) Send(msg string) (err error) {
	err = nil
	out := fmt.Sprintf("output http plugin loaded, sending HTTP request to %s", h.Url)
	log.Println(out, msg)
	return
}

func init() {
	log.Println("initializing http")
	Register("http", func() OutputPlugin { return &Http{} })
}
