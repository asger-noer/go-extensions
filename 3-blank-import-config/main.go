package main

import (
	"encoding/json"
	"log/slog"
	"os"

	"github.com/asger-noer/go-extensions/3-blank-import-config/registry"
	"github.com/asger-noer/go-extensions/runner"
	"github.com/lmittmann/tint"

	// Blank imports for the plugins to register
	_ "github.com/asger-noer/go-extensions/3-blank-import-config/extensions/gopher"
	_ "github.com/asger-noer/go-extensions/3-blank-import-config/extensions/rustacean"
)

// Config is the representation of the configuration file.
type Config struct {
	Plugins map[string]json.RawMessage `json:"plugins"`
}

func main() {
	logger := slog.New(tint.NewHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	file, err := os.Open("config.json")
	if err != nil {
		logger.Error("error opening config file", "error", err)
		return
	}

	defer file.Close()

	var config Config

	// Decode the configuration file into the Config struct. This will only
	// partially decode the file, as the extension configuration is stored in the json.RawMessage.
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		logger.Error("decoding file")
		return
	}

	var plugins []runner.Printer
	for name, rawConfig := range config.Plugins {
		// Get the configuration for the plugin. This will be registered by the extensions themselves.
		config, err := registry.Get(name)
		if err != nil {
			logger.Error("creating runner", "error", err)
			continue
		}

		// Unmarshal the raw configuration into the plugin's configuration struct.
		if err := json.Unmarshal(rawConfig, &config); err != nil {
			logger.Error("unmarshaling config", "error", err)
			continue
		}

		// Create a new printer for the from the extentsion configuration.
		printer, err := config.NewPrinter()
		if err != nil {
			logger.Error("creating runner", "error", err)
			continue
		}

		plugins = append(plugins, printer)
	}

	for _, runner := range plugins {
		runner.Print(logger)
	}
}
