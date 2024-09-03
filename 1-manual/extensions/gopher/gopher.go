package gopher

import (
	"log/slog"
)

type Gopher struct{}

// Print implements runner.Plugin.
func (g *Gopher) Print(log *slog.Logger) {
	log.Info("Hello, Gophers!")
}
