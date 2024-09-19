package backend

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestRenderError(t *testing.T) {
	w := httptest.NewRecorder()

	renderError(w, http.StatusNotFound, "Not Found", "The page you are looking for does not exist.")

	// Check if the status code
	if w.Code != http.StatusNotFound {
		t.Fatalf("expected status 404, got %v", w.Code)
	}

	// Check the response body content, considering the fallback message
	expectedBody := "Internal Server Error"
	if !contains(w.Body.String(), expectedBody) {
		t.Fatalf("expected body to contain %q, got %v", expectedBody, w.Body.String())
	}
}

// Helper function to check if a string contains another string
func contains(str, substr string) bool {
	return len(str) >= len(substr) && str[:len(substr)] == substr
}

func TestArtistHandler(t *testing.T) {
	os.Chdir("../")
	getReq, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ArtistsHandler)
	handler.ServeHTTP(rr, getReq)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
	body := rr.Body.String()
	fmt.Println(body)
	if !strings.Contains(body, "<!DOCTYPE html>") {
		t.Error("handler returned a non standard html!")
	}

	postReq, err := http.NewRequest("POST", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, postReq)
	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusMethodNotAllowed)
	}
}

func TestArtistDetailsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/relation?id=4", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ArtistDetailsHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
}

func TestLocationsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/locations?id=45", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LocationsHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
}

func TestDatesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/dates?id=12", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(DatesHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}
}
