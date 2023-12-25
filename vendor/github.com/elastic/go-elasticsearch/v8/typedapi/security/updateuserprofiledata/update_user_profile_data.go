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

// Update application specific data for the user profile of the given unique ID.
package updateuserprofiledata

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
)

const (
	uidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateUserProfileData struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	uid string
}

// NewUpdateUserProfileData type alias for index.
type NewUpdateUserProfileData func(uid string) *UpdateUserProfileData

// NewUpdateUserProfileDataFunc returns a new instance of UpdateUserProfileData with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewUpdateUserProfileDataFunc(tp elastictransport.Interface) NewUpdateUserProfileData {
	return func(uid string) *UpdateUserProfileData {
		n := New(tp)

		n._uid(uid)

		return n
	}
}

// Update application specific data for the user profile of the given unique ID.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-update-user-profile-data.html
func New(tp elastictransport.Interface) *UpdateUserProfileData {
	r := &UpdateUserProfileData{
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
func (r *UpdateUserProfileData) Raw(raw io.Reader) *UpdateUserProfileData {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *UpdateUserProfileData) Request(req *Request) *UpdateUserProfileData {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *UpdateUserProfileData) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for UpdateUserProfileData: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == uidMask:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("profile")
		path.WriteString("/")

		path.WriteString(r.uid)
		path.WriteString("/")
		path.WriteString("_data")

		method = http.MethodPut
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
func (r UpdateUserProfileData) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the UpdateUserProfileData query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a updateuserprofiledata.Response
func (r UpdateUserProfileData) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the UpdateUserProfileData headers map.
func (r *UpdateUserProfileData) Header(key, value string) *UpdateUserProfileData {
	r.headers.Set(key, value)

	return r
}

// Uid A unique identifier for the user profile.
// API Name: uid
func (r *UpdateUserProfileData) _uid(uid string) *UpdateUserProfileData {
	r.paramSet |= uidMask
	r.uid = uid

	return r
}

// IfSeqNo Only perform the operation if the document has this sequence number.
// API name: if_seq_no
func (r *UpdateUserProfileData) IfSeqNo(sequencenumber string) *UpdateUserProfileData {
	r.values.Set("if_seq_no", sequencenumber)

	return r
}

// IfPrimaryTerm Only perform the operation if the document has this primary term.
// API name: if_primary_term
func (r *UpdateUserProfileData) IfPrimaryTerm(ifprimaryterm string) *UpdateUserProfileData {
	r.values.Set("if_primary_term", ifprimaryterm)

	return r
}

// Refresh If 'true', Elasticsearch refreshes the affected shards to make this operation
// visible to search, if 'wait_for' then wait for a refresh to make this
// operation
// visible to search, if 'false' do nothing with refreshes.
// API name: refresh
func (r *UpdateUserProfileData) Refresh(refresh refresh.Refresh) *UpdateUserProfileData {
	r.values.Set("refresh", refresh.String())

	return r
}

// Data Non-searchable data that you want to associate with the user profile.
// This field supports a nested data structure.
// API name: data
func (r *UpdateUserProfileData) Data(data map[string]json.RawMessage) *UpdateUserProfileData {

	r.req.Data = data

	return r
}

// Labels Searchable data that you want to associate with the user profile. This
// field supports a nested data structure.
// API name: labels
func (r *UpdateUserProfileData) Labels(labels map[string]json.RawMessage) *UpdateUserProfileData {

	r.req.Labels = labels

	return r
}
