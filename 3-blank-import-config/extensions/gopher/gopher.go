package gopher

import (
	"log/slog"

	"github.com/asger-noer/go-extensions/3-blank-import-config/registry"
	"github.com/asger-noer/go-extensions/runner"
)

func init() {
	registry.MustRegister("gophers", &GopherConfig{})
}

type GopherConfig struct {
	Version string `json:"version"`
}

// NewPrinter implements runner.Config.
func (g *GopherConfig) NewPrinter() (runner.Printer, error) {
	return &Gopher{version: g.Version}, nil
}

type Gopher struct {
	version string
}

// Print implements runner.Plugin.
func (g *Gopher) Print(log *slog.Logger) {
	log.Info("Hello, Gophers!", "version", g.version)
}
