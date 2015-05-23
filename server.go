package main

import (
	"fmt"
	"net/http"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/lambda"
	"github.com/gin-gonic/gin"
)

type codeToInvoke struct {
	Source string
}

func server() {
	svc := lambda.New(&aws.Config{Region: "us-east-1"})

	router := gin.Default()
	router.POST("/", func(c *gin.Context) {
		var fnc codeToInvoke
		c.Bind(&fnc)

		output, _ := invoke(svc, &fnc)

		c.String(http.StatusUnauthorized, string(output.Payload))
	})
	router.Run(":8080")
}

func invoke(svc *lambda.Lambda, fnc *codeToInvoke) (*lambda.InvokeOutput, error) {
	params := &lambda.InvokeInput{
		FunctionName: aws.String("weblambda"),
		Payload:      []byte(fnc.Source),
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
