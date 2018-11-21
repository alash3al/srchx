package main

// QueryRequest - represents a query request
type QueryRequest struct {
	QueryString string   `json:"query" form:"query" query:"query"`
	Offset      int      `json:"offset" form:"offset" query:"offset"`
	Size        int      `json:"size" form:"size" query:"size"`
	Sort        []string `json:"sort" form:"sort" query:"sort"`
}
