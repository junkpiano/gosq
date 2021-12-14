package gosq

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestClient(t *testing.T) {
	_, err := NewClient("https://example.com", "test", "test")

	if err != nil {
		t.Error("failed to create a client")
	}
}

func TestComponent(t *testing.T) {
	testServe := setup()
	defer testServe.Close()

	c, err := NewClient(testServe.URL, "testtoken", "")

	if err != nil {
		t.Error(err)
	}

	_, err = c.Component("test", "ncloc")

	if err != nil {
		t.Error(err)
	}
}

func TestSystemInfo(t *testing.T) {
	testServe := setup()
	defer testServe.Close()

	c, err := NewClient(testServe.URL, "testtoken", "")

	if err != nil {
		panic(err)
	}

	_, err = c.SystemInfo()

	if err != nil {
		panic(err)
	}
}

func setup() *httptest.Server {
	muxAPI := http.NewServeMux()
	testAPIServer := httptest.NewServer(muxAPI)

	muxAPI.HandleFunc("/api/measures/component", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("json/component.json")
		if err != nil {
			// either the file does not exist
			// or something is seriously wrong with the testing environment
			fmt.Fprintf(w, "error: %s", err)
		}
		defer f.Close()

		io.Copy(w, f)
	})

	muxAPI.HandleFunc("/api/system/info", func(w http.ResponseWriter, r *http.Request) {
		f, err := os.Open("json/systeminfo.json")
		if err != nil {
			// either the file does not exist
			// or something is seriously wrong with the testing environment
			fmt.Fprintf(w, "error: %s", err)
		}
		defer f.Close()

		io.Copy(w, f)
	})

	return testAPIServer
}
