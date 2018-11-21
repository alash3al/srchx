package main

import (
	"github.com/alash3al/bbadger"
	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/index/scorch"
	"github.com/blevesearch/bleve/index/store/boltdb"
	"github.com/blevesearch/bleve/index/store/goleveldb"
	"github.com/blevesearch/bleve/mapping"
	"github.com/wrble/flock/index/upsidedown"
)

// create/open a new boltdb based index
func initBoltIndex(path string, mapping mapping.IndexMapping) (*Index, error) {
	typ := upsidedown.Name
	ndx, err := bleve.NewUsing(path, mapping, typ, boltdb.Name, nil)
	if err != nil && err == bleve.ErrorIndexPathExists {
		ndx, err = bleve.Open(path)
	}
	return NewIndex(ndx), err
}

// create/open a new leveldb based index
func initLevelIndex(path string, mapping mapping.IndexMapping) (*Index, error) {
	typ := upsidedown.Name
	ndx, err := bleve.NewUsing(path, mapping, typ, goleveldb.Name, nil)
	if err != nil && err == bleve.ErrorIndexPathExists {
		ndx, err = bleve.Open(path)
	}
	return NewIndex(ndx), err
}

// create/open a new badgerdb based index
func initBadgerIndex(path string, mapping mapping.IndexMapping) (*Index, error) {
	ndx, err := bbadger.BleveIndex(path, mapping)
	return NewIndex(ndx), err
}

// create/open a new scorch based index
func initScorchIndex(path string, mapping mapping.IndexMapping) (*Index, error) {
	typ := scorch.Name
	ndx, err := bleve.NewUsing(path, mapping, typ, scorch.Name, nil)
	if err != nil && err == bleve.ErrorIndexPathExists {
		ndx, err = bleve.Open(path)
	}
	return NewIndex(ndx), err
}
