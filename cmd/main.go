package main

import (
	"encoding/json"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// HealthCheckResponseBody estrutura de resposta da função Lambda
type HealthCheckResponseBody struct {
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
}

// handler função principal da Lambda
func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := HealthCheckResponseBody{
		Timestamp: time.Now().Unix(),
		Message:   "Lambda está funcionando corretamente!",
	}

	body, err := json.Marshal(response)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Erro ao gerar a resposta",
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(body),
	}, nil
}

func main() {
	lambda.Start(handler)
}
