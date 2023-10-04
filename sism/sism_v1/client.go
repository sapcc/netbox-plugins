package sism_v1

import (
	"crypto/tls"
	"net/http"
	"net/url"
)

type Client struct {
	HttpClient *http.Client
	BaseUrl    url.URL
	AuthToken  string
}

func (c *Client) SetAuthToken(header *http.Header) {
	header.Add("Authorization", "Token "+c.AuthToken)
	header.Set("Content-Type", "application/json")
}

func New(baseUrl string, authToken string, insecureSkipVerify bool) (*Client, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	tr := http.DefaultTransport.(*http.Transport)
	tr.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: insecureSkipVerify,
	}
	res := &Client{}
	res.BaseUrl = *u
	res.HttpClient = &http.Client{
		Transport: tr,
	}
	res.AuthToken = authToken
	return res, nil
}
