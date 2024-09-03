package rustacean

import (
	"fmt"
	"log/slog"

	"github.com/asger-noer/go-extensions/3-blank-import-config/registry"
	"github.com/asger-noer/go-extensions/runner"
)

func init() {
	registry.MustRegister("rustaceans", &RustaceanConfig{})
}

type RustaceanConfig struct {
	Overlord string `json:"overlord"`
}

// NewPrinter implements runner.Config.
func (r *RustaceanConfig) NewPrinter() (runner.Printer, error) {
	return &Rustacean{
		overload: r.Overlord,
	}, nil
}

type Rustacean struct {
	overload string
}

// Print implements runner.Plugin.
func (h *Rustacean) Print(log *slog.Logger) {
	msg := fmt.Sprintf("Hello, Rustaceans! ooh, I mean %s!", h.overload)
	log.Info(msg)
}
