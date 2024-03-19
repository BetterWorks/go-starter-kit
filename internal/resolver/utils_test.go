package resolver

import (
	"log/slog"
	"testing"

	"github.com/BetterWorks/go-starter-kit/internal/core/logger"
	te "github.com/BetterWorks/go-starter-kit/test/testutils/errors"
)

func TestLogLevel(t *testing.T) {
	// Test case 1: Debug level
	debugLevel := logLevel(logger.LevelDebug)
	if debugLevel != slog.LevelDebug {
		te.NewLineErrorf(t, slog.LevelDebug, debugLevel)
	}

	// Test case 2: Info level
	infoLevel := logLevel(logger.LevelInfo)
	if infoLevel != slog.LevelInfo {
		te.NewLineErrorf(t, slog.LevelInfo, infoLevel)
	}

	// Test case 3: Warn level
	warnLevel := logLevel(logger.LevelWarn)
	if warnLevel != slog.LevelWarn {
		te.NewLineErrorf(t, slog.LevelWarn, warnLevel)
	}

	// Test case 4: Error level
	errorLevel := logLevel(logger.LevelError)
	if errorLevel != slog.LevelError {
		te.NewLineErrorf(t, slog.LevelError, errorLevel)
	}

	// Test case 5: Default level
	defaultLevel := logLevel("invalid")
	if defaultLevel != slog.LevelInfo {
		te.NewLineErrorf(t, slog.LevelInfo, defaultLevel)
	}
}
