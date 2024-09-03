package runner

import "log/slog"

type Config interface {
	// NewPrinter creates a new Printer.
	NewPrinter() (Printer, error)
}

type Printer interface {
	// Print prints a message.
	Print(*slog.Logger)
}
