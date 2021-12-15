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

	_, err = c.Component("testproject", "ncloc")

	if err != nil {
		t.Error(err)
	}
}

func TestSystemInfo(t *testing.T) {
	testServe := setup()
	defer testServe.Close()

	c, err := NewClient(testServe.URL, "testtoken", "")

	if err != nil {
		t.Error(err)
	}

	_, err = c.SystemInfo()

	if err != nil {
		t.Error(err)
	}
}

func TestBranchList(t *testing.T) {
	testServe := setup()
	defer testServe.Close()

	c, err := NewClient(testServe.URL, "testtest", "")

	if err != nil {
		t.Error(err)
	}

	_, err = c.BranchList("testproject")

	if err != nil {
		t.Error(err)
	}
}

func TestProjects(t *testing.T) {
	testServe := setup()
	defer testServe.Close()

	c, err := NewClient(testServe.URL, "testtest", "")

	if err != nil {
		t.Error(err)
	}

	_, err = c.Projects()

	if err != nil {
		t.Error(err)
	}
}

func setup() *httptest.Server {
	muxAPI := http.NewServeMux()
	testAPIServer := httptest.NewServer(muxAPI)

	listOfAPIs := make(map[string]string)
	listOfAPIs["/api/measures/component"] = "json/component.json"
	listOfAPIs["/api/system/info"] = "json/systeminfo.json"
	listOfAPIs["/api/project_branches/list"] = "json/branch_list.json"
	listOfAPIs["/api/projects/search"] = "json/projects.json"

	for k, v := range listOfAPIs {
		muxAPI.HandleFunc(k, func(w http.ResponseWriter, r *http.Request) {
			f, err := os.Open(v)
			if err != nil {
				fmt.Fprintf(w, "error: %s", err)
			}
			defer f.Close()

			io.Copy(w, f)
		})
	}

	return testAPIServer
}
