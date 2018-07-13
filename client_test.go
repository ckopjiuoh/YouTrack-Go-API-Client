package ytclient

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"fmt"
)

func TestNewYouTrackClient_Negative(t *testing.T) {
	_, err := NewYouTrackClient(
		"some",
		"http://127.0.0.1",
		true,
	)

	if err == nil {
		t.Error(err.Error())
		t.Fail()
	}
}

func TestNewYouTrackClient_Positive(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintln(w, "Hello, client")
	}))

	defer ts.Close()

	_, err := NewYouTrackClient(
		"some",
		ts.URL,
		true,
	)

	if err != nil {
		t.Error(err.Error())
		t.Fail()
	}
}

func TestNewYouTrackClient_Negative_WrongStatusCode(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(403)
		fmt.Fprintln(w, "Hello, client")
	}))

	defer ts.Close()

	_, err := NewYouTrackClient(
		"some",
		ts.URL,
		true,
	)

	if err == nil {
		t.Error(err.Error())
		t.Fail()
	}
}
