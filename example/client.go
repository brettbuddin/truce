package example

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

var _ = bytes.Compare

type Client struct {
	client *http.Client
	host   *url.URL
}

func NewClient(host string) (*Client, error) {
	u, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	return &Client{client: http.DefaultClient, host: u}, nil
}

func (c *Client) GetPost(ctxt context.Context, v0 string) (rtn Post, err error) {
	u, err := c.host.Parse(fmt.Sprintf("/api/v1/posts/%v", v0))
	if err != nil {
		return
	}

	var (
		body io.Reader
		resp *http.Response
	)

	req, err := http.NewRequest("GET", u.String(), body)
	if err != nil {
		return
	}

	resp, err = c.client.Do(req.WithContext(ctxt))
	if err != nil {
		return
	}

	defer func() {
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		_ = resp.Body.Close()
	}()

	if err = checkResponse(resp); err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&rtn)

	return
}
func (c *Client) GetPosts(ctxt context.Context) (rtn []Post, err error) {
	u, err := c.host.Parse(fmt.Sprintf("/api/v1/posts"))
	if err != nil {
		return
	}

	var (
		body io.Reader
		resp *http.Response
	)

	req, err := http.NewRequest("GET", u.String(), body)
	if err != nil {
		return
	}

	resp, err = c.client.Do(req.WithContext(ctxt))
	if err != nil {
		return
	}

	defer func() {
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		_ = resp.Body.Close()
	}()

	if err = checkResponse(resp); err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&rtn)

	return
}
func (c *Client) GetUser(ctxt context.Context, v0 string) (rtn User, err error) {
	u, err := c.host.Parse(fmt.Sprintf("/api/v1/users/%v", v0))
	if err != nil {
		return
	}

	var (
		body io.Reader
		resp *http.Response
	)

	req, err := http.NewRequest("GET", u.String(), body)
	if err != nil {
		return
	}

	resp, err = c.client.Do(req.WithContext(ctxt))
	if err != nil {
		return
	}

	defer func() {
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		_ = resp.Body.Close()
	}()

	if err = checkResponse(resp); err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&rtn)

	return
}
func (c *Client) GetUsers(ctxt context.Context) (rtn []User, err error) {
	u, err := c.host.Parse(fmt.Sprintf("/api/v1/users"))
	if err != nil {
		return
	}

	var (
		body io.Reader
		resp *http.Response
	)

	req, err := http.NewRequest("GET", u.String(), body)
	if err != nil {
		return
	}

	resp, err = c.client.Do(req.WithContext(ctxt))
	if err != nil {
		return
	}

	defer func() {
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		_ = resp.Body.Close()
	}()

	if err = checkResponse(resp); err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&rtn)

	return
}
func (c *Client) PatchPost(ctxt context.Context, v0 string, v1 PatchPostRequest) (rtn Post, err error) {
	u, err := c.host.Parse(fmt.Sprintf("/api/v1/posts/%v", v0))
	if err != nil {
		return
	}

	var (
		body io.Reader
		resp *http.Response
	)

	buf := &bytes.Buffer{}
	body = buf
	if err = json.NewEncoder(buf).Encode(v1); err != nil {
		return
	}

	req, err := http.NewRequest("PATCH", u.String(), body)
	if err != nil {
		return
	}

	resp, err = c.client.Do(req.WithContext(ctxt))
	if err != nil {
		return
	}

	defer func() {
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		_ = resp.Body.Close()
	}()

	if err = checkResponse(resp); err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&rtn)

	return
}
func (c *Client) PatchUser(ctxt context.Context, v0 string, v1 PatchUserRequest) (rtn User, err error) {
	u, err := c.host.Parse(fmt.Sprintf("/api/v1/users/%v", v0))
	if err != nil {
		return
	}

	var (
		body io.Reader
		resp *http.Response
	)

	buf := &bytes.Buffer{}
	body = buf
	if err = json.NewEncoder(buf).Encode(v1); err != nil {
		return
	}

	req, err := http.NewRequest("PATCH", u.String(), body)
	if err != nil {
		return
	}

	resp, err = c.client.Do(req.WithContext(ctxt))
	if err != nil {
		return
	}

	defer func() {
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		_ = resp.Body.Close()
	}()

	if err = checkResponse(resp); err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&rtn)

	return
}
func (c *Client) PutPost(ctxt context.Context, v0 PutPostRequest) (rtn Post, err error) {
	u, err := c.host.Parse(fmt.Sprintf("/api/v1/posts"))
	if err != nil {
		return
	}

	var (
		body io.Reader
		resp *http.Response
	)

	buf := &bytes.Buffer{}
	body = buf
	if err = json.NewEncoder(buf).Encode(v0); err != nil {
		return
	}

	req, err := http.NewRequest("PUT", u.String(), body)
	if err != nil {
		return
	}

	resp, err = c.client.Do(req.WithContext(ctxt))
	if err != nil {
		return
	}

	defer func() {
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		_ = resp.Body.Close()
	}()

	if err = checkResponse(resp); err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&rtn)

	return
}
func (c *Client) PutUser(ctxt context.Context, v0 PutUserRequest) (rtn User, err error) {
	u, err := c.host.Parse(fmt.Sprintf("/api/v1/users"))
	if err != nil {
		return
	}

	var (
		body io.Reader
		resp *http.Response
	)

	buf := &bytes.Buffer{}
	body = buf
	if err = json.NewEncoder(buf).Encode(v0); err != nil {
		return
	}

	req, err := http.NewRequest("PUT", u.String(), body)
	if err != nil {
		return
	}

	resp, err = c.client.Do(req.WithContext(ctxt))
	if err != nil {
		return
	}

	defer func() {
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		_ = resp.Body.Close()
	}()

	if err = checkResponse(resp); err != nil {
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&rtn)

	return
}

func checkResponse(resp *http.Response) error {
	switch resp.StatusCode {
	case http.StatusOK:
		return nil
	case 404:
		v := NotFound{}
		if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
			return err
		}

		return v
	case 401:
		v := NotAuthorized{}
		if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
			return err
		}

		return v

	default:
		return errors.New("unexpected status code")
	}
}
