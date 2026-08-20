package protocol

import "context"

func RunExportContext(ctx context.Context) error {
	svc := ExportContextService{}
	return svc.Execute(ctx, func(callCtx context.Context) error {
		return callCtx.Err()
	})
}
