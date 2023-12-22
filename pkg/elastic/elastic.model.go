package elastic

type SearchResult struct {
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []interface{} `json:"hits"`
	} `json:"hits"`
}
