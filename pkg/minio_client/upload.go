package minio_client

import (
	"context"
	"fmt"
	"mime/multipart"
	"path/filepath"

	minioSDK "github.com/minio/minio-go/v7"
)

func (c *Client) Upload(ctx context.Context, folder string, file multipart.File, filename string) (string, error) {
	object := fmt.Sprintf("%s/%s", folder, filepath.Base(filename))

	_, err := c.sdk.PutObject(ctx, c.bucket, object, file, -1, minioSDK.PutObjectOptions{})
	if err != nil {
		return "", fmt.Errorf("upload file: %w", err)
	}

	return fmt.Sprintf("%s/%s/%s", c.publicURL, c.bucket, object), nil
}
