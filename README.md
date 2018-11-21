SRCHX
======
> A standalone lightweight full-text search engine built on top of `blevesearch` and `Go` with multiple storage (`scorch`, `boltdb`, `leveldb`, `badgerdb`)

Features
==========
- Standanlone.
- Light & Fast.
- Utilizes the full CPU cores, thanks to `Go` runtime.
- Pluggable storage engines, `badgerdb` (pure Go rocksdb alternative), `boltdb`, `leveldb`, `scorch`.
- Simple & Neat RESTful API.
- Dynamic Index Creation, you don't need to create an index, just `POST` your data to the indexing endpoint.
- You can search for your documents instantly.

Installation
=============
1. Goto [Releases Page](https://github.com/alash3al/srchx/releases)
2. Choose your platform based version
3. Download it
4. Copy/Rename it as `./srchx`
5. Run `chmod +x ./srchx`
6. Run `./srchx --help` to see help info

Examples
=========
```bash

# Example 1
# Add new document to the index "twitter" and type "tweet"
$ curl --request POST \
  --url 'http://localhost:2050/twitter/tweets/_doc/new' \
  --header 'Content-Type: application/json' \
  --data '{
	"user": "u5",
	"content": "this is my tweet",
	"views": 5
}'

# Example 2
# Fetch the previously added document using its ID
$ curl http://localhost:2050/twitter/tweets/_doc/2552b636-002e-4f1a-98b1-bdb06c2464ac

# Example 3
# Search for the documents that contains u5
$ curl http://localhost:2050/twitter/tweets/_search?query=+user:u5

```

API Documentation
=================
> I published the API docs on postman [here](https://documenter.getpostman.com/view/2408647/RzZFDwf4) with examples.

Refs
=====
1. [Blevesearch](http://blevesearch.com/)
2. [QueryStringQuery](http://blevesearch.com/docs/Query-String-Query/)
3. [Sorting](http://blevesearch.com/docs/Sorting/)