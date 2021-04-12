package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

// TODO: optional file compression
// TODO: add secret user key to metadata
//       allows deleting, querying meta etc.
//       send in header
//       (control this with flags: canDelete etc.)
// TODO: allow in-memory store
// TODO: UTC switch flag (vs local time)
// TODO: make -key mandatory?
// TODO: unify the data and metadata files somehow...

type Metadata struct {
	Name     string    `json:"fileName"`
	Size     int64     `json:"fileSize"`
	Uploader string    `json:"uploader"`
	Uploaded time.Time `json:"uploadTime"`
}

var (
	flagAddress     = flag.String("addr", "", "TODO")
	flagPort        = flag.Uint("port", 8080, "TODO")
	flagPath        = flag.String("path", "/", "TODO")
	flagStorage     = flag.String("dir", "database", "TODO")
	flagSecret      = flag.String("key", "", "TODO")
	flagAllowList   = flag.Bool("list", false, "TODO")
	flagAllowDelete = flag.Bool("delete", false, "TODO") // TODO: unused
)

var mutex sync.RWMutex

func main() {
	flag.Parse()

	if *flagSecret == "" {
		log.Print("Warning: Using an empty key may be insecure.")
	}

	err := os.MkdirAll(*flagStorage, os.ModePerm)
	if err != nil {
		panic(fmt.Errorf("failed to create store directory: %w", err))
	}

	if *flagAllowList {
		http.HandleFunc(*flagPath+"list", handleList) // TODO: properly join url
	}

	http.HandleFunc(*flagPath, handleRoot)

	fullAddr := fmt.Sprintf("%s:%d", *flagAddress, *flagPort)

	log.Printf("Listening on %s...\n", fullAddr)
	log.Fatal(http.ListenAndServe(fullAddr, nil))
}
