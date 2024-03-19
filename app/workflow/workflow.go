package workflow

import (
	"fmt"
	"time"

	"github.com/odyfey/temporal-decode-oneof-demo/gen/yetanothercloud/compute/v1"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/workflow"
)

func Decode(ctx workflow.Context) error {
	lgr := workflow.GetLogger(ctx)
	lgr.Info("workflow with default decode options started")

	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: 5 * time.Second,
	})

	var disks []*compute.Disk

	if err := workflow.ExecuteActivity(ctx, ListDisks).Get(ctx, &disks); err != nil {
		return fmt.Errorf("failed to list compute disks: %w", err)
	}

	for _, d := range disks {
		lgr.Info("disks", "diskId", d.Id, "diskSource", d.Source)
	}

	return nil
}

func DecodeWithOptions(ctx workflow.Context) error {
	lgr := workflow.GetLogger(ctx)
	lgr.Info("workflow with custom decode options started")

	ctx = workflow.WithActivityOptions(workflow.WithDataConverter(ctx,
		converter.NewCompositeDataConverter(
			converter.NewNilPayloadConverter(),
			converter.NewByteSlicePayloadConverter(),

			// Order is important here. Both ProtoJsonPayload and ProtoPayload converters check for the same proto.Message
			// interface. The first match (ProtoJsonPayload in this case) will always be used for serialization.
			// Deserialization is controlled by metadata, therefore both converters can deserialize corresponding data format
			// (JSON or binary proto).
			converter.NewProtoJSONPayloadConverterWithOptions(converter.ProtoJSONPayloadConverterOptions{
				AllowUnknownFields: true,
			}),
			converter.NewProtoPayloadConverter(),

			converter.NewJSONPayloadConverter(),
		)),
		workflow.ActivityOptions{
			StartToCloseTimeout: 5 * time.Second,
		},
	)

	var disks []*compute.Disk

	if err := workflow.ExecuteActivity(ctx, ListDisks).Get(ctx, &disks); err != nil {
		return fmt.Errorf("failed to list compute disks: %w", err)
	}

	for _, d := range disks {
		lgr.Info("disks", "diskId", d.Id, "diskSource", d.Source)
	}

	return nil
}
