package grpc

import (
	"fmt"
	"github.com/willianSteffler/libsrv/data"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
	"log"
	"net"
)

func Listen(port int, db *gorm.DB) error {

	lis, err := net.Listen("tcp4", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("erro ao criar conexão grpc: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	libServer := GrpcLibServer{}
	libServer.Init(db)
	data.RegisterLibraryServer(grpcServer, &libServer)
	reflection.Register(grpcServer)

	log.Printf("grpc aguarando conexões em %s",lis.Addr().String())

	err = grpcServer.Serve(lis)

	return err
}
