package recdrive

import "net/http"

type Options struct {
	// `x-auth-token` field in headers, used to auth by recdrive
	AuthToken   string
	ApiEndpoint string
	HttpClient  *http.Client
}

func (opt Options) Build() Options {
	if opt.ApiEndpoint == "" {
		opt.ApiEndpoint = apiEndpointDefault
	}

	if opt.HttpClient == nil {
		opt.HttpClient = http.DefaultClient
	}

	return opt
}
