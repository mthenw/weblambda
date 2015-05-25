# weblambda

## Example

1. Install weblambda function on AWS Lambda

        weblambda install --role <IAM role ARN> --region us-east-1

2. Run HTTP server

        weblambda server --region us-east-1

3. Run function

		$ http POST 127.0.0.1:8080 source="{\"source\":\"context.succeed({'test': 'Hello world'});\"}"
		HTTP/1.1 200 OK
		Content-Length: 22
		Content-Type: text/plain; charset=utf-8
		Date: Sat, 23 May 2015 10:51:00 GMT

		{"test":"Hello world"}
