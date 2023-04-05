package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(request events.ALBTargetGroupRequest) (events.ALBTargetGroupResponse, error) {
	// Replace the URL with your target API URL
	redirectURL := os.Getenv("APIURL")
	statusCode := 307 // Change this to 308 for a 308 redirect

	response := events.ALBTargetGroupResponse{
		StatusCode: statusCode,
		StatusDescription: "307 Temporary Redirect", // Change this to "308 Permanent Redirect" for a 308 redirect
		Headers: map[string]string{
			"Location": redirectURL,
			"Content-Type": "text/plain",
		},
		Body: "Redirecting...",
	}

	return response, nil
}

func main() {
	lambda.Start(HandleRequest)
}
