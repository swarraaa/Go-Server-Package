package goServer

import (
	"log"
	"net/http"
	"sync"
	"time"
)

type Server struct {
    Router      *Router         
    middlewares []Middleware  
}

func NewServer() *Server {
    return &Server{
        Router: NewRouter(), 
    }
}

func (s *Server) AddRoute(path string, handlerFunc http.HandlerFunc) {
    s.Router.AddRoute(path, handlerFunc)
}

func (s *Server) StartServer(port string, wg *sync.WaitGroup) {
    srv := &http.Server{
        Addr:    ":" + port,
        Handler: s.Router,
    }

    go func() {
        log.Printf("Starting server at port %s", port)
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Could not listen on %s: %v\n", port, err)
        }
    }()

    wg.Add(1)
    go func() {
        defer wg.Done()
        time.Sleep(10 * time.Second) 
        log.Println("Shutting down server gracefully...")
        srv.Close()
    }()
}

func (s *Server) AddMiddleware(mw Middleware) {
    s.middlewares = append(s.middlewares, mw)
}

func (s *Server) StartServerWithCron(port string, wg *sync.WaitGroup, cronSchedule string, task func()) {
    s.StartServer(port, wg)
    StartCronWithRobfig(cronSchedule, task)
}
