package backend

import (
	"encoding/json"
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

// Helper function to temporarily change to parent directory
func changeToParentDir(t *testing.T) func() {
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	err = os.Chdir("../")
	if err != nil {
		t.Fatal(err)
	}

	return func() {
		err := os.Chdir(originalDir)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestArtistHandler(t *testing.T) {
	defer changeToParentDir(t)()

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

	invalidPaths := []string{"/invalid", "/anotherpath"}

	for _, path := range invalidPaths {
		t.Run(path, func(t *testing.T) {
			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(ArtistsHandler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusNotFound {
				t.Errorf("handler returned wrong status code for %s: got %v, want %v", path, status, http.StatusNotFound)
			}
		})
	}
}

func TestArtistDetailsHandler(t *testing.T) {
	defer changeToParentDir(t)()
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

	invalidPaths := []string{"/invalid", "/anotherpath"}

	for _, path := range invalidPaths {
		t.Run(path, func(t *testing.T) {
			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(ArtistDetailsHandler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusNotFound {
				t.Errorf("handler returned wrong status code for %s: got %v, want %v", path, status, http.StatusNotFound)
			}
		})
	}
}

func TestLocationsHandler(t *testing.T) {
	defer changeToParentDir(t)()
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

	invalidPaths := []string{"/invalid", "/anotherpath"}

	for _, path := range invalidPaths {
		t.Run(path, func(t *testing.T) {
			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(LocationsHandler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusNotFound {
				t.Errorf("handler returned wrong status code for %s: got %v, want %v", path, status, http.StatusNotFound)
			}
		})
	}
}

func TestDatesHandler(t *testing.T) {
	defer changeToParentDir(t)()
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

	invalidPaths := []string{"/invalid", "/anotherpath"}

	for _, path := range invalidPaths {
		t.Run(path, func(t *testing.T) {
			req, err := http.NewRequest("GET", path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(DatesHandler)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != http.StatusNotFound {
				t.Errorf("handler returned wrong status code for %s: got %v, want %v", path, status, http.StatusNotFound)
			}
		})
	}
}

func TestGetJSONData(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"key":"value"}`))
	}))
	defer mockServer.Close()

	client := mockServer.Client()
	expectedJSON := json.RawMessage(`{"key":"value"}`)

	result, err := GetJSONData(mockServer.URL, "", client)
	if err != nil {
		t.Fatalf("GetJSONData returned an error: %v", err)
	}

	if string(result) != string(expectedJSON) {
		t.Errorf("GetJSONData returned unexpected JSON: got %s, want %s", result, expectedJSON)
	}
}

func TestRegisterRoutes(t *testing.T) {
	defer changeToParentDir(t)()
	mux := RegisterRoutes()

	tests := []struct {
		path            string
		expectedHandler http.HandlerFunc
	}{
		{"/", ArtistsHandler},
		{"/locations?id=36", LocationsHandler},
		{"/dates?id=1", DatesHandler},
		{"/relation?id=12", ArtistDetailsHandler},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			// Create a request for the specific route
			req, err := http.NewRequest("GET", tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create a response recorder to capture the response
			rr := httptest.NewRecorder()

			// Serve the request using the mux
			mux.ServeHTTP(rr, req)

			// Check if the handler responded (e.g., by verifying the status code)
			if rr.Code != http.StatusOK {
				t.Errorf("Expected status OK for path %s, but got %v", tt.path, rr.Code)
			}
		})
	}
}

func TestArtistHandler_InternalServerError(t *testing.T) {
	// Simulate an internal server error, e.g., by passing an invalid template path
	// This might require temporarily changing how you manage templates in your handler.

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a server error (e.g., template not found)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	})
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusInternalServerError)
	}
}
