package workflow

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/converter"

	"github.com/odyfey/temporal-decode-oneof-demo/gen/yetanothercloud/compute/v1"
)

func TestProtoJSONPayload_Disk(t *testing.T) {
	disk := &compute.Disk{
		Id:     "trdc7u8s75v5je2gpoth",
		Name:   "disk-test-1",
		Status: compute.Disk_STATUS_READY,
		Source: &compute.Disk_SourceImageId{
			SourceImageId: "iw92psl6l0dyp8ctoqm",
		},
	}

	pc := converter.NewProtoJSONPayloadConverter()
	payload, err := pc.ToPayload(disk)
	require.NoError(t, err)

	disk2 := &compute.Disk{}
	err = pc.FromPayload(payload, disk2)
	require.NoError(t, err)
	assert.Equal(t, disk.Id, disk2.Id)
	assert.Equal(t, disk.Name, disk2.Name)
}
