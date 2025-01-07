package config

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"let-you-cook/utils/helper"
	"mime"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client

func InitializeMinioClient() error {
	endpoint := helper.GetENV("MINIO_URI", "127.0.0.1:9000")
	accessKeyID := helper.GetENV("MINIO_ACCESS", "minio")
	secretAccessKey := helper.GetENV("MINIO_SECRET", "minio123")

	logger.Infof("initializing MinIO client with endpoint: %s\n", endpoint)

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		logger.Errorf("error initializing MinIO client: %v\n", err)
		return fmt.Errorf("error initializing MinIO client: %v", err)
	}

	minioClient = client

	logger.Infof("initialized MinIO client: %v\n", minioClient)

	return nil
}

func CheckOrCreateBucket(bucketName string) error {
	// Check if the bucket exists
	exists, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		return fmt.Errorf("error checking if bucket exists: %v", err)
	}

	// If the bucket does not exist, create it
	if !exists {
		err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return fmt.Errorf("error creating bucket: %v", err)
		}
		fmt.Printf("Bucket %s created successfully.\n", bucketName)
	} else {
		fmt.Printf("Bucket %s already exists.\n", bucketName)
	}

	return nil
}

func DownloadFile(url string) ([]byte, error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept", "video/mp4, */*")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the data into a byte slice
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// UploadBytes uploads a byte slice to MinIO with the specified bucket and object names.
func UploadBytes(data []byte, bucketName, filename, url string) (string, error) {
	// Convert byte slice to io.Reader
	reader := bytes.NewReader(data)
	ext := GetFileExtension(url)
	fmt.Println("ext a: ", ext)
	objectName := AddFileExtension(filename, ext)

	// Upload the byte slice to MinIO
	// _, err := minioClient.PutObject(context.Background(), bucketName, objectName, reader, int64(reader.Len()), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	contentType := http.DetectContentType(data)
	_, err := minioClient.PutObject(context.Background(), bucketName, objectName, reader, int64(reader.Len()), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", fmt.Errorf("error uploading to MinIO: %v", err)
	}

	return objectName, nil
}

func UploadBytesWithCookie(data []byte, bucketName, filename, url string) (string, error) {
	// Convert byte slice to io.Reader
	reader := bytes.NewReader(data)
	// ext := GetFileExtension(url)
	// fmt.Println("ext a: ", ext)
	// objectName := AddFileExtension(filename, ext)

	// Upload the byte slice to MinIO
	// _, err := minioClient.PutObject(context.Background(), bucketName, objectName, reader, int64(reader.Len()), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	contentType := http.DetectContentType(data)
	_, err := minioClient.PutObject(context.Background(), bucketName, filename, reader, int64(reader.Len()), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", fmt.Errorf("error uploading to MinIO: %v", err)
	}

	return filename, nil
}

func GetFileExtension(url string) string {
	ext := ""
	// Try to determine extension from content type
	resp, err := http.Head(url)
	if err == nil {
		contentType := resp.Header.Get("Content-Type")
		fmt.Println("content type: ", contentType)
		if contentType != "" {
			exts, _ := mime.ExtensionsByType(contentType)
			if len(exts) > 0 {
				ext = exts[0]
			}

			if contentType == "image/jpeg" {
				ext = ".jpeg"
			}

			if contentType == "application/json; charset=utf-8" {
				ext = ".mp4"
			}
		}
	}
	return ext
}

// AddFileExtension adds an extension to a file name
func AddFileExtension(fileName, ext string) string {
	if !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + ext
}

func DownloadFileWithCookie(url, cookieValue, originValue, refererValue string) ([]byte, string, error) {
	// Buat HTTP client baru
	client := &http.Client{}

	// Buat request baru
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, "", err
	}

	// Tambahkan cookie ke header
	req.Header.Set("User-Agent", "PostmanRuntime/7.28.4")
	req.Header.Set("Accept", "video/mp4, */*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Cookie", cookieValue)
	req.Header.Set("Origin", originValue)
	req.Header.Set("Referer", refererValue)

	// Kirim request
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	// Periksa apakah request berhasil
	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("bad status: %s", resp.Status)
	}

	ext := ""

	contentType := resp.Header.Get("Content-Type")
	fmt.Println("content type: ", contentType)

	if contentType != "" {
		exts, _ := mime.ExtensionsByType(contentType)
		if len(exts) > 0 {
			ext = exts[0]
		}

		if contentType == "image/jpeg" {
			ext = ".jpeg"
		} else if contentType == "video/mp4" {
			ext = ".mp4"
		}
	}

	// Baca data ke dalam byte slice
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, "", err
	}

	return buf.Bytes(), ext, nil
}

func UploadImageToMinio(file *multipart.FileHeader, bucketName string) (string, error) {
	fileName := uuid.New().String() + filepath.Ext(file.Filename)
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	_, err = minioClient.PutObject(context.Background(), bucketName, fileName, src, -1, minio.PutObjectOptions{
		ContentType: file.Header.Get("Content-Type"),
	})
	if err != nil {
		return "", err
	}

	return fileName, nil
}
