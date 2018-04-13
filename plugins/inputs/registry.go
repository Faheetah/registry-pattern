package inputs

// InputPlugin defines the interface that can interact with the registry
type InputPlugin interface {
	Message() string
	Start() (string, error)
}

// Register can be called from init() on a plugin in this package
// It will automatically be added to the Inputs map to be called externally
func Register(name string, plugin InputPlugin) {
	Inputs[name] = plugin
}

// The registry
var Inputs = make(map[string]InputPlugin)
