package pkg

import (
	"io"
	"net/http"
	"testing"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) { return f(r) }

func TestWAClientSendsBearerRequest(t *testing.T) {
	called := false
	client := NewWAClient("http://wa.local", "key")
	client.Client = &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		called = true
		if r.Header.Get("Authorization") != "Bearer key" {
			t.Fatalf("missing bearer")
		}
		if r.URL.Path != "/api/send-message" {
			t.Fatalf("path = %s", r.URL.Path)
		}
		return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(http.NoBody), Header: http.Header{}}, nil
	})}
	if err := client.SendTextMessage("6281", "hi"); err != nil {
		t.Fatal(err)
	}
	if !called {
		t.Fatal("server not called")
	}
}
