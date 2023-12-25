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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// JobStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/_types/Job.ts#L54-L59
type JobStatistics struct {
	Avg   Float64 `json:"avg"`
	Max   Float64 `json:"max"`
	Min   Float64 `json:"min"`
	Total Float64 `json:"total"`
}

func (s *JobStatistics) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "avg":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Avg = f
			case float64:
				f := Float64(v)
				s.Avg = f
			}

		case "max":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Max = f
			case float64:
				f := Float64(v)
				s.Max = f
			}

		case "min":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Min = f
			case float64:
				f := Float64(v)
				s.Min = f
			}

		case "total":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Total = f
			case float64:
				f := Float64(v)
				s.Total = f
			}

		}
	}
	return nil
}

// NewJobStatistics returns a JobStatistics.
func NewJobStatistics() *JobStatistics {
	r := &JobStatistics{}

	return r
}