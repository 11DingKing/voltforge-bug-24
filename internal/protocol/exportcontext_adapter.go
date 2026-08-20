package protocol

import "context"

func RunExportContext(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	svc := ExportContextService{}
	return svc.Execute(ctx, func(callCtx context.Context) error {
		if err := callCtx.Err(); err != nil {
			return err
		}
		return nil
	})
}
