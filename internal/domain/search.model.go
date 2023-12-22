package domain

type Item struct {
	Id          string `bson:"_id" json:"id"`
	Title       string `bson:"title" json:"title"`
	Description string `bson:"description" json:"description"`
	Url         string `bson:"url" json:"url"`
}

type SearchResult struct {
	Items           []*Item `json:"items"`
	PageSize        int32   `json:"pageSize"`
	PageNo          int32   `json:"pageNo"`
	TotalItemsCount int32   `json:"totalItemsCountIP"`
}
