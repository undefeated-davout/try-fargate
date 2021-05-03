package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/joho/godotenv"
)

func main() {
	// .env読み込み
	err := godotenv.Load(fmt.Sprintf("./env/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		fmt.Println("not found .env file", err)
	}

	bucketName := os.Getenv("BUCKET_NAME")
	if bucketName == "" { // .envにも環境変数にもなければエラー
		log.Fatal("not found BUCKET_NAME in environment variable")
	}

	// 環境変数が取得できているか確認用
	log.Println("BUCKET_NAME: ", bucketName)

	// アップロード用ファイルの作成
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	now := time.Now().In(jst)
	targetFilePath := "./sample.txt"
	if err := writeLine(targetFilePath); err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(targetFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		file.Close()
		if err := os.Remove(targetFilePath); err != nil {
			fmt.Println(err)
		}
	}()

	objectKey := fmt.Sprintf("./%s.txt", now.Format("2006-01-02_150405"))

	// Uploaderを作成し、ローカルファイルをアップロード
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   file,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("successflul S3 upload")
}

// サンプルファイル作成処理
func writeLine(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines = []string{"１２３４５\n", "あいうえお\n", "1234567890\n"}
	for _, line := range lines {
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}
	return nil
}
