package main

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/blevesearch/bleve"
)

// Store our main store wrapper
type Store struct {
	engine      string
	datapath    string
	indexes     map[string]*Index
	indexesLock sync.RWMutex
}

// NewStore initialize a new store
func NewStore(engine, path string) (*Store, error) {
	s := new(Store)
	s.engine = strings.ToLower(engine)
	s.datapath = filepath.Join(path, s.engine)
	s.indexes = map[string]*Index{}
	s.indexesLock = sync.RWMutex{}

	os.MkdirAll(s.datapath, 0744)

	return s, nil
}

// GetIndex load/init an index and return it
func (s *Store) GetIndex(name string) (*Index, error) {
	var err error

	name = strings.ToLower(name)
	ndx, ok := s.indexes[name]

	if !ok {
		ndx, err = s.InitIndex(name)
	}

	if err != nil {
		return nil, err
	}

	return ndx, nil
}

// InitIndex create an index and register it in our main registry
func (s *Store) InitIndex(name string) (ndx *Index, err error) {
	engine := s.engine
	name = strings.ToLower(name)
	indexPath := path.Join(s.datapath, name)

	s.indexesLock.Lock()
	defer s.indexesLock.Unlock()

	if err = os.MkdirAll(indexPath, 0744); err != nil && err != os.ErrExist {
		return nil, err
	}

	indexMapping := bleve.NewIndexMapping()
	indexMapping.DefaultAnalyzer = "srchx"

	switch engine {
	case "boltdb":
		ndx, err = initBoltIndex(indexPath, indexMapping)
	case "leveldb":
		ndx, err = initLevelIndex(indexPath, indexMapping)
	case "badgerdb", "rocksdb":
		ndx, err = initBadgerIndex(indexPath, indexMapping)
	case "scorch":
		ndx, err = initScorchIndex(indexPath, indexMapping)
	default:
		err = errors.New("unknown engine (" + (engine) + ") specfied ")
	}

	if err != nil {
		return nil, err
	}

	s.indexes[name] = ndx

	return ndx, err
}
