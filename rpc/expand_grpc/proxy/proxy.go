package main

import (
	"context"
	pb "github.com/feitianlove/multiplePracticeTaking/rpc/expand_grpc"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net/http"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	err := pb.RegisterRestServiceHandlerFromEndpoint(ctx, mux, "localhost:5000", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		panic(err)
	}
	http.ListenAndServe(":8080", mux)
}
