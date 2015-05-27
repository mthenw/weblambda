package main

import (
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
		output, err := invoke(svc, code)

		if err == nil {
			c.String(http.StatusOK, string(output.Payload))
		} else {
			c.String(http.StatusInternalServerError, string(output.Payload))
		}
	})
	router.Run(":8080")
}

func invoke(svc *lambda.Lambda, code []byte) (*lambda.InvokeOutput, error) {
	params := &lambda.InvokeInput{
		FunctionName: aws.String(FunctionName),
		Payload:      code,
	}

	return svc.Invoke(params)
}
