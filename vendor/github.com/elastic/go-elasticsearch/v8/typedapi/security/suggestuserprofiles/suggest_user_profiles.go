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

// Get suggestions for user profiles that match specified search criteria.
package suggestuserprofiles

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type SuggestUserProfiles struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int
}

// NewSuggestUserProfiles type alias for index.
type NewSuggestUserProfiles func() *SuggestUserProfiles

// NewSuggestUserProfilesFunc returns a new instance of SuggestUserProfiles with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewSuggestUserProfilesFunc(tp elastictransport.Interface) NewSuggestUserProfiles {
	return func() *SuggestUserProfiles {
		n := New(tp)

		return n
	}
}

// Get suggestions for user profiles that match specified search criteria.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-suggest-user-profile.html
func New(tp elastictransport.Interface) *SuggestUserProfiles {
	r := &SuggestUserProfiles{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),

		req: NewRequest(),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *SuggestUserProfiles) Raw(raw io.Reader) *SuggestUserProfiles {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *SuggestUserProfiles) Request(req *Request) *SuggestUserProfiles {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *SuggestUserProfiles) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if len(r.deferred) > 0 {
		for _, f := range r.deferred {
			deferredErr := f(r.req)
			if deferredErr != nil {
				return nil, deferredErr
			}
		}
	}

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for SuggestUserProfiles: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("profile")
		path.WriteString("/")
		path.WriteString("_suggest")

		method = http.MethodPost
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

	if req.Header.Get("Content-Type") == "" {
		if r.buf.Len() > 0 {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=8")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r SuggestUserProfiles) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the SuggestUserProfiles query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a suggestuserprofiles.Response
func (r SuggestUserProfiles) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
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

// Header set a key, value pair in the SuggestUserProfiles headers map.
func (r *SuggestUserProfiles) Header(key, value string) *SuggestUserProfiles {
	r.headers.Set(key, value)

	return r
}

// Data List of filters for the `data` field of the profile document.
// To return all content use `data=*`. To return a subset of content
// use `data=<key>` to retrieve content nested under the specified `<key>`.
// By default returns no `data` content.
// API name: data
func (r *SuggestUserProfiles) Data(data ...string) *SuggestUserProfiles {
	r.req.Data = data

	return r
}

// Hint Extra search criteria to improve relevance of the suggestion result.
// Profiles matching the spcified hint are ranked higher in the response.
// Profiles not matching the hint don't exclude the profile from the response
// as long as the profile matches the `name` field query.
// API name: hint
func (r *SuggestUserProfiles) Hint(hint *types.Hint) *SuggestUserProfiles {

	r.req.Hint = hint

	return r
}

// Name Query string used to match name-related fields in user profile documents.
// Name-related fields are the user's `username`, `full_name`, and `email`.
// API name: name
func (r *SuggestUserProfiles) Name(name string) *SuggestUserProfiles {

	r.req.Name = &name

	return r
}

// Size Number of profiles to return.
// API name: size
func (r *SuggestUserProfiles) Size(size int64) *SuggestUserProfiles {

	r.req.Size = &size

	return r
}