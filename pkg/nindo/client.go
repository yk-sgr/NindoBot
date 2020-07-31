package nindo

import (
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"net/url"
)

var json = jsoniter.ConfigFastest

var DefaultClient = New(nil)

type Client struct {
	baseURL    *url.URL
	httpClient *http.Client
}

type Options struct {
	BaseURL    *url.URL
	HTTPClient *http.Client
}

func New(opts *Options) *Client {
	if opts == nil {
		opts = &Options{}
	}
	return createClient(opts)
}

func URL(scheme, host string) *url.URL {
	return &url.URL{Scheme: scheme, Host: host}
}

func createClient(opts *Options) *Client {
	if opts.BaseURL == nil {
		opts.BaseURL = &url.URL{Scheme: DefaultURLScheme, Host: DefaultURLHost}
	}
	if opts.HTTPClient == nil {
		opts.HTTPClient = &http.Client{}
	}
	return &Client{
		baseURL:    opts.BaseURL,
		httpClient: opts.HTTPClient,
	}
}
