package main

import (
	"flag"
	"log"
	"math/rand"
	"runtime"

	"github.com/alash3al/libsrchx"
	"github.com/icrowley/fake"
)

func init() {
	var err error
	flag.Parse()

	runtime.GOMAXPROCS(*flagWorkers)

	store, err = srchx.NewStore(*flagEngine, *flagStoragePath)
	if err != nil {
		log.Fatal(err)
	}

	if *flagGenFakeData > 0 {
		go func() {
			ndx, _ := store.GetIndex("test/fake")
			fake.SetLang("en")
			for i := 0; i < *flagGenFakeData; i++ {
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
}
