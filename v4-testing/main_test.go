package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeCustomFS_DotFilesBlocked_No_Escape(t *testing.T) {
	fsys, err := NewCustomFS("testdata", false, false)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRequest(http.MethodGet, "/.secretFile", nil)
	w := httptest.NewRecorder()
	serveCustomFS(fsys, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}

	r = httptest.NewRequest(http.MethodGet, "/.secretDir/testes.txt", nil)
	w = httptest.NewRecorder()
	serveCustomFS(fsys, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}

	// should get file not found when asking for a directory that doesn't contain index.html
	r = httptest.NewRequest(http.MethodGet, "/noindex", nil)
	w = httptest.NewRecorder()
	serveCustomFS(fsys, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}

	// should status ok 200 when accessing a directory with an index.html
	r = httptest.NewRequest(http.MethodGet, "/static", nil)
	w = httptest.NewRecorder()
	serveCustomFS(fsys, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	// shoudn't be able to esacpe to a directory at a higher level than we start
	r = httptest.NewRequest(http.MethodGet, "/escape/README.md", nil)
	w = httptest.NewRecorder()
	serveCustomFS(fsys, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}
}

func TestServeCustomFS_DotFilesBlocked_With_Escape(t *testing.T) {
	fsys, err := NewCustomFS("testdata", true, false)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRequest(http.MethodGet, "/.secretFile", nil)
	w := httptest.NewRecorder()
	serveCustomFS(fsys, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}

	r = httptest.NewRequest(http.MethodGet, "/.secretDir/testes.txt", nil)
	w = httptest.NewRecorder()
	serveCustomFS(fsys, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}

	// should get file not found when asking for a directory that doesn't contain index.html
	r = httptest.NewRequest(http.MethodGet, "/noindex", nil)
	w = httptest.NewRecorder()
	serveCustomFS(fsys, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}

	// should status ok 200 when accessing a directory with an index.html
	r = httptest.NewRequest(http.MethodGet, "/static", nil)
	w = httptest.NewRecorder()
	serveCustomFS(fsys, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	// should be able to esacpe to a directory at a higher level than we start
	r = httptest.NewRequest(http.MethodGet, "/escape/README.md", nil)
	w = httptest.NewRecorder()
	serveCustomFS(fsys, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestServeCustomFS_DotFilesEnabled_No_Escape(t *testing.T) {
	fsys, err := NewCustomFS("testdata", false, true)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRequest(http.MethodGet, "/.secretFile", nil)
	w := httptest.NewRecorder()
	serveCustomFS(fsys, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	r = httptest.NewRequest(http.MethodGet, "/.secretDir/testes.txt", nil)
	w = httptest.NewRecorder()
	serveCustomFS(fsys, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	// should get file not found when asking for a directory that doesn't contain index.html
	r = httptest.NewRequest(http.MethodGet, "/noindex", nil)
	w = httptest.NewRecorder()
	serveCustomFS(fsys, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}

	// should status ok 200 when accessing a directory with an index.html
	r = httptest.NewRequest(http.MethodGet, "/static", nil)
	w = httptest.NewRecorder()
	serveCustomFS(fsys, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	// shouldn't be able to esacpe to a directory at a higher level than we start
	r = httptest.NewRequest(http.MethodGet, "/escape/README.md", nil)
	w = httptest.NewRecorder()
	serveCustomFS(fsys, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}
}

func TestServeCustomFS_DotFilesEnabled_With_Escape(t *testing.T) {
	fsys, err := NewCustomFS("testdata", true, true)
	if err != nil {
		t.Fatal(err)
	}

	r := httptest.NewRequest(http.MethodGet, "/.secretFile", nil)
	w := httptest.NewRecorder()
	serveCustomFS(fsys, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	r = httptest.NewRequest(http.MethodGet, "/.secretDir/testes.txt", nil)
	w = httptest.NewRecorder()
	serveCustomFS(fsys, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	// should get file not found when asking for a directory that doesn't contain index.html
	r = httptest.NewRequest(http.MethodGet, "/noindex", nil)
	w = httptest.NewRecorder()
	serveCustomFS(fsys, w, r)
	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}

	// should status ok 200 when accessing a directory with an index.html
	r = httptest.NewRequest(http.MethodGet, "/static", nil)
	w = httptest.NewRecorder()
	serveCustomFS(fsys, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	// shoudn be able to esacpe to a directory at a higher level than we start
	r = httptest.NewRequest(http.MethodGet, "/escape/README.md", nil)
	w = httptest.NewRecorder()
	serveCustomFS(fsys, w, r)
	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}
