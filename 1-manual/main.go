package main

import (
	"log/slog"
	"os"

	"github.com/asger-noer/go-extensions/1-manual/extensions/gopher"
	"github.com/asger-noer/go-extensions/1-manual/extensions/rustacean"
	"github.com/asger-noer/go-extensions/1-manual/registry"
	"github.com/lmittmann/tint"
)

func main() {
	logger := slog.New(tint.NewHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Register the plugins as a part of the main function.
	registry.MustRegister("gopher", &gopher.Gopher{})
	registry.MustRegister("rustacean", &rustacean.Rustacean{})
	runners := registry.GetAll()

	for _, runner := range runners {
		runner.Print(logger)
	}
}
