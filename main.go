package main

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

type HealthCheckResponseBody struct {
	Timestamp time.Duration
	Message   string `json:"message"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	res, _ := json.Marshal(&HealthCheckResponseBody{
		Timestamp: time.Nanosecond,
		Message:   "Lambda OK!",
	})

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(res),
	}, nil
}
