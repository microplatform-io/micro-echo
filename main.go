package main

import (
	"log"

	"github.com/microplatform-io/platform"
	"strconv"
	"time"
)

var delay = platform.Getenv("DELAY", "0")

func main() {
	service, err := platform.NewBasicService("micro-echo")
	if err != nil {
		log.Fatalf("failed to create service: %s", err)
	}

	service.AddHandler("/platform/create/echo", platform.HandlerFunc(func(responseSender platform.ResponseSender, request *platform.Request) {
		d, err := strconv.Atoi(delay)
		if err != nil {
			d = 0
		}

		time.Sleep(time.Duration(d) * time.Second)

		responseSender.Send(platform.GenerateResponse(request, &platform.Request{
			Routing:   platform.RouteToUri("resource:///platform/reply/echo"),
			Context:   request.Context,
			Payload:   request.Payload,
			Completed: platform.Bool(true),
		}))
	}))

	service.AddHandler("/platform/get/documentation", platform.HandlerFunc(func(responseSender platform.ResponseSender, request *platform.Request) {
		responseSender.Send(platform.GenerateResponse(request, &platform.Request{
			Routing: platform.RouteToUri("resource:///platform/reply/documentation"),
			Payload: GetProtoBytes(&platform.Documentation{
				ServiceRoutes: []*platform.ServiceRoute{
					&platform.ServiceRoute{
						Description: platform.String("Create Echo "),
						Request:     &platform.Route{Uri: platform.String("microservice:///platform/create/echo")},
						Responses: []*platform.Route{
							&platform.Route{Uri: platform.String("resource:///platform/reply/echo")},
						},
						Version: platform.String("1.0"),
					},
				},
			}),
			Completed: platform.Bool(true),
		}))
	}))

	service.Run()
}

func GetProtoBytes(message platform.Message) []byte {
	protoBytes, _ := platform.Marshal(message)
	return protoBytes
}
