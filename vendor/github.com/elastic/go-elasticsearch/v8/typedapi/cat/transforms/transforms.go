// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Gets configuration and usage information about transforms.
package transforms

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/cattransformcolumn"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/timeunit"
)

const (
	transformidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Transforms struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	transformid string
}

// NewTransforms type alias for index.
type NewTransforms func() *Transforms

// NewTransformsFunc returns a new instance of Transforms with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewTransformsFunc(tp elastictransport.Interface) NewTransforms {
	return func() *Transforms {
		n := New(tp)

		return n
	}
}

// Gets configuration and usage information about transforms.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/cat-transforms.html
func New(tp elastictransport.Interface) *Transforms {
	r := &Transforms{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Transforms) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("transforms")

		method = http.MethodGet
	case r.paramSet == transformidMask:
		path.WriteString("/")
		path.WriteString("_cat")
		path.WriteString("/")
		path.WriteString("transforms")
		path.WriteString("/")

		path.WriteString(r.transformid)

		method = http.MethodGet
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r Transforms) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Transforms query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a transforms.Response
func (r Transforms) Do(ctx context.Context) (Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	return nil, errorResponse
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r Transforms) IsSuccess(ctx context.Context) (bool, error) {
	res, err := r.Perform(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(ioutil.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	return false, nil
}

// Header set a key, value pair in the Transforms headers map.
func (r *Transforms) Header(key, value string) *Transforms {
	r.headers.Set(key, value)

	return r
}

// TransformId A transform identifier or a wildcard expression.
// If you do not specify one of these options, the API returns information for
// all transforms.
// API Name: transformid
func (r *Transforms) TransformId(transformid string) *Transforms {
	r.paramSet |= transformidMask
	r.transformid = transformid

	return r
}

// AllowNoMatch Specifies what to do when the request: contains wildcard expressions and
// there are no transforms that match; contains the `_all` string or no
// identifiers and there are no matches; contains wildcard expressions and there
// are only partial matches.
// If `true`, it returns an empty transforms array when there are no matches and
// the subset of results when there are partial matches.
// If `false`, the request returns a 404 status code when there are no matches
// or only partial matches.
// API name: allow_no_match
func (r *Transforms) AllowNoMatch(allownomatch bool) *Transforms {
	r.values.Set("allow_no_match", strconv.FormatBool(allownomatch))

	return r
}

// From Skips the specified number of transforms.
// API name: from
func (r *Transforms) From(from int) *Transforms {
	r.values.Set("from", strconv.Itoa(from))

	return r
}

// H Comma-separated list of column names to display.
// API name: h
func (r *Transforms) H(cattransformcolumns ...cattransformcolumn.CatTransformColumn) *Transforms {
	tmp := []string{}
	for _, item := range cattransformcolumns {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// S Comma-separated list of column names or column aliases used to sort the
// response.
// API name: s
func (r *Transforms) S(cattransformcolumns ...cattransformcolumn.CatTransformColumn) *Transforms {
	tmp := []string{}
	for _, item := range cattransformcolumns {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// Time The unit used to display time values.
// API name: time
func (r *Transforms) Time(time timeunit.TimeUnit) *Transforms {
	r.values.Set("time", time.String())

	return r
}

// Size The maximum number of transforms to obtain.
// API name: size
func (r *Transforms) Size(size int) *Transforms {
	r.values.Set("size", strconv.Itoa(size))

	return r
}