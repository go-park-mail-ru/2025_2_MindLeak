package minio_client

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func (c *Client) DeleteAvatar(ctx context.Context, userID uuid.UUID) error {
	filename := fmt.Sprintf("avatars/%s", userID)
	return c.sdk.RemoveObject(ctx, c.bucket, filename, minio.RemoveObjectOptions{})
}

func (c *Client) DeleteCover(ctx context.Context, userID uuid.UUID) error {
	filename := fmt.Sprintf("covers/%s", userID)
	return c.sdk.RemoveObject(ctx, c.bucket, filename, minio.RemoveObjectOptions{})
}
