package main

import (
	"flag"
	"log"
	"os"
	"path"

	"github.com/alash3al/libsrchx"
)

var (
	flagListenAddr  = flag.String("listen", ":2050", "the restful server listen address")
	flagEngine      = flag.String("engine", "boltdb", "the engine to be used as a backend")
	flagStoragePath = flag.String("storage", path.Join(path.Dir(os.Args[0]), "data"), "the storage path")
)

var (
	store *srchx.Store
)

func init() {
	var err error
	flag.Parse()

	store, err = srchx.NewStore(*flagEngine, *flagStoragePath)
	if err != nil {
		log.Fatal(err)
	}
}
