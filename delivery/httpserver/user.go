package httpserver

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/yazdanbhd/Music-Cloud/delivery/authjwt"
	"github.com/yazdanbhd/Music-Cloud/repository/mysqldb"
	"github.com/yazdanbhd/Music-Cloud/service/userservice"
	"log"
	"net/http"
	"strings"
)

func (s Server) UserRegister(c echo.Context) error {
	var req userservice.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	db, err := mysqldb.New(s.dbConfig)
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

	db, err := mysqldb.New(s.dbConfig)
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

func (s Server) UploadMusic(c echo.Context) error {
	// Get the access token from the header
	accessToken := c.Request().Header.Get("Authorization")
	accessToken = strings.Replace(accessToken, "Bearer ", "", -1)

	token := authjwt.New([]byte(`secret-key`), jwt.SigningMethodHS256)
	verifyErr := token.VerifyToken(accessToken)

	if verifyErr != nil {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	// Minio Connection
	ctx := context.Background()

	endpoint := "localhost:9000"
	accessKeyID := "tcazzdiIuNJmgmCjXVdm"
	secretAccessKey := "WV2oRUF3NkRadBBuJvo6ifXBGM67roVxyXTeBcbD"

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucket called testbucket.
	bucketName := "music"
	location := "us-east-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}

	file, header, err := c.Request().FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	fmt.Println(header)

	// Upload the test file
	// Change the value of filePath if the file is in another location
	objectName := "testdata.jpg"
	//filePath := "/tmp/testdata"
	contentType := "application/octet-stream"

	info, err := minioClient.PutObject(ctx, bucketName, objectName, file, 16, minio.PutObjectOptions{ContentType: contentType})

	// Upload the test file with FPutObject
	//info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)

	return c.JSON(http.StatusCreated, info)
}
