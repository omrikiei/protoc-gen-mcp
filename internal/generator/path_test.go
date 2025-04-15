package generator

import (
	"testing"
)

type testRequest struct {
	Name     string `mcp_path_param:"name" mcp_query_param:"name"`
	Language string `mcp_query_param:"lang"`
	Formal   bool   `mcp_query_param:"formal"`
	Repeat   int32  `mcp_query_param:"repeat"`
}

func TestParsePathParams(t *testing.T) {
	tests := []struct {
		name       string
		template   string
		path       string
		wantParams map[string]string
		wantErr    bool
	}{
		{
			name:     "simple path parameter",
			template: "/v1/hello/{name}",
			path:     "/v1/hello/john",
			wantParams: map[string]string{
				"name": "john",
			},
			wantErr: false,
		},
		{
			name:     "multiple path parameters",
			template: "/v1/users/{user_id}/posts/{post_id}",
			path:     "/v1/users/123/posts/456",
			wantParams: map[string]string{
				"user_id": "123",
				"post_id": "456",
			},
			wantErr: false,
		},
		{
			name:       "invalid path",
			template:   "/v1/hello/{name}",
			path:       "/v1/hello",
			wantParams: nil,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotParams, err := ParsePathParams(tt.template, tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePathParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !compareMaps(gotParams, tt.wantParams) {
				t.Errorf("ParsePathParams() = %v, want %v", gotParams, tt.wantParams)
			}
		})
	}
}

func TestApplyPathParams(t *testing.T) {
	tests := []struct {
		name    string
		req     *testRequest
		params  map[string]string
		wantErr bool
	}{
		{
			name: "valid path parameter",
			req:  &testRequest{},
			params: map[string]string{
				"name": "john",
			},
			wantErr: false,
		},
		{
			name:    "missing required parameter",
			req:     &testRequest{},
			params:  map[string]string{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ApplyPathParams(tt.req, tt.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyPathParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && tt.req.Name != tt.params["name"] {
				t.Errorf("ApplyPathParams() = %v, want %v", tt.req.Name, tt.params["name"])
			}
		})
	}
}

func TestApplyQueryParams(t *testing.T) {
	tests := []struct {
		name    string
		req     *testRequest
		params  map[string][]string
		wantErr bool
	}{
		{
			name: "valid query parameters",
			req:  &testRequest{},
			params: map[string][]string{
				"name":   {"john"},
				"lang":   {"en"},
				"formal": {"true"},
				"repeat": {"3"},
			},
			wantErr: false,
		},
		{
			name: "invalid boolean parameter",
			req:  &testRequest{},
			params: map[string][]string{
				"formal": {"invalid"},
			},
			wantErr: true,
		},
		{
			name: "invalid integer parameter",
			req:  &testRequest{},
			params: map[string][]string{
				"repeat": {"invalid"},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ApplyQueryParams(tt.req, tt.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyQueryParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if tt.req.Name != tt.params["name"][0] {
					t.Errorf("ApplyQueryParams() name = %v, want %v", tt.req.Name, tt.params["name"][0])
				}
				if tt.req.Language != tt.params["lang"][0] {
					t.Errorf("ApplyQueryParams() language = %v, want %v", tt.req.Language, tt.params["lang"][0])
				}
				if tt.req.Formal != true {
					t.Errorf("ApplyQueryParams() formal = %v, want %v", tt.req.Formal, true)
				}
				if tt.req.Repeat != 3 {
					t.Errorf("ApplyQueryParams() repeat = %v, want %v", tt.req.Repeat, 3)
				}
			}
		})
	}
}

func compareMaps(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
