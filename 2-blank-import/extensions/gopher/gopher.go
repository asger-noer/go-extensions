package gopher

import (
	"log/slog"

	"github.com/asger-noer/go-extensions/2-blank-import/registry"
)

func init() {
	registry.MustRegister("gopher", &Gopher{})
}

type Gopher struct{}

// Print implements runner.Plugin.
func (g *Gopher) Print(log *slog.Logger) {
	log.Info("Hello, Gophers!")
}
