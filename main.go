package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var port string
var rnd *rand.Rand

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	rnd = rand.New(s1)

	flag.StringVar(&port, "port", "8000", "Port to listen on")
	flag.Parse()

	if err := startServer(port); err != nil {
		log.Fatal(err)
	}
}

func startServer(port string) error {

	http.HandleFunc("/", handler)
	http.HandleFunc("/health", health)

	sig := make(chan os.Signal)
	done := make(chan bool)
	errc := make(chan error)

	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	go func(sig <-chan os.Signal, done chan<- bool) {
		s := <-sig
		log.Printf("Received %v", s)
		done <- true
	}(sig, done)

	go func(errc chan<- error) {
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			errc <- err
		}
	}(errc)

	log.Printf("Listening on port %s", port)

	select {
	case err := <-errc:
		return err
	case <-done:
		log.Printf("Shutting down")
		return nil
	}
}

func health(w http.ResponseWriter, r *http.Request) {
	log.Print("Health check")
	fmt.Fprintf(w, "ok")
}

func handler(w http.ResponseWriter, r *http.Request) {
	num := rnd.Intn(2000)
	log.Printf("Sleep %v\n", num)

	time.Sleep(time.Duration(num) * time.Millisecond)
	fmt.Fprintf(w, "slept %v ms", num)
}
