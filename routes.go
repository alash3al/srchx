package main

import (
	"strings"

	"github.com/blevesearch/bleve"
	"github.com/blevesearch/bleve/search/query"
	"github.com/labstack/echo"
)

/**
 * routeHome - the home route
 */
func routeHome(c echo.Context) error {
	return c.JSON(200, map[string]interface{}{
		"success": true,
		"message": "I'm ready",
	})
}

/**
 * routeIndex - index a document
 */
func routeIndex(c echo.Context) error {
	var doc map[string]interface{}
	if err := c.Bind(&doc); err != nil {
		return c.JSON(500, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	ndx, typ, id := c.Param("index"), c.Param("type"), c.Param("id")
	index, err := store.GetIndex(ndx + "/" + typ)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	if strings.ToLower(id) == "new" {
		doc["id"] = ""
	} else {
		doc["id"] = id
	}

	doc, err = index.Put(doc)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"success": true,
		"payload": doc,
	})
}

/**
 * routeBatchIndex - index multiple documents
 */
func routeBatchIndex(c echo.Context) error {
	var docs []map[string]interface{}
	if err := c.Bind(&docs); err != nil {
		return c.JSON(400, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	ndx, typ := c.Param("index"), c.Param("type")
	index, err := store.GetIndex(ndx + "/" + typ)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	success := []string{}
	errs := []string{}

	for _, doc := range docs {
		doc, err = index.Put(doc)
		if err == nil {
			success = append(success, doc["id"].(string))
		} else {
			errs = append(errs, err.Error())
		}
	}

	return c.JSON(200, map[string]interface{}{
		"success": len(success) > len(errs),
		"errors":  errs,
		"payload": success,
	})
}

/**
 * routeGet - get a document
 */
func routeGet(c echo.Context) error {
	ndx, typ, id := c.Param("index"), c.Param("type"), c.Param("id")
	index, err := store.GetIndex(ndx + "/" + typ)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	doc, err := index.Get(id)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"success": true,
		"payload": doc,
	})
}

/**
 * routeDelete - delete a document
 */
func routeDelete(c echo.Context) error {
	ndx, typ, id := c.Param("index"), c.Param("type"), c.Param("id")
	index, err := store.GetIndex(ndx + "/" + typ)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	index.Delete(id)

	return c.JSON(200, map[string]interface{}{
		"success": true,
	})
}

/**
 * routeSearch - search for documents
 */
func routeSearch(c echo.Context) error {
	var input QueryRequest
	var q query.Query

	if err := c.Bind(&input); err != nil {
		return c.JSON(400, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	ndx, typ := c.Param("index"), c.Param("type")
	index, err := store.GetIndex(ndx + "/" + typ)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	if input.QueryString != "" {
		q = query.Query(bleve.NewQueryStringQuery(input.QueryString))
	}

	if q == nil {
		q = query.Query(bleve.NewMatchAllQuery())
	}

	res, err := index.Search(q, input.Offset, input.Size, input.Sort)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"success": false,
			"error":   err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"success": true,
		"payload": res,
	})
}
