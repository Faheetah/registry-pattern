package outputs

import (
	"errors"
	_ "github.com/burntsushi/toml"
)

// OutputPlugin defines the interface that can interact with the registry
type OutputPlugin interface {
	Send(string) (error)
}

// Register can be called from init() on a plugin in this package
// It will automatically be added to the Inputs map to be called externally
func Register(name string, plugin OutputPlugin) {
	Outputs[name] = plugin
}

var Outputs = make(map[string]OutputPlugin)
