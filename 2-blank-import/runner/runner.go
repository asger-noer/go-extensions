package runner

import "log/slog"

type Printer interface {
	// Print prints a message.
	Print(*slog.Logger)
}
