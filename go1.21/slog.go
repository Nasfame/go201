// https://go.dev/play/p/cZFrHP2rEMd

package main

import (
    "log/slog"
)

func main() {
    slog.Debug("Debug message")
    slog.Info("Info message")
    slog.Warn("Warning message")
    slog.Error("Error message")
}

// https://betterstack.com/community/guides/logging/logging-in-go/

// not as performant as zap but very mem efficent, very less allocs
// performance : zerolog>zap> slog