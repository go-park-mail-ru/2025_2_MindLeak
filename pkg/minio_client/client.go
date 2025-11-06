package minio_client

import (
	"fmt"

	minioSDK "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Client struct {
	sdk       *minioSDK.Client
	bucket    string
	publicURL string
}

func NewClient(endpoint, user, password, bucket, publicURL string) (*Client, error) {
	sdk, err := minioSDK.New(endpoint, &minioSDK.Options{
		Creds:  credentials.NewStaticV4(user, password, ""),
		Secure: false,
	})
	if err != nil {
		return nil, fmt.Errorf("init minio_client client: %w", err)
	}

	return &Client{
		sdk:       sdk,
		bucket:    bucket,
		publicURL: publicURL,
	}, nil
}

func (c *Client) GetDefaultAvatar() string {
	return fmt.Sprintf("https://mindleak.ru/%s/%s/defaultAvatar.svg", c.publicURL, c.bucket)
}

func (c *Client) GetDefaultCover() string {
	return fmt.Sprintf("https://mindleak.ru/%s/%s/cover-pic.jpg", c.publicURL, c.bucket)
}

func (c *Client) GetDefaultMedia() string {
	return fmt.Sprintf("https://mindleak.ru/%s/%s/default.jpg", c.publicURL, c.bucket)
}
