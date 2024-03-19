# temporal-decode-oneof-demo

This demo reproduces a potential bug when decoding a protobuf message that contains a oneof field.

## Run demo

### Start temporal dev server

```sh
temporal server start-dev --log-format pretty --log-level warn
```

### Start worker

```sh
make worker
```

### Start workflows

```sh
make starter
```

## Workflows execution result

```sh
2024/02/23 16:09:10 INFO  No logger configured for temporal client. Created default one.
2024/02/23 16:09:10 INFO Decode workflow started WorkflowID=decode_0 RunID=df54fb2e-d6a9-46a5-9cd9-0bf19327fdb4
2024/02/23 16:09:10 ERROR error during decode workflow error="failed to get workflow result: workflow execution error (type: Decode, workflowID: decode_0, runID: df54fb2e-d6a9-46a5-9cd9-0bf19327fdb4): failed to list compute disks: payload item 0: unable to decode: json: cannot unmarshal object into Go struct field Disk.Source of type compute.isDisk_Source (type: wrapError, retryable: true): payload item 0: unable to decode: json: cannot unmarshal object into Go struct field Disk.Source of type compute.isDisk_Source (type: wrapError, retryable: true): unable to decode: json: cannot unmarshal object into Go struct field Disk.Source of type compute.isDisk_Source (type: wrapError, retryable: true): unable to decode"
2024/02/23 16:09:10 INFO DecodeWithOptions workflow started WorkflowID=decode_options_1 RunID=61e203cc-84ed-4bb5-809b-ff36bb5dc78a
2024/02/23 16:09:10 ERROR error during decode with options workflow error="failed to get workflow result: workflow execution error (type: DecodeWithOptions, workflowID: decode_options_1, runID: 61e203cc-84ed-4bb5-809b-ff36bb5dc78a): failed to list compute disks: payload item 0: unable to decode: json: cannot unmarshal object into Go struct field Disk.Source of type compute.isDisk_Source (type: wrapError, retryable: true): payload item 0: unable to decode: json: cannot unmarshal object into Go struct field Disk.Source of type compute.isDisk_Source (type: wrapError, retryable: true): unable to decode: json: cannot unmarshal object into Go struct field Disk.Source of type compute.isDisk_Source (type: wrapError, retryable: true): unable to decode"
```
