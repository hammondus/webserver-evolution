package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeCustomFS(t *testing.T) {
	type requestCase struct {
		path         string
		expectedCode int
	}

	tests := []struct {
		name        string
		allowEscape bool
		allowDot    bool
		cases       []requestCase
	}{
		{
			name:        "dot files blocked, no escape",
			allowEscape: false,
			allowDot:    false,
			cases: []requestCase{
				{"/.secretFile", http.StatusNotFound},
				{"/.secretDir/testes.txt", http.StatusNotFound},
				{"/noindex", http.StatusNotFound},
				{"/static", http.StatusOK},
				{"/escape/README.md", http.StatusNotFound},
			},
		},
		{
			name:        "dot files blocked, with escape",
			allowEscape: true,
			allowDot:    false,
			cases: []requestCase{
				{"/.secretFile", http.StatusNotFound},
				{"/.secretDir/testes.txt", http.StatusNotFound},
				{"/noindex", http.StatusNotFound},
				{"/static", http.StatusOK},
				{"/escape/README.md", http.StatusOK},
			},
		},
		{
			name:        "dot files enabled, no escape",
			allowEscape: false,
			allowDot:    true,
			cases: []requestCase{
				{"/.secretFile", http.StatusOK},
				{"/.secretDir/testes.txt", http.StatusOK},
				{"/noindex", http.StatusNotFound},
				{"/static", http.StatusOK},
				{"/escape/README.md", http.StatusNotFound},
			},
		},
		{
			name:        "dot files enabled, with escape",
			allowEscape: true,
			allowDot:    true,
			cases: []requestCase{
				{"/.secretFile", http.StatusOK},
				{"/.secretDir/testes.txt", http.StatusOK},
				{"/noindex", http.StatusNotFound},
				{"/static", http.StatusOK},
				{"/escape/README.md", http.StatusOK},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fsys, err := NewCustomFS("testdata", tt.allowEscape, tt.allowDot)
			if err != nil {
				t.Fatal(err)
			}
			for _, rc := range tt.cases {
				t.Run(rc.path, func(t *testing.T) {
					r := httptest.NewRequest(http.MethodGet, rc.path, nil)
					w := httptest.NewRecorder()
					serveCustomFS(fsys, w, r)
					if w.Code != rc.expectedCode {
						t.Errorf("expected %d, got %d", rc.expectedCode, w.Code)
					}
				})
			}
		})
	}
}
