package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Echo instance
	e := echo.New()

	// .env読み込み
	err := godotenv.Load(fmt.Sprintf("./env/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		e.Logger.Fatal(err)
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.GET("/hello/task", runHelloTask)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}

// Handler
func hello(c echo.Context) error {
	return c.JSON(http.StatusOK,
		map[string]string{
			"success": "echo server is working!!",
		})
}

func runHelloTask(c echo.Context) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("ECS_REGION"))},
	)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	svc := ecs.New(sess)
	input := &ecs.RunTaskInput{
		Cluster:        aws.String(os.Getenv("ECS_CLUSTER")),         // クラスター名
		TaskDefinition: aws.String(os.Getenv("ECS_TASK_DEFINITION")), // バージョン指定しない場合はLATESTが選択される
	}

	input.NetworkConfiguration = &ecs.NetworkConfiguration{
		AwsvpcConfiguration: &ecs.AwsVpcConfiguration{
			Subnets: aws.StringSlice([]string{ // サブネットID
				os.Getenv("ECS_SUBNET_1"),
				os.Getenv("ECS_SUBNET_2"),
			}),
			AssignPublicIp: aws.String("ENABLED"), // 必要に応じて
		},
	}

	input.LaunchType = aws.String("FARGATE")

	result, err := svc.RunTask(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		}
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(result)
	return c.JSON(http.StatusOK,
		map[string]string{
			"success": "runHelloTask is success",
		})
}
