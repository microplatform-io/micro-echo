package main

import (
	"log"

	"github.com/microplatform-io/platform"
)

func main() {
	service, err := platform.NewBasicService()
	if err != nil {
		log.Fatalf("failed to create service: %s", err)
	}

	service.AddHandler("/platform/create/echo", platform.HandlerFunc(func(responseSender platform.ResponseSender, request *platform.Request) {
		responseSender.Send(platform.GenerateResponse(request, &platform.Request{
			Routing:   platform.RouteToUri("resource:///platform/reply/echo"),
			Context:   request.Context,
			Payload:   request.Payload,
			Completed: platform.Bool(true),
		}))
	}), 1)

	service.Run()
}
