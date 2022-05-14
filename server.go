package main

import (
	"log"
	"net/http"
	"time"

	"github.com/nicholasaiello/go-http2-api-example/routes"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

// Chunked 50 + 150
// ----------------------
//      | h1.1  | h2   | h2 push |
// 1000 | 3     | 4    | 2.1     |
// 2000 | 3.7   | 5.6  | 2.4     |

// TODO
// - add streaming
func main() {
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      routes.Router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":8443",
		Handler:      routes.Router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	g.Go(func() error {
		return server02.ListenAndServeTLS("./server.crt", "./server.key")
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
