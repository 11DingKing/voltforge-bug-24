package protocol

import (
	"context"
	"fmt"
)

type ExportContextService struct{}

func (s ExportContextService) Execute(ctx context.Context, action func(context.Context) error) error {
	if ctx == nil {
		return fmt.Errorf("exportcontext: nil context")
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("exportcontext before action: %w", err)
	}
	if err := action(ctx); err != nil {
		return fmt.Errorf("exportcontext action: %w", err)
	}
	if err := ctx.Err(); err != nil {
		return fmt.Errorf("exportcontext after action: %w", err)
	}
	return nil
}
