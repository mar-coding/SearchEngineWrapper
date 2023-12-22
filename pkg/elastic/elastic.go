package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"

	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type Elastic interface {
	ExecuteQuery(ctx context.Context, indexName string, fields []string, query string, page int32, pageSize int32, response any) (int, error)
}

type Client struct {
	client *elasticsearch.Client
}

func NewElasticSearch(ctx context.Context, addresses []string, username, password string, development bool) (*Client, error) {
	cfg := elasticsearch.Config{
		Addresses: addresses,
		Username:  username,
		Password:  password,
	}

	if development {
		cfg.EnableDebugLogger = true
	}

	cli, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	cli.Ping.WithContext(ctx)

	return &Client{client: cli}, nil
}

func (e *Client) ExecuteQuery(ctx context.Context, indexName string, fields []string, query string, pageNum int32, pageSize int32, response any) (int, error) {
	jsonQuery, err := json.Marshal(fields)
	if err != nil {
		return 0, err
	}

	query = strings.TrimSpace(query)
	pageStart := (pageNum - 1) * pageSize

	req := esapi.SearchRequest{
		Index: []string{indexName},
		Body:  strings.NewReader(fmt.Sprintf(`{"query": {"query_string": {"query": "%s", "fields": %s}}}`, query, jsonQuery)),
		From:  esapi.IntPtr(int(pageStart)),
		Size:  esapi.IntPtr(int(pageSize)),
	}

	res, err := req.Do(ctx, e.client)
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()

	var result SearchResult

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return 0, err
	}

	b, err := json.Marshal(result.Hits.Hits)
	if err != nil {
		return 0, err
	}

	if err := json.Unmarshal(b, response); err != nil {
		return 0, err
	}

	return result.Hits.Total.Value, nil
}
