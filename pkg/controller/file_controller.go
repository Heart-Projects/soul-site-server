package controller

import (
	"com.sj/admin/pkg/entity/vo"
	"com.sj/admin/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
	"net/http"
	"path/filepath"
	"time"
)

type IFileController interface {
	UploadFile(c *gin.Context)
}

type fileController struct {
}

func NewFileController() IFileController {
	return &fileController{}
}

var SupportedContentTypes = []string{"image/jpeg", "image/png", "image/gif", "image/bmp", "image/webp"}

func (f *fileController) UploadFile(c *gin.Context) {
	//ctx := context.Background()
	uploadFile, fileHeader, fileErr := c.Request.FormFile("file")
	if fileErr != nil {
		logrus.Errorf("MinIO upload failed: %v", fileErr)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("MinIO upload failed: %v", fileErr)})
		return
	}
	uploadFileContentType := fileHeader.Header.Get("Content-Type")
	// 是否支持的类型
	isSupported := utils.Contains(SupportedContentTypes, uploadFileContentType)
	if !isSupported {
		logrus.Errorf("Upload file type is not supported: %v", uploadFileContentType)
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Upload file type is not supported: %v", uploadFileContentType)})
		return
	}
	// 获取原始文件名的后缀
	originalFileName := fileHeader.Filename
	uploadFileExtension := filepath.Ext(originalFileName)
	newFileName := fmt.Sprintf("upload-%d%s", time.Now().UnixNano(), filepath.Ext("file")) + uploadFileExtension
	logrus.Infof("Upload file is: %v, new file name is: %v ", originalFileName, newFileName)
	endpoint := "minio-api.k8s.qc.host.dxy"
	accessKeyID := "aV6BOgTxIVEV2RDXnrWS"
	secretAccessKey := "pTu3Pdju0gSpulOQUMBlcw6RBS9rjXnRqNwxCM8D"
	useSSL := true

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		logrus.Info(err)
	}

	// Make a new bucket called testbucket.
	bucketName := "file-image"

	//location := "us-east-1"
	//err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	//if err != nil {
	//	// Check to see if we already own this bucket (which happens if you run this twice)
	//	exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
	//	if errBucketExists == nil && exists {
	//		log.Printf("We already own %s\n", bucketName)
	//	} else {
	//		log.Fatalln(err)
	//	}
	//} else {
	//	log.Printf("Successfully created %s\n", bucketName)
	//}

	// Upload the test file
	// Change the value of filePath if the file is in another location
	objectName := "testdata"
	//filePath := "/Users/shenjin/backend-project/go-project/soul-site/PEADME.md"

	// Upload the test file with FPutObject
	info, err := minioClient.PutObject(c, bucketName, newFileName, uploadFile, -1, minio.PutObjectOptions{ContentType: uploadFileContentType})
	fileURL := fmt.Sprintf("http://%s/%s/%s", endpoint, bucketName, newFileName)

	//info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		logrus.Errorf(err.Error())
	}
	fileInfo := &vo.UploadedFile{
		Key:  newFileName,
		Url:  fileURL,
		Name: originalFileName,
		Size: info.Size,
		Type: uploadFileContentType,
	}
	logrus.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	utils.DoResponseSuccessWithData(c, fileInfo)
}
