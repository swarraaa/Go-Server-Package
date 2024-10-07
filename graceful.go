package goServer

import (
	"context"
	"log"
	"net/http"
)

func (s *Server) Shutdown(ctx context.Context) {
    srv := &http.Server{Addr: ":8080"} 
    if err := srv.Shutdown(ctx); err != nil {
        log.Fatalf("Server Shutdown Failed: %+v", err)
    }
    log.Println("Server exited properly")
}
