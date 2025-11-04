package minio_client

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"mime/multipart"
	"path/filepath"
)

func (c *Client) UploadAvatar(ctx context.Context, userID uuid.UUID, file multipart.File, header *multipart.FileHeader) (string, error) {
	if err := validateImage(header); err != nil {
		return "", err
	}

	filename := fmt.Sprintf("avatars/%s%s", userID, filepath.Ext(header.Filename))
	_, err := c.sdk.PutObject(ctx, c.bucket, filename, file, header.Size, minio.PutObjectOptions{
		ContentType: header.Header.Get("Content-Type"),
	})
	if err != nil {
		return "", fmt.Errorf("upload avatar: %w", err)
	}

	return fmt.Sprintf("https://mindleak.ru/%s/%s/%s", c.publicURL, c.bucket, filename), nil
}

func (c *Client) UploadCover(ctx context.Context, userID uuid.UUID, file multipart.File, header *multipart.FileHeader) (string, error) {
	if err := validateImage(header); err != nil {
		return "", err
	}

	filename := fmt.Sprintf("covers/%s%s", userID, filepath.Ext(header.Filename))
	_, err := c.sdk.PutObject(ctx, c.bucket, filename, file, header.Size, minio.PutObjectOptions{
		ContentType: header.Header.Get("Content-Type"),
	})
	if err != nil {
		return "", fmt.Errorf("upload cover: %w", err)
	}

	return fmt.Sprintf("https://mindleak.ru/%s/%s/%s", c.publicURL, c.bucket, filename), nil
}
