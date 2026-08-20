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
	return action(context.Background())
}
