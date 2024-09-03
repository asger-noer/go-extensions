package registry

import "github.com/asger-noer/go-extensions/runner"

var registry = make(map[string]runner.Printer)

func MustRegister(name string, plugin runner.Printer) {
	registry[name] = plugin
}

func Get(name string) runner.Printer {
	return registry[name]
}

func GetAll() map[string]runner.Printer {
	return registry
}
