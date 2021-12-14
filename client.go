package gosq

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	endpoint           *url.URL
	username, password string
	httpClient         *http.Client
}

func NewClient(endpoint, username, password string) (*Client, error) {
	if !strings.HasSuffix(endpoint, "/") {
		endpoint += "/"
	}

	url, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	c := &Client{
		endpoint:   url,
		username:   username,
		password:   password,
		httpClient: http.DefaultClient,
	}

	return c, nil
}

func (c *Client) ActivityStatus() (*ActivityStatus, error) {
	var as ActivityStatus

	err := c.newRequest(http.MethodGet, "api/ce/activity_status", nil, &as)

	if err != nil {
		return nil, err
	}

	return &as, nil
}

func (c *Client) SystemInfo() (*SystemInfo, error) {
	var result SystemInfo

	err := c.newRequest(http.MethodGet, "api/system/info", nil, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) Component(component, metricsKeys string) (*ComponentResponse, error) {
	var as ComponentResponse

	params := make(map[string]string)
	params["component"] = component
	params["metricsKeys"] = metricsKeys

	err := c.newRequest(http.MethodGet, "api/measures/component", params, &as)

	if err != nil {
		return nil, err
	}

	return &as, nil
}

func (c *Client) newRequest(method, path string, in map[string]string, out interface{}) error {
	c.endpoint.Path = path

	params := url.Values{}

	for k, v := range in {
		params.Add(k, v)
	}

	req, err := http.NewRequest(method, c.endpoint.String(), strings.NewReader(params.Encode()))

	if err != nil {
		log.Fatal(err)
	}

	req.SetBasicAuth(c.username, c.password)

	if method == http.MethodPost || method == http.MethodPut {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return err
	}

	return decodeBody(resp, &out)
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

//go:generate gojson -o component.go -name "ComponentResponse" -pkg "gosq" -input json/component.json
//go:generate gojson -o system_info.go -name "SystemInfo" -pkg "gosq" -input json/systeminfo.json
