package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "go-go/go-go.grpc/searchgoods"
)

func main() {
	conn, err := grpc.Dial(fmt.Sprintf("localhost:%d", 9000), grpc.WithInsecure()) // grpc: no transport security set (use grpc.WithInsecure() explicitly or set credentials)

	if nil != err {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewSearchGoodsClient(conn)

	var (
		sku string = "123456789"
		id  int32  = 1
	)
	// 一元RPC
	resp, err := client.Search(context.Background(), &pb.SearchRequest{Sku: &sku, Id: &id})
	if err != nil {
		log.Fatalf("%v.Search(_) = _, %v: ", client, err)
	}
	log.Println(resp)

	// 服务器端流式 RPC
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.List(ctx, &pb.SearchRequest{Sku: &sku, Id: &id})
	if err != nil {
		log.Fatalf("%v.list(_) = _, %v: ", stream, err)
	}
	for {
		goods, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.list(_) = _, %v", client, err)
		}
		log.Println(goods)
	}

	// 客户端流式 RPC
	cliStream, err := client.Record(context.Background())
	if err != nil {
		log.Fatalf("%v.Record(_) = _, %v", client, err)
	}
	var rets []*pb.SearchRequest
	rets = append(rets, &pb.SearchRequest{Sku: &sku, Id: &id})
	rets = append(rets, &pb.SearchRequest{Sku: &sku, Id: &id})
	for _, r := range rets {
		cliStream.Send(r)
	}

	cliResp, err := cliStream.CloseAndRecv()
	log.Println(cliResp)

	// 双向流式 RPC
	dbStream, err := client.Route(context.Background())
	for _, r := range rets {
		dbStream.Send(r)
		resp, err := dbStream.Recv()
		log.Println(resp)
		if err == io.EOF {
			break
		}
	}
}
