package httpserver

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/yazdanbhd/Music-Cloud/repository/mysqldb"
	"github.com/yazdanbhd/Music-Cloud/service/userservice"
	"log"
	"net/http"
	"path/filepath"
)

func (s Server) UserRegister(c echo.Context) error {
	var req userservice.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	db, err := mysqldb.New(s.cfg.DataBase)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)

	}

	userSvc := userservice.New(db)

	response, err := userSvc.UserRegister(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)

	}

	return c.JSON(http.StatusCreated, response)
}

func (s Server) UserLogin(c echo.Context) error {
	var req userservice.LoginRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	db, err := mysqldb.New(s.cfg.DataBase)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)

	}

	userSvc := userservice.New(db)

	response, err := userSvc.UserLogin(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}
	return c.JSON(http.StatusOK, response)
}

func getClaims(c echo.Context) *jwt.MapClaims {
	return c.Get("user").(*jwt.MapClaims)
}

func (s Server) UploadMusic(c echo.Context) error {
	// Get the access token from the header
	//accessToken := c.Request().Header.Get("Authorization")
	//accessToken = strings.Replace(accessToken, "Bearer ", "", -1)
	//
	//token := authjwt.New([]byte(`secret-key`), jwt.SigningMethodHS256)
	//claims, err := token.VerifyToken(accessToken)
	//
	//if err != nil {
	//	return echo.NewHTTPError(http.StatusUnauthorized)
	//}

	// Minio Connection
	ctx := context.Background()

	// Initialize minio client object.

	minioClient, err := minio.New(s.cfg.MinioS3.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(s.cfg.MinioS3.AccessKeyID, s.cfg.MinioS3.SecretAccessKey, ""),
		Secure: s.cfg.MinioS3.UserSSL,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err)
	}

	bucketName := "music"
	location := "us-east-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			return echo.NewHTTPError(http.StatusBadGateway, err)
		}
	}

	file, metaData, err := c.Request().FormFile("file")

	fileExt := filepath.Ext(metaData.Filename)
	if fileExt != ".mp3" {
		return echo.NewHTTPError(http.StatusBadRequest, "The file is not in mp3 format")
	}

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	// TODO - Apply limitation for the size of upload file. => We can use API Gateway or apply in code.

	claims := getClaims(c)

	username := (*claims)["username"].(string)

	objectName := fmt.Sprintf("%s/%s", username, metaData.Filename)
	//filePath := "/tmp/testdata"
	contentType := "application/octet-stream"

	info, err := minioClient.PutObject(ctx, bucketName, objectName, file, -1, minio.PutObjectOptions{ContentType: contentType})

	if err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, err)
	}

	return c.JSON(http.StatusCreated, info)
}
