package minio_client

import (
	"context"
	"fmt"
	"strings"

	minioSDK "github.com/minio/minio-go/v7"
)

func (c *Client) Delete(ctx context.Context, fileURL string) error {
	object := strings.TrimPrefix(fileURL, fmt.Sprintf("%s/%s/", c.publicURL, c.bucket))

	return c.sdk.RemoveObject(ctx, c.bucket, object, minioSDK.RemoveObjectOptions{})
}
