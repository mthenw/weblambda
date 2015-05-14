package main

import (
	"archive/zip"
	"bytes"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/lambda"
)

var runtimeFunction = `
exports.handler = function(event, context) {
    eval(event.source);
};
`

func install(role string) {
	svc := lambda.New(&aws.Config{Region: "us-east-1"})

	params := &lambda.CreateFunctionInput{
		Code: &lambda.FunctionCode{
			ZipFile: zipRuntime(),
		},
		FunctionName: aws.String("weblambda"),
		Handler:      aws.String("handler"),
		Role:         aws.String(role),
		Runtime:      aws.String("nodejs"),
	}

	_, err := svc.CreateFunction(params)
	if err != nil {
		panic(err)
	}
}

func zipRuntime() []byte {
	buffer := bytes.NewBuffer(nil)
	archive := zip.NewWriter(buffer)

	fwriter, _ := archive.Create("index.js")
	fwriter.Write([]byte(runtimeFunction))

	archive.Close()

	return buffer.Bytes()
}
