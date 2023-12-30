package domain

type Item struct {
	Text        []string `json:"text"`
	Description string   `json:"description"`
	Title       string   `json:"title"`
	URL         string   `json:"url"`
}

type ElasticResult struct {
	ID      string   `json:"_id"`
	Ignored []string `json:"_ignored"`
	Index   string   `json:"_index"`
	Score   float64  `json:"_score"`
	Source  Item     `json:"_source"`
}

type SearchResult struct {
	Items           []*Item `json:"items"`
	PageSize        int32   `json:"pageSize"`
	PageNo          int32   `json:"pageNo"`
	TotalItemsCount int32   `json:"totalItemsCountIP"`
}
