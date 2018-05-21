package outputs

// OutputPlugin defines the interface that can interact with the registry
type OutputPlugin interface {
	Send(string) error
}

// Register can be called from init() on a plugin in this package
// It will automatically be added to the Inputs map to be called externally
func Register(name string, output OutputFactory) {
	Outputs[name] = output
}

// OutputFactory lets us use a closure to get instances of the plugin struct
type OutputFactory func() OutputPlugin

// Outputs registry
var Outputs = map[string]OutputFactory{}
