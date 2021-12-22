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

	err := c.Request(http.MethodGet, "api/ce/activity_status", nil, &as)

	if err != nil {
		return nil, err
	}

	return &as, nil
}

func (c *Client) SystemInfo() (*SystemInfo, error) {
	var result SystemInfo

	err := c.Request(http.MethodGet, "api/system/info", nil, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *Client) Component(component, metricsKeys string, options ...Option) (*ComponentResponse, error) {
	var as ComponentResponse

	params := make(map[string]string)
	params["component"] = component
	params["metricKeys"] = metricsKeys

	for _, option := range options {
		option(params)
	}

	err := c.Request(http.MethodGet, "api/measures/component", params, &as)

	if err != nil {
		return nil, err
	}

	return &as, nil
}

func (c *Client) BranchList(project string) (*BranchList, error) {
	var as BranchList

	params := make(map[string]string)
	params["project"] = project

	err := c.Request(http.MethodGet, "api/project_branches/list", params, &as)

	if err != nil {
		return nil, err
	}

	return &as, nil
}

func (c *Client) Projects() (*Projects, error) {
	var as Projects

	params := make(map[string]string)

	err := c.Request(http.MethodGet, "api/projects/search", params, &as)

	if err != nil {
		return nil, err
	}

	return &as, nil
}

func (c *Client) AlmSettings() (*AlmSettings, error) {
	var as AlmSettings

	params := make(map[string]string)

	err := c.Request(http.MethodGet, "api/alm_settings/list", params, &as)

	if err != nil {
		return nil, err
	}

	return &as, nil
}

func (c *Client) Request(method, path string, in map[string]string, out interface{}) error {
	c.endpoint.Path = path

	params := url.Values{}
	q := c.endpoint.Query()

	for k, v := range in {
		q.Set(k, v)
		params.Add(k, v)
	}

	if method == http.MethodGet || method == http.MethodDelete {
		c.endpoint.RawQuery = q.Encode()
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
