# weblambda

[![Build Status](https://travis-ci.org/mthenw/weblambda.svg?branch=master)](https://travis-ci.org/mthenw/weblambda)

## Installation

```
go get github.com/mthenw/weblambda
```

*Remember about [configuring credentials](https://github.com/awslabs/aws-sdk-go/#configuring-credentials)*

## Example

1. Create IAM role for weblambda function

2. Install weblambda function on AWS Lambda

        weblambda install --role <IAM role ARN> --region us-east-1

3. Run HTTP server

        weblambda server --region us-east-1

4. Run function

        $ curl 127.0.0.1:8080 --data-binary "var result = 'Hello world!'; context.succeed(result);"
        "Hello world!"

## Usage

```
$ weblambda help
NAME:
   weblambda - Run JavaScript code on AWS Lambda with HTTP call

USAGE:
   weblambda [global options] command [command options] [arguments...]

VERSION:
   0.0.0

COMMANDS:
   install	Install weblambda function on AWS Lambda
   server	Start HTTP server
   help, h	Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h		show help
   --version, -v	print the version

```

## AWS Lambda limits

Number requests per second is limited by [AWS Lambda limits](http://docs.aws.amazon.com/lambda/latest/dg/limits.html).