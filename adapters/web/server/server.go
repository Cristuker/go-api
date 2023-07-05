package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/codeedu/go-hexagonal/adapters/web/server/handler"
	"github.com/codeedu/go-hexagonal/application"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {
	// Roteador, ajuda a trabalhar com as rotas
	r := mux.NewRouter()
	// Middleware
	n := negroni.New(
		negroni.NewLogger(),
	)
	handler.MakeProductHandler(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "WebServer Log:", log.Lshortfile),
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Running at 8080")
}
