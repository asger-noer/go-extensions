package rustacean

import (
	"log/slog"

	"github.com/asger-noer/go-extensions/2-blank-import/registry"
)

func init() {
	registry.MustRegister("rustacean", &Rustacean{})
}

type Rustacean struct{}

// Print implements runner.Plugin.
func (h *Rustacean) Print(log *slog.Logger) {
	log.Info("Hello, Rustaceans! ooh, I mean Gophers!")
}
