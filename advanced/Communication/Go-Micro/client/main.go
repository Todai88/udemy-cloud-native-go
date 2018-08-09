package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	proto "github.com/Todai88/udemy-cloud-native-go/advanced/Communication/Go-Micro/proto"
	hystrix "github.com/afex/hystrix-go/hystrix"
	micro "github.com/micro/go-micro"
	breaker "github.com/micro/go-plugins/wrapper/breaker/hystrix"
)

func main() {
	service := micro.NewService(
		micro.Name("greeter"),
		micro.Version("latest"),
		micro.Metadata(map[string]string{
			"type": "helloworld",
		}),
	)

	service.Init(
		micro.WrapClient(breaker.NewClientWrapper()),
	)

	greeter := proto.NewGreeterClient("greeter", service.Client())
	hystrix.DefaultVolumeThreshold = 3
	hystrix.DefaultErrorPercentThreshold = 75
	hystrix.DefaultTimeout = 500
	hystrix.DefaultSleepWindow = 3500

	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()

	go http.ListenAndServe(net.JoinHostPort("", "8081"), hystrixStreamHandler)

	callEvery(3*time.Second, greeter, hello)
}

func hello(t time.Time, greeter proto.GreeterClient) {
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Beth, calling at " + t.String()})
	if err != nil {
		if err.Error() == "hystrix: timeout" {
			fmt.Printf("%s. Insert fallback logic here \n", err.Error())
		} else {
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Printf("%s\n", rsp.Greeting)
}

func callEvery(d time.Duration, greeter proto.GreeterClient, f func(time.Time, proto.GreeterClient)) {
	for x := range time.Tick(d) {
		f(x, greeter)
	}
}
