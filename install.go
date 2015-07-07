package main

import (
	"archive/zip"
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
)

var runtimeFunction = `
exports.handler = function(event, context) {
	eval(event.source);
};
`

func install(role string, region string, upgrade bool) {
	svc := lambda.New(&aws.Config{Region: region})

	if functionExists(svc) {
		if upgrade {
			deleteFunction(svc)
		} else {
			println("Function already exits")
		}
	}

	if upgrade {
		createFunction(svc, role)
	}
}

func functionExists(svc *lambda.Lambda) bool {
	params := &lambda.GetFunctionInput{
		FunctionName: aws.String(FunctionName),
	}

	_, err := svc.GetFunction(params)

	return err == nil
}

func createFunction(svc *lambda.Lambda, role string) {
	params := &lambda.CreateFunctionInput{
		Code: &lambda.FunctionCode{
			ZipFile: zipRuntime(),
		},
		FunctionName: aws.String(FunctionName),
		Handler:      aws.String("index.handler"),
		Runtime:      aws.String("nodejs"),
		Role:         aws.String(role),
	}

	_, err := svc.CreateFunction(params)
	if err != nil {
		panic(err)
	}
}

func deleteFunction(svc *lambda.Lambda) {
	params := &lambda.DeleteFunctionInput{
		FunctionName: aws.String(FunctionName),
	}

	_, err := svc.DeleteFunction(params)
	if err != nil {
		panic(err)
	}
}

func zipRuntime() []byte {
	buf := bytes.NewBuffer(nil)
	arch := zip.NewWriter(buf)

	fwriter, _ := arch.Create("index.js")
	fwriter.Write([]byte(runtimeFunction))

	arch.Close()
	return buf.Bytes()
}
