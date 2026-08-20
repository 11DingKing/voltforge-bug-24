package protocol

import (
	"context"
	"errors"
	"testing"
)

func TestVoltForge24(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	err := RunExportContext(ctx)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected cancellation, got %v", err)
	}
}
