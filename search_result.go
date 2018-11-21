package main

// SearchResult - a search result
type SearchResult struct {
	Totals uint64                   `json:"totals"`
	Time   int64                    `time:"time"`
	Docs   []map[string]interface{} `docs:"docs"`
}
