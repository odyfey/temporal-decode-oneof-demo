package workflow

import (
	"context"

	"github.com/odyfey/temporal-decode-oneof-demo/gen/yetanothercloud/compute/v1"
)

func ListDisks(_ context.Context) ([]*compute.Disk, error) {
	return []*compute.Disk{
		{
			Id:     "trdc7u8s75v5je2gpoth",
			Name:   "list-disk-test-1",
			Status: compute.Disk_STATUS_READY,
			Source: &compute.Disk_SourceImageId{
				SourceImageId: "iw92psl6l0dyp8ctoqm",
			},
		},
	}, nil
}

func GetDisk(_ context.Context, diskID string) (*compute.Disk, error) {
	return &compute.Disk{
		Id:     diskID,
		Name:   "get-disk-test-1",
		Status: compute.Disk_STATUS_READY,
		Source: &compute.Disk_SourceImageId{
			SourceImageId: "iw92psl6l0dyp8ctoqm",
		},
	}, nil
}
