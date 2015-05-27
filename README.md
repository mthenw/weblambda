# weblambda

[![Build Status](https://travis-ci.org/mthenw/weblambda.svg?branch=master)](https://travis-ci.org/mthenw/weblambda)

## Installation

        go get github.com/mthenw/weblambda

**Remember about [configuring credentials](https://github.com/awslabs/aws-sdk-go/#configuring-credentials)**

## Usage

1. Install weblambda function on AWS Lambda

        weblambda install --role <IAM role ARN> --region us-east-1

2. Run HTTP server

        weblambda server --region us-east-1

3. Run function

        $ curl 127.0.0.1:8080 --data-binary "{\"source\":\"context.succeed({'test': 'Hello world'});\"}"
        {"test":"Hello world"}%
