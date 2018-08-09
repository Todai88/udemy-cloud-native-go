package main

import (
	"context"
	"fmt"
	"time"

	proto "github.com/Todai88/udemy-cloud-native-go/advanced/Communication/Go-Micro/proto"
	micro "github.com/micro/go-micro"
)

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	service.Init()

	greeter := proto.NewGreeterClient("greeter", service.Client())
	callEvery(3*time.Second, greeter, hello)
}

func hello(t time.Time, greeter proto.GreeterClient) {
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Beth, calling at " + t.String()})
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("%s\n", rsp.Greeting)
}

func callEvery(d time.Duration, greeter proto.GreeterClient, f func(time.Time, proto.GreeterClient)) {
	for x := range time.Tick(d) {
		f(x, greeter)
	}
}
