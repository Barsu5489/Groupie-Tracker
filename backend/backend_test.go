package backend

import (
	"net/http"
	"net/http/httptest"
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
