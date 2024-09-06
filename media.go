package subclub

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type (
	// MediaParams holds valid values for uploading files.
	MediaParams struct {
		FileName string
	}

	Media struct {
		ID         string `json:"id"`
		Type       string `json:"type"`
		MediaType  string `json:"mediaType"`
		URL        string `json:"url"`
		PreviewURL string `json:"previewUrl"`
	}
)

// UploadFile uploads a file, and returns a Media struct.
func (c *Client) UploadMedia(mp *MediaParams) (*Media, error) {
	f, err := os.Open(mp.FileName)
	if err != nil {
		return nil, fmt.Errorf("open file: %s", err)
	}
	defer f.Close()

	body := &bytes.Buffer{}
	w := multipart.NewWriter(body)

	part, err := w.CreateFormFile("file", filepath.Base(f.Name()))
	if err != nil {
		return nil, fmt.Errorf("create form file: %s", err)
	}
	_, err = io.Copy(part, f)
	if err != nil {
		return nil, fmt.Errorf("copy file: %s", err)
	}

	err = w.Close()
	if err != nil {
		return nil, fmt.Errorf("close writer: %s", err)
	}

	url := fmt.Sprintf("%s%s", c.Config.BaseURL, "/media")
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, fmt.Errorf("create request: %s", err)
	}
	req.Header.Add("User-Agent", c.Config.UserAgent)
	req.Header.Add("Content-Type", w.FormDataContentType())
	req.Header.Add("Authorization", "Bearer "+c.APIKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request: %s", err)
	}
	defer resp.Body.Close()

	m := &Media{}
	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		return nil, err
	}
	/*
		if m.Code != http.StatusCreated {
			return nil, fmt.Errorf("%s", m.ErrorMessage)
		}
	*/

	return m, nil
}
