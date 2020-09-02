package main

import (
	"log"
	"time"

	"github.com/leaf-rain/grpc-lb/balancer"
	"github.com/leaf-rain/grpc-lb/examples/proto"
	registry "github.com/leaf-rain/grpc-lb/registry/zookeeper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	registry.RegisterResolver("zk", []string{"10.0.101.68:2189"}, "/backend/services", "test", "1.0")
	c, err := grpc.Dial("zk:///", grpc.WithInsecure(), grpc.WithBalancerName(balancer.RoundRobin))
	if err != nil {
		log.Printf("grpc dial: %s", err)
		return
	}
	defer c.Close()
	client := proto.NewTestClient(c)

	for i := 0; i < 5000; i++ {
		resp, err := client.Say(context.Background(), &proto.SayReq{Content: "round robin"})
		if err != nil {
			log.Println("aa:", err)
			time.Sleep(time.Second)
			continue
		}
		time.Sleep(time.Second)
		log.Printf(resp.Content)
	}
}
