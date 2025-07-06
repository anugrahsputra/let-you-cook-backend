package minio_util

import (
	"bytes"
	"context"
	"fmt"
	"let-you-cook/config"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

var minioClient *minio.Client

var imageExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
}

// InitMinio initializes the MinIO client
func InitMinio() error {
	var err error
	minioClient, err = config.InitMinio()
	if err != nil {
		return err
	}
	return nil
}

// GetMinioClient returns the initialized MinIO client
func GetMinioClient() *minio.Client {
	return minioClient
}

// SetPublicBucketPolicy sets a public read policy for the specified bucket
func SetPublicBucketPolicy(bucketName string) error {
	ctx := context.Background()

	// Check if bucket exists, create if it doesn't
	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if !exists {
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}

	// Define public read policy as a JSON string
	policy := `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":"*","Action":["s3:GetObject"],"Resource":["arn:aws:s3:::` + bucketName + `/*"]}]}`

	// Apply the policy
	err = minioClient.SetBucketPolicy(ctx, bucketName, policy)
	if err != nil {
		return err
	}

	return nil
}

// UploadFile uploads a file to the specified bucket and returns the upload info
func UploadFile(bucketName, objectName string, fileBuffer []byte, contentType string) (minio.UploadInfo, error) {
	ctx := context.Background()

	// Check if the bucket exists
	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return minio.UploadInfo{}, err
	}
	if !exists {
		// Create the bucket if it does not exist
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return minio.UploadInfo{}, err
		}
	}

	reader := bytes.NewReader(fileBuffer)

	uploadInfo, err := minioClient.PutObject(ctx, bucketName, objectName, reader, reader.Size(), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return minio.UploadInfo{}, err
	}

	return uploadInfo, nil
}

// GetObjectURL returns the direct URL to the object (public access)
func GetObjectURL(bucketName, objectName string) string {
	protocol := "http"
	if config.MinioUseSSL {
		protocol = "https"
	}
	return protocol + "://" + config.MinioEndpoint + "/" + bucketName + "/" + objectName
}

// TrimMinioURLPrefix removes the MinIO endpoint and bucket name from a given URL.
func TrimMinioURLPrefix(fullURL string, minioEndpoint string, minioUseSSL bool) string {
	if strings.HasPrefix(fullURL, "http") {
		prefix := ""
		if minioUseSSL {
			prefix = "https://" + minioEndpoint + "/let-you-cook/"
		} else {
			prefix = "http://" + minioEndpoint + "/let-you-cook/"
		}
		return strings.TrimPrefix(fullURL, prefix)
	}
	return fullURL
}

func UploadPhoto(file *multipart.FileHeader) (string, error) {
	bucketName := "let-you-cook"

	ctx := context.Background()

	// Check if bucket exists, create if it doesn't
	exists, err := minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return "", err
	}
	if !exists {
		err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return "", err
		}
		// Set public policy
		policy := `{"Version":"2012-10-17","Statement":[{"Effect":"Allow","Principal":"*","Action":["s3:GetObject"],"Resource":["arn:aws:s3:::` + bucketName + `/*"]}]}`
		err = minioClient.SetBucketPolicy(ctx, bucketName, policy)
		if err != nil {
			return "", err
		}
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Generate a unique object name
	extension := strings.ToLower(filepath.Ext(file.Filename))
	if !imageExtensions[extension] {
		return "", fmt.Errorf("file %s is not an image", extension)
	}
	objectName := fmt.Sprintf("%s-%d%s", uuid.New().String(), time.Now().Unix(), extension)

	// Upload the file
	_, err = minioClient.PutObject(ctx, bucketName, objectName, src, file.Size, minio.PutObjectOptions{ContentType: file.Header.Get("Content-Type")})
	if err != nil {
		return "", err
	}

	// Construct the URL
	url := GetObjectURL(bucketName, objectName)

	return url, nil
}

