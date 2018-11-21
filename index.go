package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/blevesearch/bleve/search/query"

	"github.com/imdario/mergo"

	"github.com/blevesearch/bleve"
	"github.com/satori/go.uuid"
)

// Index - the index wrapper
type Index struct {
	bleve bleve.Index
}

// NewIndex - initialize a new index wrapper
func NewIndex(ndx bleve.Index) *Index {
	i := new(Index)
	i.bleve = ndx

	return i
}

// Delete - delete a document from the index
func (i *Index) Delete(id string) {
	i.bleve.Delete(id)
	i.bleve.DeleteInternal(i.formatInternalKey(id))
}

// Get - loads a document from the index
func (i *Index) Get(id string) (data map[string]interface{}, err error) {
	if _, err = i.bleve.Document(id); err != nil {
		return nil, err
	}

	databytes, err := i.bleve.GetInternal(i.formatInternalKey(id))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(databytes, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Put - set/update a document
func (i *Index) Put(data map[string]interface{}) (document map[string]interface{}, err error) {
	if _, found := data["id"]; !found || data["id"] == "" {
		uid, err := uuid.NewV4()
		if err != nil {
			return nil, err
		}
		data["id"] = uid.String()
	}

	id, ok := data["id"].(string)
	if !ok {
		return nil, errors.New("invalid id specified, it must be string")
	}

	if document, err = i.Get(id); err == nil && document != nil {
		document["id"] = id
	}

	if err = mergo.Map(&document, data, mergo.WithOverride); err != nil {
		return nil, err
	}

	databytes, err := json.Marshal(document)
	if err != nil {
		return nil, err
	}

	if err = i.bleve.Index(id, document); err != nil {
		return nil, err
	}

	if err = i.bleve.SetInternal(i.formatInternalKey(id), databytes); err != nil {
		i.bleve.Delete(id)
		return nil, err
	}

	return document, nil
}

// Search - search in the index for the specified query
func (i *Index) Search(q query.Query, offset, size int, sort []string) (*SearchResult, error) {
	if len(sort) < 1 {
		sort = []string{"_score"}
	}
	if size < 1 {
		size = 10
	}

	searchRequest := bleve.NewSearchRequest(q)
	searchRequest.Fields = []string{"*"}
	searchRequest.IncludeLocations = true
	searchRequest.From = offset
	searchRequest.Size = size
	searchRequest.SortBy(sort)

	res, err := i.bleve.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	docs := []map[string]interface{}{}
	for _, v := range res.Hits {
		doc, err := i.Get(v.ID)
		if err != nil {
			continue
		}
		docs = append(docs, doc)
	}

	return &SearchResult{
		Totals: res.Total,
		Docs:   docs,
		Time:   int64(res.Took),
	}, nil
}

// formatInternalKey - normalize an id for the internal storage
func (i *Index) formatInternalKey(id string) []byte {
	return []byte(fmt.Sprintf("raw:document:%s", id))
}
