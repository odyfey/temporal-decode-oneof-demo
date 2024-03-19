package main

import (
	"log/slog"

	T "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/converter"
	"go.temporal.io/sdk/worker"

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

	wrk := worker.New(tc, "demo-protobuf-oneof", worker.Options{})
	wrk.RegisterActivity(workflow.GetDisk)
	wrk.RegisterActivity(workflow.ListDisks)
	wrk.RegisterWorkflow(workflow.DecodeDisk)
	wrk.RegisterWorkflow(workflow.DecodeListDisks)
	wrk.RegisterWorkflow(workflow.DecodeWithOptions)

	if err := wrk.Run(worker.InterruptCh()); err != nil {
		slog.Error("error during worker run", "error", err)
	}
}
