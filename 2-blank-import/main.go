package main

import (
	"log/slog"
	"os"

	"github.com/asger-noer/go-extensions/2-blank-import/registry"
	"github.com/lmittmann/tint"

	// Blank imports for the plugins to register
	_ "github.com/asger-noer/go-extensions/2-blank-import/extensions/gopher"
	_ "github.com/asger-noer/go-extensions/2-blank-import/extensions/rustacean"
)

func main() {
	logger := slog.New(tint.NewHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	runners := registry.GetAll()

	for _, runner := range runners {
		runner.Print(logger)
	}
}
