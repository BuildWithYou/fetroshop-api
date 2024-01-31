package miniohelper

import (
	"context"
	"mime/multipart"

	"github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

type Minio struct {
	Client     *minio.Client
	BucketName string
}

func GetMinio(config *viper.Viper, logger *logger.Logger) *Minio {
	endpoint := config.GetString("minio.endpoint")
	accessKeyID := config.GetString("minio.accessKeyId")
	secretAccessKey := config.GetString("minio.secretAccessKey")
	useSSL := config.GetBool("minio.useSSL")
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		logger.Info("Minio client initialized failed")
		logger.Fatal(err.Error())
	}
	logger.Info("Minio client initialized successfully")

	return &Minio{
		Client:     minioClient,
		BucketName: config.GetString("minio.bucketName"),
	}
}

func (m *Minio) Upload(ctx context.Context, file *multipart.FileHeader, filePath string) (result minio.UploadInfo, err error) {
	// get buffer
	buffer, err := file.Open()
	if err != nil {
		return minio.UploadInfo{}, err
	}
	defer buffer.Close()

	contentType := file.Header["Content-Type"][0]
	fileSize := file.Size

	// Upload the zip file with PutObject
	return m.Client.PutObject(ctx, m.BucketName, filePath, buffer, fileSize, minio.PutObjectOptions{ContentType: contentType})
}

func (m *Minio) Remove(ctx context.Context, filePath string) (err error) {
	return m.Client.RemoveObject(ctx, m.BucketName, filePath, minio.RemoveObjectOptions{})
}
