package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/feifeijun97/GenshinBackEnd/modules/character/characterpb"
	"google.golang.org/grpc"
	// "github.com/feifeijun97/GenshinBackEnd/modules/character/characterpb"
)

const (
	defaultName = "world"
)

var (
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	fmt.Println("im client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	fmt.Println("connected to tcp server")

	c := characterpb.NewCharacterListServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.CharacterList(ctx, &characterpb.CharacterListRequest{Name: name})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.GetName())
}
