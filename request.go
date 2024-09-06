package subclub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) get(path string, r interface{}) (interface{}, error) {
	return c.request(http.MethodGet, path, nil, r)
}

func (c *Client) post(path string, data, r interface{}) (interface{}, error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(data)
	return c.request(http.MethodPost, path, b, r)
}

func (c *Client) put(path string, data, r interface{}) (interface{}, error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(data)
	return c.request(http.MethodPut, path, b, r)
}

func (c *Client) delete(path string, data map[string]string) (interface{}, error) {
	r, err := c.buildRequest(http.MethodDelete, path, nil)
	if err != nil {
		return nil, err
	}

	q := r.URL.Query()
	for k, v := range data {
		q.Add(k, v)
	}
	r.URL.RawQuery = q.Encode()

	return c.doRequest(r, nil)
}

func (c *Client) request(method, path string, data io.Reader, result interface{}) (interface{}, error) {
	r, err := c.buildRequest(method, path, data)
	if err != nil {
		return nil, err
	}

	return c.doRequest(r, result)
}

func (c *Client) buildRequest(method, path string, data io.Reader) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", c.Config.BaseURL, path)
	r, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, fmt.Errorf("Create request: %v", err)
	}
	c.prepareRequest(r)

	return r, nil
}

func (c *Client) doRequest(r *http.Request, result interface{}) (interface{}, error) {
	resp, err := c.Client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("Request: %v", err)
	}
	defer resp.Body.Close()

	var env interface{}
	if result != nil {
		env = result

		err = json.NewDecoder(resp.Body).Decode(&env)
		if err != nil {
			return nil, err
		}
	}

	return env, nil
}

func (c *Client) prepareRequest(r *http.Request) {
	ua := c.Config.UserAgent
	if ua == "" {
		ua = "go-subclub"
	}
	r.Header.Set("User-Agent", ua)
	r.Header.Add("Content-Type", "application/json")
	if c.APIKey != "" {
		r.Header.Add("Authorization", "Bearer "+c.APIKey)
	}
}
