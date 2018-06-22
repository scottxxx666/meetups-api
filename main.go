package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	graphql "github.com/graph-gophers/graphql-go"
)

var mainSchema *graphql.Schema

type resolver struct{}

var Schema = `
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

type person struct {
	FirstName graphql.ID
}

type personResolver struct {
	p *person
}

func (r *personResolver) ID() graphql.ID {
	return r.p.FirstName
}

func (r *resolver) Person(args struct{ ID graphql.ID }) *personResolver {
	p := person{"AAA"}
	return &personResolver{&p}
}

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

	mainSchema = graphql.MustParseSchema(Schema, &resolver{})
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
