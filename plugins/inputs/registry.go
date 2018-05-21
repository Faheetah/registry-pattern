package inputs

// InputPlugin defines the interface that can interact with the registry
type InputPlugin interface {
	Message() string
	Start() (string, error)
}

// Register can be called from init() on a plugin in this package
// It will automatically be added to the Inputs map to be called externally
func Register(name string, input InputFactory) {
	Inputs[name] = input
}

// InputFactory lets us use a closure to get intsances of the plugin struct
type InputFactory func() InputPlugin

// Inputs registry
var Inputs = map[string]InputFactory{}
