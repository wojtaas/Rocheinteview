package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"rocheinteview"
	"rocheinteview/service"
)

var VERSION string

func main() {
	Run()
}

func Run() {
	var cfg rocheinteview.Config
	envconfig.MustProcess("", &cfg)

	sv := service.NewService(cfg, VERSION)
	api := service.NewApi(cfg.ServiceName, sv)
	server := service.NewServer(sv)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.HTTPPort),
		Handler: api.GetHandler(),
	}

	go func() {
		fmt.Printf("Starting http server on port: %d \n", cfg.HTTPPort)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Errorf("http ListenAndServe err: %s", err)
		}
	}()

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPCPort))
	if err != nil {
		fmt.Errorf("net.Listen() %s", err)
	}

	grpcServer := grpc.NewServer()
	server.Register(grpcServer)

	fmt.Printf("server grpc is listening on port: %d", cfg.GRPCPort)
	if err := grpcServer.Serve(listen); err != nil && err != grpc.ErrServerStopped {
		fmt.Errorf("grpcServer.Serve() %s", err)
	}
}
