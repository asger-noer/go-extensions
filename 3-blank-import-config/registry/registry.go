package registry

import (
	"fmt"

	"github.com/asger-noer/go-extensions/runner"
)

// registry is the map of registered exention configurations. We hold the plugin
// name as the key, so we can easily look up the configuration for a plugin.
var registry = make(map[string]runner.Config)

func MustRegister(name string, plugin runner.Config) {
	registry[name] = plugin
}

func Get(name string) (runner.Config, error) {
	if cfg, ok := registry[name]; ok {
		return cfg, nil
	}

	return nil, fmt.Errorf("plugin %s not found", name)
}
