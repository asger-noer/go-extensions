package rustacean

import (
	"log/slog"
)

type Rustacean struct{}

// Print implements runner.Plugin.
func (h *Rustacean) Print(log *slog.Logger) {
	log.Info("Hello, Rustaceans! ooh, I mean Gophers!")
}
