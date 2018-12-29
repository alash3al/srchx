package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"path"
	"runtime"

	"github.com/icrowley/fake"

	"github.com/alash3al/libsrchx"
)

var (
	flagListenAddr  = flag.String("listen", ":2050", "the restful server listen address")
	flagEngine      = flag.String("engine", "boltdb", "the engine to be used as a backend")
	flagStoragePath = flag.String("storage", path.Join(path.Dir(os.Args[0]), "data"), "the storage path")
	flagWorkers     = flag.Int("workers", runtime.NumCPU(), "number of workers to be used")
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

	go func() {
		ndx, _ := store.GetIndex("test/fake")
		fake.SetLang("en")
		for i := 0; i < 100000; i++ {
			ndx.Put(map[string]interface{}{
				"full_name": fake.FullName(),
				"country":   fake.Country(),
				"brand":     fake.Brand(),
				"email":     fake.EmailAddress(),
				"ip":        fake.IPv4(),
				"industry":  fake.Industry(),
				"age":       rand.Intn(100),
				"salary":    rand.Intn(20) * 1000,
				"power":     rand.Intn(10),
				"family":    rand.Intn(10),
			})
		}
	}()
}
