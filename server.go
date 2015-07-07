package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/gin-gonic/gin"
)

// Name of a function declared in AWS Lambda
const FunctionName string = "weblambda"

type weblambdaRequest struct {
	Source string `json:"source"`
}

func server(region string, port string) {
	svc := lambda.New(&aws.Config{Region: region})

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/", func(c *gin.Context) {
		source, _ := ioutil.ReadAll(c.Request.Body)
		req := &weblambdaRequest{
			Source: string(source[:]),
		}

		output, err := invoke(svc, req)

		if err == nil {
			c.String(http.StatusOK, string(output.Payload))
		} else {
			c.String(http.StatusInternalServerError, string(output.Payload))
		}
	})
	router.Run(":" + port)
}

func invoke(svc *lambda.Lambda, req *weblambdaRequest) (*lambda.InvokeOutput, error) {
	reqJSON, _ := json.Marshal(req)

	params := &lambda.InvokeInput{
		FunctionName: aws.String(FunctionName),
		Payload:      reqJSON,
	}

	return svc.Invoke(params)
}
