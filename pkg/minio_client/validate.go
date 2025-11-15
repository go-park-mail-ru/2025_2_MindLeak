package minio_client

import (
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
)

const (
	maxAvatarSize = 50 * 1024 * 1024 // 5 MB
)

var allowedPictureExt = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".webp": true,
	".bmp":  true,
	".tiff": true,
	".svg":  true,
	".ico":  true,
}

func validateImage(header *multipart.FileHeader) error {
	ext := filepath.Ext(header.Filename)
	if !allowedPictureExt[ext] {
		return fmt.Errorf("unsupported file type: %s", ext)
	}
	if header.Size > maxAvatarSize {
		return errors.New("file size exceeds 5 MB limit")
	}
	return nil
}
