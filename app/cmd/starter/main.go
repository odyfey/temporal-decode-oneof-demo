package main

import (
	"context"
	"fmt"
	"log/slog"

	T "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/converter"

	"github.com/odyfey/temporal-decode-oneof-demo/app/workflow"
)

func main() {
	tc, err := T.Dial(T.Options{
		DataConverter: converter.NewCompositeDataConverter(
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
		),
	})
	if err != nil {
		slog.Error("error during dial to temporal server", "error", err)
		return
	}
	defer tc.Close()

	queue := "demo-protobuf-oneof"

	if err := runDecodeWorkflow(tc, queue); err != nil {
		slog.Error("decode workflow failed", "error", err)
	}

	if err := runDecodeWithOptionsWorkflow(tc, queue); err != nil {
		slog.Error("decode with options workflow failed", "error", err)
	}
}

func runDecodeWorkflow(tc T.Client, queue string) error {
	opts := T.StartWorkflowOptions{
		ID:        "decode_0",
		TaskQueue: queue,
	}

	wr, err := tc.ExecuteWorkflow(context.Background(), opts, workflow.Decode)
	if err != nil {
		return fmt.Errorf("failed to execute workflow: %w", err)
	}

	slog.Info("Decode workflow started", "WorkflowID", wr.GetID(), "RunID", wr.GetRunID())

	if err = wr.Get(context.Background(), nil); err != nil {
		return fmt.Errorf("failed to get workflow result: %w", err)
	}

	return nil
}

func runDecodeWithOptionsWorkflow(tc T.Client, queue string) error {
	opts := T.StartWorkflowOptions{
		ID:        "decode_options_1",
		TaskQueue: queue,
	}

	wr, err := tc.ExecuteWorkflow(context.Background(), opts, workflow.DecodeWithOptions)
	if err != nil {
		return fmt.Errorf("failed to execute workflow: %w", err)
	}

	slog.Info("DecodeWithOptions workflow started", "WorkflowID", wr.GetID(), "RunID", wr.GetRunID())

	if err = wr.Get(context.Background(), nil); err != nil {
		return fmt.Errorf("failed to get workflow result: %w", err)
	}

	return nil
}
