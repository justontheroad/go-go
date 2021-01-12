package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "go-go/go-go.grpc/searchgoods"
)

type SearchGoodsServer struct{}

func (s *SearchGoodsServer) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	goods, err := findGoods(*r.Sku)
	if nil == err {
		return &pb.SearchResponse{Id: goods.id, Sku: goods.sku, Name: goods.name}, nil
	}

	return &pb.SearchResponse{Id: 0, Sku: "", Name: ""}, nil
}

func (s *SearchGoodsServer) List(r *pb.SearchRequest, stream pb.SearchGoods_ListServer) error {
	for _, goods := range demoData {
		if err := stream.Send(&pb.SearchResponse{Id: goods.id, Sku: goods.sku, Name: goods.name}); err != nil {
			return err
		}
	}

	return nil
}

func (s *SearchGoodsServer) Record(stream pb.SearchGoods_RecordServer) error {
	var rets []*pb.SearchRequest
	for {
		r, err := stream.Recv()
		rets = append(rets, r)
		if err == io.EOF {
			goods, err := findGoods(*rets[0].Sku)
			if nil == err {
				return stream.SendAndClose(&pb.SearchResponse{Id: goods.id, Sku: goods.sku, Name: goods.name})
			}

			return stream.SendAndClose(&pb.SearchResponse{Id: 0, Sku: "", Name: ""})
		}
	}

	return nil
}

func (s *SearchGoodsServer) Route(stream pb.SearchGoods_RouteServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Println(in)
		for _, goods := range demoData {
			if err := stream.Send(&pb.SearchResponse{Id: goods.id, Sku: goods.sku, Name: goods.name}); err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	initDemoData()

	grpcSev := grpc.NewServer()
	pb.RegisterSearchGoodsServer(grpcSev, &SearchGoodsServer{})
	grpcSev.Serve(lis)
}

type goodsInfo struct {
	id   int32
	sku  string
	name string
}

var (
	demoData = make(map[string]*goodsInfo)
)

func initDemoData() {
	demoData["123456789"] = &goodsInfo{id: 1, sku: "123456789", name: "test1"}
	demoData["223456789"] = &goodsInfo{id: 2, sku: "223456789", name: "test2"}
	demoData["323456789"] = &goodsInfo{id: 3, sku: "323456789", name: "test3"}
}

// func (s *SearchGoodsServer) demoData() map[string]*goodsInfo {
// 	demoData["123456789"] = &goodsInfo{id: 1, sku: "123456789", name: "test1"}
// 	demoData["223456789"] = &goodsInfo{id: 2, sku: "223456789", name: "test2"}
// 	demoData["323456789"] = &goodsInfo{id: 3, sku: "323456789", name: "test3"}

// 	return demoData
// }

func findGoods(sku string) (goods *goodsInfo, err error) {
	goods, ok := demoData[sku]
	if ok {
		return
	}

	err = errors.New("goods not found")
	// err = fmt.Errorf("goods not found")
	return
}
