package query

var (
	// EsQuery query elastic for archive index
	EsQuery string = `{
		"query": {
		  "nested": {
			"path": "author",
			"query": {
			  "bool": {
				"must": [
				  { "match": {"author.value": "Agustian"}},
				  { "match": {"author.value": "Gofar"}}
				]
			  }
			}
		  }
		}
	  }`
)
