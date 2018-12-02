package main

import (
	"flag"
	"log"
	"os"
	"path"
	"runtime"

	"github.com/alash3al/libsrchx"
)

var (
	flagListenAddr  = flag.String("listen", ":2050", "the restful server listen address")
	flagEngine      = flag.String("engine", "boltdb", "the engine to be used as a backend")
	flagStoragePath = flag.String("storage", path.Join(path.Dir(os.Args[0]), "data"), "the storage path")
	flagWorkers     = flag.Int("workers", runtime.NumCPU()*4, "number of workers to be used")
)

var (
	store *srchx.Store
)

func init() {
	var err error
	flag.Parse()

	runtime.GOMAXPROCS(*flagWorkers)

	store, err = srchx.NewStore(*flagEngine, *flagStoragePath)
	if err != nil {
		log.Fatal(err)
	}
}
