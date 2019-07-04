package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lucid-bunch/gone-alexa/alexa"
)

// Handler - entry point for Lambda Alexa Skill
func Handler(request alexa.Request) (alexa.Response, error) {
	return intentDispatcher(request), nil
}

func main() {
	lambda.Start(Handler)
}

func handleGoneSearchEvent(request alexa.Request) alexa.Response {
	return alexa.NewSimpleResponse("Search Response", "You wanted to know about "+request.Body.Intent.Slots["query"].Value)
}

func handleDefaultIntent(request alexa.Request) alexa.Response {
	var builder alexa.SSMLBuilder
	builder.Say("Welcome to Gone.")
	builder.Pause("2000")
	builder.Say("Invoke this skill by saying. Alexa open gone and search followed by what you want to know more about.")
	return alexa.NewSSMLResponse("Gone Default", builder.Build())
}

func intentDispatcher(request alexa.Request) alexa.Response {
	var response alexa.Response
	switch request.Body.Intent.Name {
	case "GoneSearchIntent":
		response = handleGoneSearchEvent(request)
	default:
		response = handleDefaultIntent(request)
	}
	return response
}
