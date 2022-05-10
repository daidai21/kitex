/*
 * Copyright 2021 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package descriptor

import (
	"encoding/json"
	"net/http"
	"net/url"
)

// Cookies ...
type Cookies map[string]string

// HTTPRequest ...
type HTTPRequest struct {
	Header  http.Header
	Query   url.Values
	Cookies Cookies
	Method  string
	Host    string
	Path    string
	Params  *Params // path params
	RawBody []byte  // 二进制后的http的参数body信息
	Body    map[string]interface{}
}

// HTTPResponse ...
type HTTPResponse struct {
	Header     http.Header
	StatusCode int32
	Body       map[string]interface{}
}

// NewHTTPResponse ...
func NewHTTPResponse() *HTTPResponse {
	return &HTTPResponse{
		Header: http.Header{},
		Body:   map[string]interface{}{},
	}
}

// Write to ResponseWriter
func (resp *HTTPResponse) Write(w http.ResponseWriter) error {
	w.WriteHeader(int(resp.StatusCode))
	for k := range resp.Header {
		w.Header().Set(k, resp.Header.Get(k))
	}
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(resp.Body)
}

// Param in request path
type Param struct {
	Key   string
	Value string
}

// Params and recyclable
type Params struct {
	params   []Param
	recycle  func(*Params)
	recycled bool
}

// Recycle the Params
func (ps *Params) Recycle() {
	if ps.recycled {
		return
	}
	ps.recycled = true
	ps.recycle(ps)
}

// ByName search Param by given name
func (ps *Params) ByName(name string) string {
	for _, p := range ps.params {
		if p.Key == name {
			return p.Value
		}
	}
	return ""
}
