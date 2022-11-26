package main

import (
	"context"
	"log"
	"net"

	pb "github.com/examplev2/product"
	sh "github.com/examplev2/shop"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	// this will allow forward compatibility
	// but new function will always return error unless implemented in code
	pb.UnimplementedChatServiceServer
}

func (s *Server) AddProduct(ctx context.Context, in *pb.AddProduct) (*pb.AddProductResp, error) {
	log.Printf("Receive message: %s", in.GetBody())
	a := sh.Shop{}
	return &product.AddProductResp{}, nil
}

func (s *Server) GetProduct(ctx context.Context, in *product.GetProductReq) (*product.GetProductResp, error) {
	log.Printf("Receive message: %s", in.GetBody())
	return &product.GetProductResp{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8999")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Server{}

	grpcServer := grpc.NewServer()

	product.RegisterChatServiceServer(grpcServer, &s)

	// register server using reflection
	reflection.Register(grpcServer)

	log.Println("start listening on port: 8999")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
