package recdrive

type Options struct {
	// `x-auth-token` field in headers, used to auth by recdrive
	AuthToken   string
	ApiEndpoint string
}

func (opt *Options) Build() *Options {
	if opt.ApiEndpoint == "" {
		opt.ApiEndpoint = apiEndpointDefault
	}

	return opt
}
