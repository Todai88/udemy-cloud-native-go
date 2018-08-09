package main

import (
	"fmt"
	"time"

	proto "github.com/Todai88/udemy-cloud-native-go/advanced/Communication/Go-Micro/proto"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

type Greeter struct{}

var counter int

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	counter++

	if counter > 7 && counter < 15 {
		time.Sleep(1000 * time.Millisecond)
	} else {
		time.Sleep(100 * time.Millisecond)
	}

	rsp.Greeting = "Hello " + req.Name
	fmt.Printf("Responding with %s\n", rsp.Greeting)
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("1.0.0"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
