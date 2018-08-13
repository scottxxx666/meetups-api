package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/scottxxx666/meetups-api/resolver"
	"github.com/scottxxx666/meetups-api/schema"
)

func Handler(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}

	if err := json.Unmarshal([]byte(request.Body), &params); err != nil {
		log.Print("Could not decode body", err)
	}

	mainSchema := graphql.MustParseSchema(schema.GetSchema(), &resolver.Resolver{})
	response := mainSchema.Exec(context, params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)

	if err != nil {
		log.Print("Could not decode body")
	}

	return events.APIGatewayProxyResponse{Body: string(responseJSON), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
