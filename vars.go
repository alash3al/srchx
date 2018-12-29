package main

import (
	"flag"
	"os"
	"path"
	"runtime"

	"github.com/alash3al/libsrchx"
)

var (
	flagListenAddr  = flag.String("listen", ":2050", "the restful server listen address")
	flagEngine      = flag.String("engine", "boltdb", "the engine to be used as a backend")
	flagStoragePath = flag.String("storage", path.Join(path.Dir(os.Args[0]), "data"), "the storage path")
	flagWorkers     = flag.Int("workers", runtime.NumCPU(), "number of workers to be used")
	flagGenFakeData = flag.Int("testdata", 0, "this will generate a `test/fake` collection with fake data just for testing")
)

var (
	store *srchx.Store
)
