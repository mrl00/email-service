package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	appconfig "github.com/mrl00/email-service/src/app_config"
	"github.com/mrl00/email-service/src/web"
)

func main() {
	appConfig, err := appconfig.GetAppConfiguration()
	if err != nil {
		panic("cannot load app configuration")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/health_check", web.HealthCheck)

	server := &http.Server{
		Handler: mux,
	}

	listener, err := net.Listen("tcp", ":"+appConfig.Application.Port)
	if err != nil {
		log.Fatalf("Listener failed: %v", err)
	}

	addr := listener.Addr().String()
	fmt.Printf("listen to: %v", addr)

	go func() {
		if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	select {}
}
