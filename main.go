package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	graphql "github.com/graph-gophers/graphql-go"
	"github.com/scottxxx666/meetups-api/resolvers"
)

var mainSchema *graphql.Schema

func getSchema() (string, error) {
	path, _ := filepath.Abs("schema.graphql")
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Print("Get schema failed", err)
		return "", err
	}

	return string(b), nil
}

var s1 = `
	schema {
		query: Query
	}
	type Person{
		id: ID!
	}
	type Query{
		person(id: ID!): Person
	}
`

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

	s, err := getSchema()
	mainSchema = graphql.MustParseSchema(s, &resolvers.Resolver{})
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
