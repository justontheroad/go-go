package main

import (
	"context"
	"flag"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "go-go/go-go.grpc/routeguide"
)

var (
	serverAddr = flag.String("server_addr", "localhost:10000", "The server address in the format of host:port")
)

func main() {
	// 创建存根
	// conn, err := grpc.Dial(*serverAddr)
	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure()) // grpc: no transport security set (use grpc.WithInsecure() explicitly or set credentials)
	if nil != err {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)
	// 调用服务方法
	// 简单 RPC
	feature, err := client.GetFeature(context.Background(), &pb.Point{Latitude: 409146138, Longitude: -746188906})
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v: ", client, err)
	}
	log.Println(feature)

	rect := &pb.Rectangle{
		Lo: &pb.Point{Latitude: 400000000, Longitude: -750000000},
		Hi: &pb.Point{Latitude: 420000000, Longitude: -730000000},
	}

	// 服务器端流式 RPC
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, err := client.ListFeatures(ctx, rect)
	if err != nil {
		log.Fatalf("%v.GetFeatures(_) = _, %v: ", stream, err)
	}
	for {
		feature, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", client, err)
		}
		log.Printf("Feature: name: %q, point:(%v, %v)", feature.GetName(),
			feature.GetLocation().GetLatitude(), feature.GetLocation().GetLongitude())
	}

	notes := []*pb.RouteNote{
		{Location: &pb.Point{Latitude: 0, Longitude: 1}, Message: "First message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 2}, Message: "Second message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 3}, Message: "Third message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 1}, Message: "Fourth message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 2}, Message: "Fifth message"},
		{Location: &pb.Point{Latitude: 0, Longitude: 3}, Message: "Sixth message"},
	}

	// 双向流式 RPC
	dstream, err := client.RouteChat(context.Background())
	waitc := make(chan struct{})

	go func() {
		for {
			in, err := dstream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a note : %v", err)
			}
			log.Printf("Got message %s at point(%d, %d)", in.Message, in.Location.Latitude, in.Location.Longitude)
		}
	}()
	for _, note := range notes {
		if err := dstream.Send(note); err != nil {
			log.Fatalf("Failed to send a note: %v", err)
		}
	}
	dstream.CloseSend()
	<-waitc
}
