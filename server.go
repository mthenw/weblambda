package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/lambda"
	"github.com/gin-gonic/gin"
)

// Name of a function declared in AWS Lambda
const FunctionName string = "weblambda"

func server(region string) {
	svc := lambda.New(&aws.Config{Region: region})

	router := gin.Default()
	router.POST("/", func(c *gin.Context) {
		code, _ := ioutil.ReadAll(c.Request.Body)
		output, _ := invoke(svc, code)

		c.String(http.StatusOK, string(output.Payload))
	})
	router.Run(":8080")
}

func invoke(svc *lambda.Lambda, code []byte) (*lambda.InvokeOutput, error) {
	params := &lambda.InvokeInput{
		FunctionName: aws.String(FunctionName),
		Payload:      code,
	}

	resp, err := svc.Invoke(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	return resp, err
}
