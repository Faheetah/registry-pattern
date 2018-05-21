package main

import (
	"log"

	"github.com/burntsushi/toml"
	inputs "github.com/faheetah/registry-pattern/plugins/inputs"
	outputs "github.com/faheetah/registry-pattern/plugins/outputs"
)

// Config base structure that holds our configuration, but not any actual plugins
type Config struct {
	Inputs  map[string]toml.Primitive
	Outputs map[string]toml.Primitive
}

// App holds the actual app that is built from configuration
type App struct {
	Inputs  map[string]inputs.InputPlugin
	Outputs map[string]outputs.OutputPlugin
}

func main() {
	var md toml.MetaData
	var err error
	var conf Config

	// Create an app struct to hold the live inputs/outputs, explained below
	app := App{
		Inputs:  make(map[string]inputs.InputPlugin),
		Outputs: make(map[string]outputs.OutputPlugin),
	}

	// Initial read of the config, this tells us what plugins to load too
	if md, err = toml.DecodeFile("registry-pattern.toml", &conf); err != nil {
		log.Fatal(err)
	}

	// Build the inputList
	for n, pv := range conf.Inputs {
		var i inputs.InputPlugin

		// Get the input object from the registry
		// Always validate the plugin exists to avoid a NPE
		if p, ok := inputs.Inputs[n]; ok {
			i = p()
		} else {
			log.Fatal(err)
		}

		// The way PrimitiveDecode works is to initialize the struct
		// It does not actually update the conf struct's maps, so we can't call conf.Inputs["file"].Start()
		if err = md.PrimitiveDecode(pv, i); err != nil {
			log.Fatal("could not find input plugin %s", n)
		}

		// For simplicity, we will just store these in an App struct
		app.Inputs[n] = i
	}

	// Do the same for outputs
	for n, pv := range conf.Outputs {
		var o outputs.OutputPlugin

		if p, ok := outputs.Outputs[n]; ok {
			o = p()
		} else {
			log.Fatal("could not find output plugin %s", n)
		}

		if err = md.PrimitiveDecode(pv, o); err != nil {
			log.Fatal(err)
		}

		app.Outputs[n] = o
	}

	// Do something with the inputs
	for _, i := range app.Inputs {
		i.Start()
		log.Println("message:", i.Message())
	}

	// Do something with the outputs, i.e. output the inputs
	for _, o := range app.Outputs {
		o.Send("Sent some output here")
	}
}
